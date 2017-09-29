package application

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultApplication(t *testing.T) {
	a := Application{}
	err := a.ParseYaml([]byte{})
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, "", a.Name)
}
