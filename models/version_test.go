package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersionSimpleGetLatest(t *testing.T) {
	set := Versions{
		Version{
			Version: "set1",
			Number:  "0.0.0",
			Status:  string(VersionStatusActive),
		},
	}

	latest, err := set.GetLatest()

	assert.NoError(t, err, "GetLatest returned an error")
	assert.Equal(t, "set1", latest.Version, "latest has the wrong version")
}

func TestVersionWithMultipleGetLatest(t *testing.T) {
	set := Versions{
		Version{
			Version: "set1",
			Number:  "0.0.0",
			Status:  string(VersionStatusActive),
		},
		Version{
			Version: "set2",
			Number:  "1.0.0",
			Status:  string(VersionStatusActive),
		},
		Version{
			Version: "set3",
			Number:  "0.1.0",
			Status:  string(VersionStatusActive),
		},
		Version{
			Version: "set4",
			Number:  "0.0.1",
			Status:  string(VersionStatusActive),
		},
	}

	latest, err := set.GetLatest()

	assert.NoError(t, err, "GetLatest returned an error")
	assert.Equal(t, "set2", latest.Version, "latest has the wrong version")
}

func TestVersionWithRevokedGetLatest(t *testing.T) {
	set := Versions{
		Version{
			Version: "set1",
			Number:  "0.0.0",
			Status:  string(VersionStatusActive),
		},
		Version{
			Version: "set2",
			Number:  "0.0.1",
			Status:  string(VersionStatusRevoked),
		},
	}

	latest, err := set.GetLatest()

	assert.NoError(t, err, "GetLatest returned an error")
	assert.Equal(t, "set1", latest.Version, "latest has the wrong version")
}

func TestVersionWithUnreleasedGetLatest(t *testing.T) {
	set := Versions{
		Version{
			Version: "set2",
			Number:  "0.0.1",
			Status:  string(VersionStatusRevoked),
		},
		Version{
			Version: "set1",
			Number:  "0.0.0",
			Status:  string(VersionStatusActive),
		},
		Version{
			Version: "set3",
			Number:  "0.0.2",
			Status:  string(VersionStatusUnreleased),
		},
	}

	latest, err := set.GetLatest()

	assert.NoError(t, err, "GetLatest returned an error")
	assert.Equal(t, "set1", latest.Version, "latest has the wrong version")
}
