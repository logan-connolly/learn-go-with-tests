package main

import (
	"io"
	"io/fs"
	"strings"
)

type Post struct {
	Title string
}

func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}

	var posts []Post
	for _, filePath := range dir {
		file, _ := fileSystem.Open(filePath.Name())
		post := makePostFromFile(file)
		posts = append(posts, post)
	}

	return posts, nil
}

func makePostFromFile(file io.Reader) Post {
	content, _ := io.ReadAll(file)
	title := extractTitle(content)
	return Post{Title: title}
}

func extractTitle(content []byte) string {
	return strings.TrimPrefix(string(content), "Title: ")
}
