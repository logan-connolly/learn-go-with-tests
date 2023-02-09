package main

import (
	"log"
	"net/http"

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

func main() {
	s := blog.NewServer(samplePosts)

	http.HandleFunc("/", s.IndexHandler)
	http.HandleFunc("/post/", s.PostHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
