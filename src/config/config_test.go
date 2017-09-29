package config

import (
	"fmt"
	"log"
	"testing"
)

func TestJsonSchema(t *testing.T) {
	str, err := GenerateJsonSchema()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println()
	fmt.Println(str)
	t.Fail()
}
