package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/hashicorp/go-version"
)

type VersionStatus string

const (
	VersionStatusUnreleased VersionStatus = "unreleased"
	VersionStatusActive     VersionStatus = "active"
	VersionStatusRevoked    VersionStatus = "revoked"
)

type VersionRequest struct {
	Version struct {
		Version     string `json:"version"`
		Description string `json:"description"`
	} `json:"version"`
}

func (v VersionRequest) ToVersion() Version {
	return Version{
		Version:             v.Version.Version,
		DescriptionMarkdown: v.Version.Description,
		Number:              v.Version.Version,
	}
}

type Version struct {
	ID                  uuid.UUID `json:"id" db:"id"`
	CreatedAt           time.Time `json:"created_at" db:"created_at"`
	UpdatedAt           time.Time `json:"updated_at" db:"updated_at"`
	Version             string    `json:"version" db:"version"`
	Status              string    `json:"status" db:"status"`
	DescriptionHTML     string    `json:"description_html" db:"description_html"`
	DescriptionMarkdown string    `json:"description_markdown" db:"description_markdown"`
	Number              string    `json:"number" db:"number"`
	ReleaseURL          string    `json:"release_url" db:"release_url"`
	RevokeURL           string    `json:"revoke_url" db:"revoke_url"`
	BoxID               uuid.UUID `json:"-" db:"box_id"`
	Providers           Providers `has_many:"providers" order_by:"name asc"`
}

// String is not required by pop and may be deleted
func (v Version) String() string {
	jv, _ := json.Marshal(v)
	return string(jv)
}

// Versions is not required by pop and may be deleted
type Versions []Version

func (v Versions) GetLatest() (Version, error) {
	latest := Version{}
	for _, ver := range v {
		if ver.Status != string(VersionStatusActive) {
			continue
		}
		if latest.Version == "" {
			latest = ver
			continue
		}
		v1, err := version.NewVersion(latest.Number)
		if err != nil {
			return latest, err
		}
		v2, err := version.NewVersion(ver.Number)
		if err != nil {
			return latest, err
		}

		if v1.LessThan(v2) {
			latest = ver
		}
	}
	return latest, nil
}

// String is not required by pop and may be deleted
func (v Versions) String() string {
	jv, _ := json.Marshal(v)
	return string(jv)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (v *Version) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: v.Version, Name: "Version"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (v *Version) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (v *Version) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
