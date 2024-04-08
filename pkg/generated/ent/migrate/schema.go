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

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ChatsColumns holds the columns for the "chats" table.
	ChatsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "title", Type: field.TypeString},
		{Name: "models", Type: field.TypeJSON},
		{Name: "tags", Type: field.TypeJSON},
		{Name: "history", Type: field.TypeJSON},
		{Name: "messages", Type: field.TypeJSON},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeUUID},
	}
	// ChatsTable holds the schema information for the "chats" table.
	ChatsTable = &schema.Table{
		Name:       "chats",
		Columns:    ChatsColumns,
		PrimaryKey: []*schema.Column{ChatsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "chats_users_chats",
				Columns:    []*schema.Column{ChatsColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "chat_user_id",
				Unique:  false,
				Columns: []*schema.Column{ChatsColumns[7]},
			},
		},
	}
	// ModelfilesColumns holds the columns for the "modelfiles" table.
	ModelfilesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "tag_name", Type: field.TypeString, Unique: true},
		{Name: "modelfile", Type: field.TypeString, Default: ""},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeUUID},
	}
	// ModelfilesTable holds the schema information for the "modelfiles" table.
	ModelfilesTable = &schema.Table{
		Name:       "modelfiles",
		Columns:    ModelfilesColumns,
		PrimaryKey: []*schema.Column{ModelfilesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "modelfiles_users_modelfiles",
				Columns:    []*schema.Column{ModelfilesColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "modelfile_user_id_tag_name",
				Unique:  false,
				Columns: []*schema.Column{ModelfilesColumns[4], ModelfilesColumns[1]},
			},
		},
	}
	// SettingsColumns holds the columns for the "settings" table.
	SettingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "default", Type: field.TypeString, Default: ""},
		{Name: "value", Type: field.TypeString, Nullable: true},
		{Name: "is_active", Type: field.TypeBool, Default: true},
		{Name: "read_only", Type: field.TypeBool, Default: false},
		{Name: "created_at", Type: field.TypeTime},
	}
	// SettingsTable holds the schema information for the "settings" table.
	SettingsTable = &schema.Table{
		Name:       "settings",
		Columns:    SettingsColumns,
		PrimaryKey: []*schema.Column{SettingsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "setting_name",
				Unique:  true,
				Columns: []*schema.Column{SettingsColumns[1]},
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "role", Type: field.TypeEnum, Enums: []string{"admin", "user", "pending"}, Default: "pending"},
		{Name: "profile_image_url", Type: field.TypeString, Default: ""},
		{Name: "created_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "user_name_email",
				Unique:  true,
				Columns: []*schema.Column{UsersColumns[1], UsersColumns[2]},
			},
			{
				Name:    "user_role",
				Unique:  false,
				Columns: []*schema.Column{UsersColumns[4]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ChatsTable,
		ModelfilesTable,
		SettingsTable,
		UsersTable,
	}
)

func init() {
	ChatsTable.ForeignKeys[0].RefTable = UsersTable
	ModelfilesTable.ForeignKeys[0].RefTable = UsersTable
}
