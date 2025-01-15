package config

import (
	"bytes"
	"html/template"
	"io"
	"log"
)

var Templates globalTemplates

type globalTemplates struct {
	templates *template.Template
}

func (gt *globalTemplates) Exec(wr io.Writer, templateName string, data any) error {
	var buffer bytes.Buffer
	err := gt.templates.ExecuteTemplate(&buffer, templateName, data)
	if err != nil {
		return err
	}
	_, err = buffer.WriteTo(wr)

	return err
}

func InitTemplates(pattern string) {
	if Templates.templates == nil {
		Templates.templates = template.New("forum")
	}
	_, err := Templates.templates.ParseGlob(pattern)
	if err != nil {
		log.Fatal(err)
	}
}
