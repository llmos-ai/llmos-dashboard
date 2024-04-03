/*
Copyright YEAR 1block.ai.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/llmos-ai/llmos-dashboard/pkg/generated/ent/chat"
	"github.com/llmos-ai/llmos-dashboard/pkg/generated/ent/modelfile"
	"github.com/llmos-ai/llmos-dashboard/pkg/generated/ent/user"
	v1 "github.com/llmos-ai/llmos-dashboard/pkg/types/v1"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	chatFields := v1.Chat{}.Fields()
	_ = chatFields
	// chatDescTitle is the schema descriptor for title field.
	chatDescTitle := chatFields[1].Descriptor()
	// chat.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	chat.TitleValidator = chatDescTitle.Validators[0].(func(string) error)
	// chatDescCreatedAt is the schema descriptor for created_at field.
	chatDescCreatedAt := chatFields[4].Descriptor()
	// chat.DefaultCreatedAt holds the default value on creation for the created_at field.
	chat.DefaultCreatedAt = chatDescCreatedAt.Default.(time.Time)
	// chatDescID is the schema descriptor for id field.
	chatDescID := chatFields[0].Descriptor()
	// chat.DefaultID holds the default value on creation for the id field.
	chat.DefaultID = chatDescID.Default.(func() uuid.UUID)
	modelfileFields := v1.Modelfile{}.Fields()
	_ = modelfileFields
	// modelfileDescUserID is the schema descriptor for user_id field.
	modelfileDescUserID := modelfileFields[0].Descriptor()
	// modelfile.UserIDValidator is a validator for the "user_id" field. It is called by the builders before save.
	modelfile.UserIDValidator = modelfileDescUserID.Validators[0].(func(int) error)
	// modelfileDescTagName is the schema descriptor for tag_name field.
	modelfileDescTagName := modelfileFields[1].Descriptor()
	// modelfile.TagNameValidator is a validator for the "tag_name" field. It is called by the builders before save.
	modelfile.TagNameValidator = modelfileDescTagName.Validators[0].(func(string) error)
	// modelfileDescModelfile is the schema descriptor for modelfile field.
	modelfileDescModelfile := modelfileFields[2].Descriptor()
	// modelfile.DefaultModelfile holds the default value on creation for the modelfile field.
	modelfile.DefaultModelfile = modelfileDescModelfile.Default.(string)
	// modelfileDescCreatedAt is the schema descriptor for created_at field.
	modelfileDescCreatedAt := modelfileFields[3].Descriptor()
	// modelfile.DefaultCreatedAt holds the default value on creation for the created_at field.
	modelfile.DefaultCreatedAt = modelfileDescCreatedAt.Default.(time.Time)
	userFields := v1.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[1].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = userDescName.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[2].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[3].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
	// userDescProfileImageURL is the schema descriptor for profile_image_url field.
	userDescProfileImageURL := userFields[5].Descriptor()
	// user.DefaultProfileImageURL holds the default value on creation for the profile_image_url field.
	user.DefaultProfileImageURL = userDescProfileImageURL.Default.(string)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[6].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(time.Time)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}
