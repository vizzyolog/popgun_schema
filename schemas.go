package popugschema

import (
	_ "embed"
)

//go:embed user_schema
var userSchema string

func UserSchema() string {
	return userSchema
}
