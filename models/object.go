package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/uuid"
)

type Object struct {
	ID         uuid.UUID `json:"-" db:"id"`
	CreatedAt  time.Time `json:"-" db:"created_at"`
	UpdatedAt  time.Time `json:"-" db:"updated_at"`
	UploadPath string    `json:"upload_path" db:"upload_path"`
	ProviderID uuid.UUID `json:"-" db:"provider_id"`
}

// String is not required by pop and may be deleted
func (o Object) String() string {
	jo, _ := json.Marshal(o)
	return string(jo)
}

// Objects is not required by pop and may be deleted
type Objects []Object

// String is not required by pop and may be deleted
func (o Objects) String() string {
	jo, _ := json.Marshal(o)
	return string(jo)
}
