package main

import (
	"fmt"
	"os"
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
