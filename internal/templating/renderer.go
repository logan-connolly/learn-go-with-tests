package main

import (
	"embed"
	"html/template"
	"io"
)

//go:embed "templates/*"
var blogTemplates embed.FS

// Render takes a io.Writer and Post and inserts it into
// the defined blog template.
func Render(w io.Writer, p Post) error {
	tmpl, err := template.ParseFS(blogTemplates, "templates/*.gohtml")
	if err != nil {
		return err
	}

	if err := tmpl.Execute(w, p); err != nil {
		return err
	}

	return nil
}
