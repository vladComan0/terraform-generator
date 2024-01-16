package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func (app *application) generate(file string, data interface{}) error {
	ts, ok := app.templateCache[file]
	if !ok {
		return fmt.Errorf("the template %s does not exist", file)
	}

	if err := ts.ExecuteTemplate(os.Stdout, "base", data); err != nil {
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
