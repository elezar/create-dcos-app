package config

import (
	"fmt"
	"testing"
)

func TestJsonSchema(t *testing.T) {

	c := Config{}
	c.SetDefaults("serviceName")

	str := GenerateJsonSchema(c)

	fmt.Println()
	fmt.Println(str)
	t.Fail()
}
