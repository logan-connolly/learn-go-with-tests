package main

import (
	"bufio"
	"bytes"
	"fmt"
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
	Body        string
}

func (p *Post) SanitizeTitle() string {
	return strings.ToLower(strings.Replace(p.Title, " ", "-", -1))
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
	return newPost(bufio.NewScanner(file))
}

func parseLine(scanner *bufio.Scanner, seperator string) string {
	scanner.Scan()
	return strings.TrimPrefix(scanner.Text(), seperator)
}

func parseBody(scanner *bufio.Scanner) string {
	scanner.Scan()
	buf := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}
	return strings.TrimSuffix(buf.String(), "\n")
}

func newPost(scanner *bufio.Scanner) (Post, error) {
	post := Post{
		Title:       parseLine(scanner, titleSeperator),
		Description: parseLine(scanner, descriptionSeperator),
		Tags:        strings.Split(parseLine(scanner, tagsSeperator), ", "),
		Body:        parseBody(scanner),
	}
	return post, nil
}
