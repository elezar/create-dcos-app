package pkg

import (
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDetaultPackageInfo(t *testing.T) {
	p := Package{}
	err := p.ParseYaml([]byte{})
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, "", p.Name)
	assert.Equal(t, DefaultPackagingVerison, p.PackagingVersion)
	assert.Equal(t, false, p.Selected)
	assert.Equal(t, true, p.Framework)
	assert.Equal(t, DefaultMinDcosReleaseVersion, p.MinDcosReleaseVersion)
	assert.Equal(t, "", p.Description)
	assert.Equal(t, []string{"{{upgrades-from}}"}, p.UpgradesFrom)
	assert.Equal(t, []string{"{{downgrades-to}}"}, p.DowngradesTo)
}

func TestDefaultPackageVersionIsCorrect(t *testing.T) {
	p := NewPackage()

	assert.Equal(t, "4.0", p.PackagingVersion)
}

func TestDefaultIsSetFromYaml(t *testing.T) {
	packageText := `
packagingVersion: 4.0
name: foo
bar: baz
`
	var p Package

	fmt.Print(packageText)
	err := p.ParseYaml([]byte(packageText))
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, "foo", p.Name)

}

func TestLoadYamlFromFile(t *testing.T) {
	p, err := LoadYamlFile("examples/package-only.yml")
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, "package-name", p.Name)
	assert.Equal(t, DefaultPackagingVerison, p.PackagingVersion)
}
