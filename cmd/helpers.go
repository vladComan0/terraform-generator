package main

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

func (app *application) generateHCL(file string, data interface{}, destination io.Writer) error {
	ts, ok := app.templateCache[file]
	if !ok {
		return fmt.Errorf("the template %s does not exist", file)
	}

	if err := ts.ExecuteTemplate(destination, "base", data); err != nil {
		return err
	}

	return nil
}

func readYAMLConfig(filename string, out interface{}) error {
	yamlConfig, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(yamlConfig, out)
}
