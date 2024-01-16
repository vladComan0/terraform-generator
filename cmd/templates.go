package main

import (
	"io/fs"
	"path/filepath"
	"text/template"

	"github.com/vladComan0/terraform-generator/config"
)

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	files, err := fs.Glob(config.Files, "files/*.tmpl")
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		name := filepath.Base(file)

		patterns := []string{
			"base.tmpl",
			file,
		}

		ts, err := template.New(name).ParseFS(config.Files, patterns...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
