package model

type EntityType struct {
	ID          string `json:"id"`
	Name        string `json:"name"` // e.g., "content"
	Description string `json:"description"`
}
