package main

import "io/fs"

type Post struct {
	name string
}

func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}

	var posts []Post
	for _, filePath := range dir {
		posts = append(posts, Post{name: filePath.Name()})
	}

	return posts, nil
}
