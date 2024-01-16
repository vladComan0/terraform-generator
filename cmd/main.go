package main

import (
	"flag"
	"log"
	"os"
	"text/template"
)

type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	templateCache map[string]*template.Template
}

type ec2Config struct {
	Name  string `yaml:"name"`
	Image string `yaml:"image"`
	Size  string `yaml:"size"`
}

type resource struct {
	EC2 ec2Config `yaml:"ec2"`
}

func main() {
	configFileName := flag.String("config", "config/config.yaml", "the config file to use")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	resource := resource{}
	if err := readYAMLConfig(*configFileName, &resource); err != nil {
		errorLog.Fatal(err)
	}

	templateCache, err := newTemplateCache()
	if err != nil {
		errorLog.Fatal(err)
	}

	app := &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		templateCache: templateCache,
	}

	if err = app.generate("ec2.tmpl", resource); err != nil {
		errorLog.Fatal(err)
	}
}
