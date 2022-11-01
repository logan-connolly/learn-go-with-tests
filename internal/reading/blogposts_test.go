package main

import (
	"reflect"
	"testing"
	"testing/fstest"
)

func assertPost(t *testing.T, got, want Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected %v got %v", want, got)
	}
}

func TestNewBlogPosts(t *testing.T) {
	const (
		postOne = `Title: Post 1
Description: Description 1`
		postTwo = `Title: Post 2
Description: Description 2`
	)

	fs := fstest.MapFS{
		"post-1.md": {Data: []byte(postOne)},
		"post-2.md": {Data: []byte(postTwo)},
	}

	posts, err := NewPostsFromFS(fs)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("test number of files found", func(t *testing.T) {
		if len(posts) != len(fs) {
			t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
		}
	})

	t.Run("test title extracted", func(t *testing.T) {
		expected := []Post{{"Post 1", "Description 1"}, {"Post 2", "Description 2"}}
		for idx, want := range expected {
			got := posts[idx]
			assertPost(t, got, want)
		}
	})
}
