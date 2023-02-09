package blog

import (
	"log"
	"net/http"
	"strings"
)

type server struct {
	renderer *PostRenderer
	postMap  map[string]Post
}

func NewServer(posts []Post) server {
	renderer, err := NewPostRenderer()
	if err != nil {
		log.Fatal(err)
	}

	postMap := make(map[string]Post)

	for _, post := range posts {
		postMap[post.SanitizeTitle()] = post
	}

	return server{renderer, postMap}
}

func (s *server) IndexHandler(w http.ResponseWriter, r *http.Request) {
	var posts []Post
	for _, post := range s.postMap {
		posts = append(posts, post)
	}

	s.renderer.RenderIndex(w, posts)
}

func (s *server) PostHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/post/")
	post := s.postMap[id]

	s.renderer.Render(w, post)
}
