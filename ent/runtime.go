// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/girakdev/girack-backend/ent/channel"
	"github.com/girakdev/girack-backend/ent/schema"
	"github.com/girakdev/girack-backend/ent/user"
	"github.com/girakdev/girack-backend/internal/pulid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	channelFields := schema.Channel{}.Fields()
	_ = channelFields
	// channelDescID is the schema descriptor for id field.
	channelDescID := channelFields[0].Descriptor()
	// channel.DefaultID holds the default value on creation for the id field.
	channel.DefaultID = channelDescID.Default.(func() pulid.ID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescAge is the schema descriptor for age field.
	userDescAge := userFields[1].Descriptor()
	// user.AgeValidator is a validator for the "age" field. It is called by the builders before save.
	user.AgeValidator = userDescAge.Validators[0].(func(int) error)
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[2].Descriptor()
	// user.DefaultName holds the default value on creation for the name field.
	user.DefaultName = userDescName.Default.(string)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() pulid.ID)
}
