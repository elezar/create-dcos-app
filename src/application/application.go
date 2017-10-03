package application

import (
	"errors"
	"io/ioutil"

	"github.com/elezar/create-dcos-app/src/config"
	"github.com/elezar/create-dcos-app/src/pkg"

	yaml "gopkg.in/yaml.v2"
)

type Application struct {
	Name       string
	Title      string
	Maintainer string
	Package    pkg.Package
	Config     config.Config
}

func (c *Application) ParseYaml(data []byte) error {
	if err := yaml.Unmarshal(data, c); err != nil {
		return err
	}
	err := c.Validate()
	if err != nil {
		return err
	}
	c.SetDefaults()
	return nil
}

func (c Application) Validate() error {
	if c.Name == "" {
		return errors.New("Application.Name must be specifed")
	}
	return nil
}

func (c *Application) SetDefaults() {
	if c.Title == "" {
		c.Title = c.Name
	}

	defaultPackage := pkg.GetDefaultPackage()
	defaultPackage.Name = c.Name
	defaultPackage.Maintainer = c.Maintainer
	defaultPackage.Description = c.Title
	c.Package.SetDefaults(defaultPackage)

	c.Config.SetDefaults(c.Name)

}

func LoadApplicationYaml(filename string) (Application, error) {
	a := Application{}

	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		return a, err
	}
	err = a.ParseYaml(raw)
	if err != nil {
		return a, err
	}

	return a, nil
}
