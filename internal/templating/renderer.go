package main

import (
	"embed"
	"html/template"
	"io"
)

//go:embed "templates/*"
var blogTemplates embed.FS

type PostRenderer struct {
	tmpl *template.Template
}

func NewPostRenderer() (*PostRenderer, error) {
	tmpl, err := template.ParseFS(blogTemplates, "templates/*.tmpl")
	if err != nil {
		return nil, err
	}

	return &PostRenderer{tmpl}, nil
}

func (pr *PostRenderer) Render(w io.Writer, p Post) error {
	if err := pr.tmpl.Execute(w, p); err != nil {
		return err
	}

	return nil
}
