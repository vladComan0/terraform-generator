package main

import (
	"log"
	"os"
	"text/template"
)

type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	templateCache map[string]*template.Template
}

type resource struct {
	Name  string
	Image string
	Size  string
}

func main() {
	resource := resource{
		Name:  "ec2-instance",
		Image: "ami-test",
		Size:  "t2.micro",
	}

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

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
