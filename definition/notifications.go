package definition

import (
	"github.com/kubex/potens-go/i18n"
)

// AppNotification notification provided by your app
type AppNotification struct {
	ID          string
	Name        i18n.Translations
	Description i18n.Translations
	Message     i18n.Translations
	Icon        string
	Path        string
	Roles       []AppScope
	Permissions []AppScope
	Attributes  []AppNotificationAttribute
	Actions     []Action
}

// AppNotificationAttributeType Type of notification attribute
type AppNotificationAttributeType string

const (
	// AppNotificationAttributeTypeString String Type
	AppNotificationAttributeTypeString AppNotificationAttributeType = "string"
	// AppNotificationAttributeTypeInteger Integer Type
	AppNotificationAttributeTypeInteger AppNotificationAttributeType = "integer"
	// AppNotificationAttributeTypeFloat Float Type
	AppNotificationAttributeTypeFloat AppNotificationAttributeType = "float"
	// AppNotificationAttributeTypeBoolean Boolean Type
	AppNotificationAttributeTypeBoolean AppNotificationAttributeType = "boolean"
)

// AppNotificationAttribute Attribute on your notification
type AppNotificationAttribute struct {
	Name string
	Type AppNotificationAttributeType

	ExampleString  string  `yaml:"example_string"`
	ExampleInteger int64   `yaml:"example_integer"`
	ExampleFloat   float64 `yaml:"example_float"`
	ExampleBoolean bool    `yaml:"example_boolean"`
}
