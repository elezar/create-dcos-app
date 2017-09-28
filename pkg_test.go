package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultPackageVersionIsCorrect(t *testing.T) {
	p := NewPackage()

	assert.Equal(t, "4.0", p.packagingVersion)
}
