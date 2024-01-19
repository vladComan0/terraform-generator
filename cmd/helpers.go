package main

import (
	"fmt"
	"io"
	"os"

	"github.com/vladComan0/terraform-generator/internal/terraform"
	"gopkg.in/yaml.v3"
)

func (app *application) generateHCL(file string, data any, destination io.Writer) error {
	ts, ok := app.templateCache[file]
	if !ok {
		return fmt.Errorf("the template %s does not exist", file)
	}

	if err := ts.ExecuteTemplate(destination, "base", data); err != nil {
		return err
	}

	return nil
}

// func readYAMLConfig(filename string, out terraform.Resource) error {
// 	yamlConfig, err := os.ReadFile(filename)
// 	if err != nil {
// 		return err
// 	}
// 	return yaml.Unmarshal(yamlConfig, out)
// }

func readYAMLConfig(filename string) ([]terraform.ResourceConfig, error) {
	var config struct {
		Resources []terraform.ResourceConfig `yaml:"resources"`
	}

	yamlConfig, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(yamlConfig, &config)
	if err != nil {
		return nil, err
	}

	return config.Resources, nil
}

// func readResourceTypeFromConfig(config string) string {
// 	// Handle reading the resource type from the config file
// 	return "EC2"
// }

// func createResourceConfig(resourceType string) (terraform.Resource, error) {
// 	switch resourceType {
// 	case "EC2":
// 		return &terraform.EC2Config{}, nil
// 	// Handle other resource types
// 	default:
// 		return nil, fmt.Errorf("unsupported resource type")
// 	}
// }
