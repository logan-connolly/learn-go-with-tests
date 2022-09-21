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
		post, err := getPost(fileSystem, filePath.Name())
		if err != nil {
			return []Post{}, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func getPost(fileSystem fs.FS, fileName string) (Post, error) {
	file, err := fileSystem.Open(fileName)
	if err != nil {
		return Post{}, err
	}
	defer file.Close()
	return newPost(file)
}

func newPost(file io.Reader) (Post, error) {
	content, err := io.ReadAll(file)
	if err != nil {
		return Post{}, err
	}
	return extractPostInfo(content), nil
}

func extractPostInfo(content []byte) Post {
	title := strings.TrimPrefix(string(content), "Title: ")
	return Post{title}
}
