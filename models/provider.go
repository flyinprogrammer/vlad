package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type ProviderRequest struct {
	Provider struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"provider"`
}

func (p ProviderRequest) ToProvider() Provider {
	return Provider{
		Name:        p.Provider.Name,
		OriginalURL: p.Provider.URL,
	}
}

type Provider struct {
	ID          uuid.UUID `json:"id" db:"id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	Name        string    `json:"name" db:"name"`
	Hosted      bool      `json:"hosted" db:"hosted"`
	HostedToken string    `json:"hosted_token" db:"hosted_token"`
	OriginalURL string    `json:"original_url" db:"original_url"`
	DownloadURL string    `json:"download_url" db:"download_url"`
	VersionID   uuid.UUID `json:"-" db:"version_id"`
	Object      Object    `json:"-" has_one:"object"`
}

// String is not required by pop and may be deleted
func (p Provider) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Providers is not required by pop and may be deleted
type Providers []Provider

// String is not required by pop and may be deleted
func (p Providers) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (p *Provider) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: p.Name, Name: "Name"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (p *Provider) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (p *Provider) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
