package main

import (
	"log"
	"net/http"
	"strings"

	blog "github.com/logan-connolly/learn-go-with-tests/internal/templating"
)

var samplePosts = []blog.Post{
	{
		Title:       "Learn go with tests",
		Body:        "**This** is a _great_ first post.",
		Description: "Post numero uno",
		Tags:        []string{"go", "tdd"},
	},
	{
		Title:       "Start with NixOS",
		Body:        "**This** is a _great_ second post.",
		Description: "Post numero dos",
		Tags:        []string{"nixos", "infra"},
	},
}

type server struct {
	renderer *blog.PostRenderer
	postMap  map[string]blog.Post
}

func NewServer() server {
	renderer, err := blog.NewPostRenderer()
	if err != nil {
		log.Fatal(err)
	}

	postMap := make(map[string]blog.Post)

	for _, post := range samplePosts {
		postMap[post.SanitizeTitle()] = post
	}

	return server{renderer, postMap}
}

func (s *server) indexHandler(w http.ResponseWriter, r *http.Request) {
	var posts []blog.Post
	for _, post := range s.postMap {
		posts = append(posts, post)
	}

	s.renderer.RenderIndex(w, posts)
}

func (s *server) postHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/post/")
	post := s.postMap[id]

	s.renderer.Render(w, post)
}

func main() {
	s := NewServer()

	http.HandleFunc("/", s.indexHandler)
	http.HandleFunc("/post/", s.postHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
