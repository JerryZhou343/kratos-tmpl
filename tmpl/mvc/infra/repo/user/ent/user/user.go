// Code generated by entc, DO NOT EDIT.

package user

import (
	"github.com/mfslog/kratos-mvc/infra/repo/user/ent/schema"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAge holds the string denoting the age vertex property in the database.
	FieldAge = "age"
	// FieldName holds the string denoting the name vertex property in the database.
	FieldName = "name"

	// Table holds the table name of the user in the database.
	Table = "users"
)

// Columns holds all SQL columns are user fields.
var Columns = []string{
	FieldID,
	FieldAge,
	FieldName,
}

var (
	fields = schema.User{}.Fields()

	// descAge is the schema descriptor for age field.
	descAge = fields[0].Descriptor()
	// AgeValidator is a validator for the "age" field. It is called by the builders before save.
	AgeValidator = descAge.Validators[0].(func(int32) error)

	// descName is the schema descriptor for name field.
	descName = fields[1].Descriptor()
	// DefaultName holds the default value on creation for the name field.
	DefaultName = descName.Default.(string)
)
