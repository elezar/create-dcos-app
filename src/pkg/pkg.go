package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

// DefaultPackagingVerison is the defaul packing version
const DefaultPackagingVerison string = "4.0"

const DefaultMinDcosReleaseVersion string = "1.9"

// {
// "packagingVersion": "4.0",
// "upgradesFrom": ["{{upgrades-from}}"],
// "downgradesTo": ["{{downgrades-to}}"],
// "minDcosReleaseVersion": "1.9",
// 	"name": "beta-cassandra",
// 	"version": "{{package-version}}",
// 	"maintainer": "support@mesosphere.io",
// 	"description": "Apache Cassandra",
// 	"selected": false,
// 	"framework": true,
// 	"tags": ["cassandra"],
// 	"preInstallNotes": "This DC/OS Service is currently a beta candidate undergoing testing as part of a formal beta test program.\n\nThere may be bugs, incomplete features, incorrect documentation, or other discrepancies.\n\nDefault configuration requires 3 agent nodes each with: 0.5 CPU | 4096 MB MEM | 1 10240 MB Disk\n\nContact Mesosphere before deploying this beta candidate service. Product support is available to approved participants in the beta test program.",
// 	"postInstallNotes": "The DC/OS Apache Cassandra service is being installed!\n\n\tDocumentation: {{documentation-path}}\n\tIssues: {{issues-path}}",
// 	"postUninstallNotes": "The DC/OS Apache Cassandra service is being uninstalled.\n\nFor DC/OS versions from 1.10 no further action is required. For older DC/OS versions follow the instructions at {{documentation-path}}uninstall to remove any persistent state if required."
//   }
type Package struct {
	Name                  string   `json:"name"`
	PackagingVersion      string   `json:"packagingVersion"`
	MinDcosReleaseVersion string   `json:"minDcosReleaseVersion"`
	UpgradesFrom          []string `json:"upgradesFrom"`
	DowngradesTo          []string `json:"DowngradesTo"`
	Selected              bool     `json:"selected"`
	Framework             bool     `json:"framework"`
	Description           string   `json:"description"`
	Maintainer            string   `json:"maintainer"`
	Version               string   `json:"version"`
	Tags                  []string `json:"tags"`
	PreInstallNotes       string   `json:"preInstallNotes"`
	PostInstallNotes      string   `json:"postInstallNotes"`
	PostUninstallNotes    string   `json:"postUninstallNotes"`
}

// NewPackage constructs a new package with default values
func NewPackage() Package {
	p := Package{}
	p.PackagingVersion = DefaultPackagingVerison

	return p
}

func (c *Package) ParseYaml(data []byte) error {
	if err := yaml.Unmarshal(data, c); err != nil {
		return err
	}
	c.SetDefaults(GetDefaultPackage())
	return nil
}

func generateDefaultActionString(formatString string, names []string, defaultString string) string {
	for _, name := range names {
		if name != "" {
			return fmt.Sprintf(formatString, name)
		}
	}

	return defaultString
}

func (c *Package) SetDefaults(defaultPackage Package) {
	if c.Name == "" {
		c.Name = defaultPackage.Name
	}
	if c.Description == "" {
		if defaultPackage.Description != "" {
			c.Description = defaultPackage.Description
		} else {
			c.Description = c.Name
		}
	}
	if c.Maintainer == "" {
		c.Maintainer = defaultPackage.Maintainer
	}
	if c.PackagingVersion == "" {
		c.PackagingVersion = defaultPackage.PackagingVersion
	}
	if c.MinDcosReleaseVersion == "" {
		c.MinDcosReleaseVersion = defaultPackage.MinDcosReleaseVersion
	}
	if c.UpgradesFrom == nil {
		c.UpgradesFrom = defaultPackage.UpgradesFrom
	}
	if c.DowngradesTo == nil {
		c.DowngradesTo = defaultPackage.DowngradesTo
	}
	if c.Tags == nil {
		if defaultPackage.Tags != nil {
			c.Tags = defaultPackage.Tags
		} else if c.Name != "" {
			c.Tags = []string{c.Name}
		}
	}
	if c.Version == "" {
		c.Version = defaultPackage.Version
	}

	if c.PreInstallNotes == "" {
		c.PreInstallNotes = defaultPackage.PreInstallNotes
	}

	nameList := []string{c.Description, c.Name}

	if c.PostInstallNotes == "" {
		c.PostInstallNotes = generateDefaultActionString("The DC/OS %s service is being installed.", nameList, defaultPackage.PostInstallNotes)
	}

	if c.PostUninstallNotes == "" {
		c.PostUninstallNotes = generateDefaultActionString("The DC/OS %s service is being uninstalled.", nameList, defaultPackage.PostUninstallNotes)
	}

}

func GetDefaultPackage() Package {
	p := Package{}
	p.MinDcosReleaseVersion = DefaultMinDcosReleaseVersion
	p.PackagingVersion = DefaultPackagingVerison
	p.UpgradesFrom = []string{"{{upgrades-from}}"}
	p.DowngradesTo = []string{"{{downgrades-to}}"}
	p.Version = "{{package-version}}"
	p.Framework = true
	p.PostInstallNotes = "The DC/OS service is being installed."
	p.PostUninstallNotes = "The DC/OS service is being uninstalled."
	return p
}

func LoadYamlFile(filename string) (Package, error) {
	p := Package{}

	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		return p, err
	}
	err = p.ParseYaml(raw)
	if err != nil {
		return p, err
	}

	return p, nil
}

func LoadJsonFile(filename string) (Package, error) {
	p := Package{}
	file, err := os.Open(filename)
	if err != nil {
		return p, err
	}
	defer file.Close()

	dec := json.NewDecoder(file)

	err = dec.Decode(&p)
	if err != nil {
		return p, err
	}

	return p, err
}

func (p Package) WriteToJsonFile(outputDir string) error {

	outputFolder := fmt.Sprintf("%s/universe", outputDir)
	err := os.MkdirAll(outputFolder, 0777)
	if err != nil {
		return err
	}

	filename := fmt.Sprintf("%s/universe/package.json", outputDir)
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return err
	}
	defer file.Close()

	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")
	err = enc.Encode(p)
	if err != nil {
		return err
	}
	return nil
}
