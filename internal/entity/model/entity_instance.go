package model

import (
	"encoding/json"
	"time"
)

type EntityInstance struct {
	ID        string          `json:"id"`
	Type      string          `json:"type"`   // e.g., "content"
	Bundle    string          `json:"bundle"` // e.g., "article"
	Data      json.RawMessage `json:"data"`   // user-provided content (validated)
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}
