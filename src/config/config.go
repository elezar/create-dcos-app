package config

import (
	"encoding/json"
	"fmt"
	"os"

	generator "github.com/elezar/go-json-schema-generator"
)

type Config struct {
	Service Service `json:"service"`
}

type Service struct {
	Name                       string `json:"name" required:"true" jsonschema:"required"`
	User                       string `json:"user" required:"true" jsonschema:"required"`
	ServiceAccount             string `json:"service_account"`
	ServiceAccountSecret       string `json:"service_account_secret"`
	VirtualNetworksEnabled     bool   `json:"virtual_network_enabled"`
	VirtualNetworkName         string `json:"virtual_network_name"`
	VirtualNetworkPluginLabels string `json:"virtual_network_plugin_labels"`
	LogLevel                   string `json:"log_level"`
}

func (c *Config) SetDefaults(defaultServiceName string) {
	c.Service.Name = defaultServiceName
	c.Service.User = "nobody"
	c.Service.VirtualNetworkName = "docs"
	c.Service.LogLevel = "INFO"
}

func GenerateJsonSchema(defaultConfig Config) *generator.Document {

	d := generator.Document{}
	g := d.Read(&Config{})

	g.Properties["service"].Properties["name"].Description = "The name of the service instance"

	g.Properties["service"].Properties["name"].Title = "Service name"
	g.Properties["service"].Properties["name"].Description = "The name of the service instance"
	g.Properties["service"].Properties["name"].Default = defaultConfig.Service.Name

	g.Properties["service"].Properties["user"].Title = "User"
	g.Properties["service"].Properties["user"].Description = "The user that the service will run as."
	g.Properties["service"].Properties["user"].Default = defaultConfig.Service.User

	g.Properties["service"].Properties["service_account"].Description = "The service account for DC/OS service authentication. This is typically left empty to use the default unless service authentication is needed. The value given here is passed as the principal of Mesos framework."
	g.Properties["service"].Properties["service_account"].Default = defaultConfig.Service.ServiceAccount

	g.Properties["service"].Properties["service_account_secret"].Title = "Credential secret name (optional)"
	g.Properties["service"].Properties["service_account_secret"].Description = "Name of the Secret Store credentials to use for DC/OS service authentication. This should be left empty unless service authentication is needed."
	g.Properties["service"].Properties["service_account_secret"].Default = defaultConfig.Service.ServiceAccountSecret

	g.Properties["service"].Properties["virtual_network_enabled"].Description = "Enable virtual networking"
	g.Properties["service"].Properties["virtual_network_enabled"].Default = defaultConfig.Service.VirtualNetworksEnabled

	g.Properties["service"].Properties["virtual_network_name"].Description = "The name of the virtual network to join"
	g.Properties["service"].Properties["virtual_network_name"].Default = defaultConfig.Service.VirtualNetworkName

	g.Properties["service"].Properties["virtual_network_plugin_labels"].Description = "Labels to pass to the virtual network plugin. Comma-separated key:value pairs. For example: k_0:v_0,k_1:v_1,...,k_n:v_n"
	g.Properties["service"].Properties["virtual_network_plugin_labels"].Default = defaultConfig.Service.VirtualNetworkPluginLabels

	g.Properties["service"].Properties["log_level"].Description = "The log level for the DC/OS service."
	g.Properties["service"].Properties["log_level"].Enum = []string{
		"OFF",
		"FATAL",
		"ERROR",
		"WARN",
		"INFO",
		"DEBUG",
		"TRACE",
		"ALL",
	}
	g.Properties["service"].Properties["log_level"].Default = defaultConfig.Service.LogLevel

	return g
}

func (c Config) WriteToJsonFile(outputDir string) error {

	outputFolder := fmt.Sprintf("%s/universe", outputDir)
	err := os.MkdirAll(outputFolder, 0777)
	if err != nil {
		return err
	}

	filename := fmt.Sprintf("%s/universe/config.json", outputDir)
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return err
	}
	defer file.Close()

	g := GenerateJsonSchema(c)

	json, _ := json.MarshalIndent(g, "", "  ")
	_, err = file.Write(json)

	return err
}
