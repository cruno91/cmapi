package model

import "encoding/json"

type FieldDefinition struct {
	ID       string          `json:"id"`
	BundleID string          `json:"bundle_id"`
	Name     string          `json:"name"` // e.g., "title"
	Type     string          `json:"type"` // e.g., "text", "rich_text", "image", "reference"
	Required bool            `json:"required"`
	Settings json.RawMessage `json:"settings,omitempty"` // JSON blob for field-specific options
}
