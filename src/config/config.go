package config

import (
	"encoding/json"

	"github.com/alecthomas/jsonschema"
)

type tName struct {
	Type string
}

type Service struct {
	Name                       string `json:"name" jsonschema:"required"`
	User                       string `json:"user" jsonschema:"required"`
	ServiceAccount             string `json:"service_account"`
	ServiceAccountSecret       string `json:"service_account_secret"`
	VirtualNetworksEnabled     bool   `json:"virtual_network_enabled"`
	VirtualNetworkName         string `json:"virtual_network_name"`
	VirtualNetworkPluginLabels string `json:"virtual_network_plugin_labels"`
	LogLevel                   string `json:"log_level"`
}

func GenerateJsonSchema() (string, error) {
	s := Service{}

	r := jsonschema.Reflector{}
	r.RequiredFromJSONSchemaTags = true
	r.ExpandedStruct = true

	schema := r.Reflect(&s)
	schema.Type.Properties["name"].Title = "Service name"
	schema.Type.Properties["name"].Description = "The name of the service instance"
	schema.Type.Properties["name"].Default = "serviceName"

	schema.Type.Properties["user"].Title = "User"
	schema.Type.Properties["user"].Description = "The user that the service will run as."
	schema.Type.Properties["user"].Default = "nobody"

	schema.Type.Properties["service_account"].Description = "The service account for DC/OS service authentication. This is typically left empty to use the default unless service authentication is needed. The value given here is passed as the principal of Mesos framework."
	schema.Type.Properties["service_account"].Default = ""

	schema.Type.Properties["service_account_secret"].Title = "Credential secret name (optional)"
	schema.Type.Properties["service_account_secret"].Description = "Name of the Secret Store credentials to use for DC/OS service authentication. This should be left empty unless service authentication is needed."
	schema.Type.Properties["service_account_secret"].Default = ""

	schema.Type.Properties["virtual_network_enabled"].Description = "Enable virtual networking"
	schema.Type.Properties["virtual_network_enabled"].Default = false

	schema.Type.Properties["virtual_network_name"].Description = "The name of the virtual network to join"
	schema.Type.Properties["virtual_network_name"].Default = "dcos"

	schema.Type.Properties["virtual_network_plugin_labels"].Description = "Labels to pass to the virtual network plugin. Comma-separated key:value pairs. For example: k_0:v_0,k_1:v_1,...,k_n:v_n"
	schema.Type.Properties["virtual_network_plugin_labels"].Default = ""

	schema.Type.Properties["log_level"].Description = "The log level for the DC/OS service."
	schema.Type.Properties["log_level"].Enum = []interface{}{
		"OFF",
		"FATAL",
		"ERROR",
		"WARN",
		"INFO",
		"DEBUG",
		"TRACE",
		"ALL",
	}
	schema.Type.Properties["log_level"].Default = "INFO"

	raw, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		return "", nil
	}
	return string(raw), nil
}
