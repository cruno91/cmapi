package entity

import "encoding/json"

type FieldType interface {
	Validate(value any, settings json.RawMessage) error
	GraphQLType(settings json.RawMessage) string
	SQLType(settings json.RawMessage) string
	IsMulti(settings json.RawMessage) bool
}

var registeredFieldTypes = map[string]FieldType{}

func RegisterFieldType(name string, impl FieldType) {
	registeredFieldTypes[name] = impl
}

func GetFieldType(name string) (FieldType, bool) {
	f, ok := registeredFieldTypes[name]
	return f, ok
}
