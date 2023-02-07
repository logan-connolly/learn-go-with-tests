package main

import (
	"embed"
	"html/template"
	"io"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

//go:embed "templates/*"
var blogTemplates embed.FS

type PostRenderer struct {
	tmpl     *template.Template
	mdParser *parser.Parser
}

func NewPostRenderer() (*PostRenderer, error) {
	tmpl, err := template.ParseFS(blogTemplates, "templates/*.tmpl")
	if err != nil {
		return nil, err
	}

	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	parser := parser.NewWithExtensions(extensions)

	return &PostRenderer{tmpl, parser}, nil
}

func (pr *PostRenderer) Render(w io.Writer, p Post) error {
	pvm := newPostViewModel(p, pr)
	return pr.tmpl.ExecuteTemplate(w, "blog.html.tmpl", pvm)
}

func (pr *PostRenderer) RenderIndex(w io.Writer, posts []Post) error {
	return pr.tmpl.ExecuteTemplate(w, "index.html.tmpl", posts)
}

type postViewModel struct {
	Post     Post
	HtmlBody template.HTML
}

func newPostViewModel(p Post, pr *PostRenderer) postViewModel {
	htmlBody := markdown.ToHTML([]byte(p.Body), pr.mdParser, nil)
	return postViewModel{p, template.HTML(htmlBody)}
}
