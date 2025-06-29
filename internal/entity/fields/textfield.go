package fields

import (
	"CMAPI/internal/entity"
	"encoding/json"
	"fmt"
)

type TextField struct{}

func (t TextField) Validate(value any, settings json.RawMessage) error {
	_, ok := value.(string)
	if !ok {
		return fmt.Errorf("invalid type, expected string")
	}
	return nil
}

func (t TextField) GraphQLType(settings json.RawMessage) string {
	return "String"
}

func (t TextField) SQLType(settings json.RawMessage) string {
	return "TEXT"
}

func (t TextField) IsMulti(settings json.RawMessage) bool {
	return false
}

func init() {
	entity.RegisterFieldType("text", TextField{})
}
