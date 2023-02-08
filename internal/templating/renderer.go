package blog

import (
	"embed"
	"html/template"
	"io"

	"github.com/gomarkdown/markdown"
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

func (r *PostRenderer) Render(w io.Writer, p Post) error {
	pvm := newPostViewModel(p, r)
	return r.tmpl.ExecuteTemplate(w, "blog.html.tmpl", pvm)
}

func (r *PostRenderer) RenderIndex(w io.Writer, posts []Post) error {
	return r.tmpl.ExecuteTemplate(w, "index.html.tmpl", posts)
}

type postViewModel struct {
	Post     Post
	HtmlBody template.HTML
}

func newPostViewModel(p Post, pr *PostRenderer) *postViewModel {
	htmlBody := template.HTML(markdown.ToHTML([]byte(p.Body), nil, nil))
	return &postViewModel{p, htmlBody}
}
