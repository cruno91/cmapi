package model

type Bundle struct {
	ID           string `json:"id"`
	EntityTypeID string `json:"entity_type_id"` // FK to EntityType
	Name         string `json:"name"`           // e.g., "article"
	Label        string `json:"label"`          // Human-readable
}
