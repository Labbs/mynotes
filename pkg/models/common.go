package models

import (
	"database/sql/driver"

	"github.com/goccy/go-json"
)

// JSONB is a map of strings to interfaces
type JSONB map[string]any

// Value implements the driver.Valuer interface
func (j JSONB) Value() (driver.Value, error) {
	valueString, err := json.Marshal(j)
	return string(valueString), err
}

// Scan implements the sql.Scanner interface
func (j *JSONB) Scan(value any) error {
	return json.Unmarshal([]byte(value.(string)), j)
}
