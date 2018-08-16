package models

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type BoxRequest struct {
	Box struct {
		Username         string `json:"username"`
		Name             string `json:"name"`
		ShortDescription string `json:"short_description"`
		Description      string `json:"description"`
		IsPrivate        bool   `json:"is_private"`
	} `json:"box"`
}

func (b BoxRequest) ToBox() Box {
	return Box{
		Username:            b.Box.Username,
		Name:                b.Box.Name,
		ShortDescription:    b.Box.ShortDescription,
		DescriptionMarkdown: b.Box.Description,
		Private:             b.Box.IsPrivate,
		Tag:                 fmt.Sprintf("%s/%s", b.Box.Username, b.Box.Name),
	}
}

type Box struct {
	ID                  uuid.UUID `json:"id" db:"id"`
	CreatedAt           time.Time `json:"created_at" db:"created_at"`
	UpdatedAt           time.Time `json:"updated_at" db:"updated_at"`
	Tag                 string    `json:"tag" db:"tag"`
	Name                string    `json:"name" db:"name"`
	ShortDescription    string    `json:"short_description" db:"short_description"`
	DescriptionHTML     string    `json:"description_html" db:"description_html"`
	Username            string    `json:"username" db:"username"`
	DescriptionMarkdown string    `json:"description_markdown" db:"description_markdown"`
	Private             bool      `json:"private" db:"private"`
	Downloads           int       `json:"downloads" db:"downloads"`
	Versions            Versions  `has_many:"versions" order_by:"version asc"`
}

// String is not required by pop and may be deleted
func (b Box) String() string {
	jb, _ := json.Marshal(b)
	return string(jb)
}

// Boxes is not required by pop and may be deleted
type Boxes []Box

// String is not required by pop and may be deleted
func (b Boxes) String() string {
	jb, _ := json.Marshal(b)
	return string(jb)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (b *Box) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: b.Name, Name: "Name"},
		&validators.StringIsPresent{Field: b.Username, Name: "Username"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (b *Box) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (b *Box) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
