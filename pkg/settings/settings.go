package settings

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"strconv"
	"strings"
)

var (
	settings       = map[string]Setting{}
	provider       Provider
	InjectDefaults string

	FirstLogin      = NewSetting(FirstLoginSettingName, "true")
	UIPl            = NewSetting(UIPlSettingName, "LLMOS Dashboard")
	Signup          = NewSetting(SignupEnabledSettingName, "true")
	DefaultUserRole = NewSetting(DefaultUserRoleSettingName, "pending") // options are pending, user, admin
	WebhookURL      = NewSetting(WebhookURLSettingName, "")
	JWTSecret       = NewSetting(JWTSecretSettingName, "llmos-dashboard-secret")
	TokenExpireTime = NewSetting(TokenExpireTimeSettingName, "24h")    // supported units are "h", "m", "s"
	AllowChatDelete = NewSetting(AllowChatDeletionSettingName, "true") // allow users to delete their own chat
	ModelWhiteList  = NewSetting(ModelWhitelistSettingName, "")        // empty means allow all
)

const (
	UIPlSettingName              = "ui-pl"
	FirstLoginSettingName        = "first-login"
	SignupEnabledSettingName     = "signup-enabled"
	DefaultUserRoleSettingName   = "default-user-role"
	WebhookURLSettingName        = "webhook-url"
	JWTSecretSettingName         = "jwt-secret-name"
	TokenExpireTimeSettingName   = "token-expire-time"
	AllowChatDeletionSettingName = "allow-chat-deletion"
	ModelWhitelistSettingName    = "model-whitelist"
)

func init() {
	if InjectDefaults == "" {
		return
	}
	defaults := map[string]string{}
	if err := json.Unmarshal([]byte(InjectDefaults), &defaults); err != nil {
		return
	}
	for name, defaultValue := range defaults {
		value, ok := settings[name]
		if !ok {
			continue
		}
		value.Default = defaultValue
		settings[name] = value
	}
}

type Provider interface {
	Get(name string) string
	Set(name, value string) error
	SetIfUnset(name, value string) error
	SetAll(settings map[string]Setting) error
}

type Setting struct {
	Name     string
	Default  string
	ReadOnly bool
}

func (s Setting) SetIfUnset(value string) error {
	if provider == nil {
		return s.Set(value)
	}
	return provider.SetIfUnset(s.Name, value)
}

func (s Setting) Set(value string) error {
	if provider == nil {
		s, ok := settings[s.Name]
		if ok {
			s.Default = value
			settings[s.Name] = s
		}
	} else {
		return provider.Set(s.Name, value)
	}
	return nil
}

func (s Setting) Get() string {
	fmt.Println("get provider:", provider)
	if provider == nil {
		s := settings[s.Name]
		return s.Default
	}
	return provider.Get(s.Name)
}

func (s Setting) GetInt() int {
	v := s.Get()
	i, err := strconv.Atoi(v)
	if err == nil {
		return i
	}
	slog.Error("failed to parse setting %s=%s as int: %v", s.Name, v, err)
	i, err = strconv.Atoi(s.Default)
	if err != nil {
		return 0
	}
	return i
}

func SetProvider(p Provider) error {
	if err := p.SetAll(settings); err != nil {
		return err
	}
	provider = p
	return nil
}

func NewSetting(name, def string) Setting {
	s := Setting{
		Name:    name,
		Default: def,
	}
	settings[s.Name] = s
	return s
}

func GetEnvKey(key string) string {
	return "LLMOS" + strings.ToUpper(strings.Replace(key, "-", "_", -1))
}
