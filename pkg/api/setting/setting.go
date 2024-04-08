package setting

import (
	"context"
	"fmt"
	"os"

	entv1 "github.com/llmos-ai/llmos-dashboard/pkg/generated/ent"
	entSetting "github.com/llmos-ai/llmos-dashboard/pkg/generated/ent/setting"
	"github.com/llmos-ai/llmos-dashboard/pkg/settings"
)

type handler struct {
	client   *entv1.Client
	ctx      context.Context
	fallback map[string]string
}

func NewHandler(c *entv1.Client, ctx context.Context) handler {
	return handler{
		client:   c,
		ctx:      ctx,
		fallback: map[string]string{},
	}
}

func (h *handler) Get(name string) string {
	value := os.Getenv(settings.GetEnvKey(name))
	if value != "" {
		return value
	}
	obj, err := h.client.Setting.Query().Where(entSetting.Name(name)).Only(h.ctx)
	if err != nil {
		return h.fallback[name]
	}
	if obj.Value == "" {
		return obj.Default
	}
	return obj.Value
}

func (h *handler) Set(name, value string) error {
	envValue := os.Getenv(settings.GetEnvKey(name))
	if envValue != "" {
		return fmt.Errorf("setting %s can not be set because it is from environment variable", name)
	}

	return h.client.Setting.Update().Where(entSetting.Name(name)).SetNillableValue(&value).Exec(h.ctx)
}

func (h *handler) SetIfUnset(name, value string) error {
	obj, err := h.client.Setting.Query().Where(entSetting.Name(name)).Only(h.ctx)
	if err != nil {
		return err
	}

	if obj.Value != "" {
		return nil
	}

	return h.client.Setting.Update().Where(entSetting.Name(name)).SetValue(value).Exec(h.ctx)
}

func (h *handler) SetAll(settingsMap map[string]settings.Setting) error {
	fallback := map[string]string{}

	for name, setting := range settingsMap {
		key := settings.GetEnvKey(name)
		value := os.Getenv(key)

		obj, err := h.client.Setting.Query().Where(entSetting.Name(name)).Only(h.ctx)
		if entv1.IsNotFound(err) {

			client := h.client.Setting.Create().
				SetName(setting.Name).
				SetValue(value).
				SetDefault(setting.Default)

			fallback[setting.Name] = setting.Default
			if value != "" {
				client.SetValue(value)
				fallback[setting.Name] = value
			}

			_, err = client.Save(h.ctx)
			if err != nil {
				return err
			}

		} else if err != nil {
			return err
		} else {
			update := false
			if obj.Default != setting.Default {
				obj.Default = setting.Default
				update = true
			}
			if value != "" && obj.Value != value {
				obj.Value = value
				update = true
			}
			if obj.Value == "" {
				fallback[obj.Name] = obj.Default
			} else {
				fallback[obj.Name] = obj.Value
			}
			if update {
				_, err = h.client.Setting.UpdateOne(obj).Save(h.ctx)
				if err != nil {
					return err
				}
			}
		}
	}

	h.fallback = fallback

	return nil
}
