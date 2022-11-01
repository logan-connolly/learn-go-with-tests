package main

import (
	"bufio"
	"io"
	"io/fs"
	"strings"
)

const (
	titleSeperator       = "Title: "
	descriptionSeperator = "Description: "
	tagsSeperator        = "Tags: "
)

type Post struct {
	Title       string
	Description string
	Tags        []string
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
	scanner := bufio.NewScanner(file)

	parseLine := func(seperator string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), seperator)
	}

	return Post{
		Title:       parseLine(titleSeperator),
		Description: parseLine(descriptionSeperator),
		Tags:        strings.Split(parseLine(tagsSeperator), ", "),
	}, nil
}
