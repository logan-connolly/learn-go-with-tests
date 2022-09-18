package main

import (
	"testing"
	"testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"post-1.md": {Data: []byte("Title: Post 1")},
		"post-2.md": {Data: []byte("Title: Post 2")},
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
		expected := []string{"Post 1", "Post 2"}
		for idx, want := range expected {
			got := posts[idx].Title
			if want != got {
				t.Errorf("got %s posts, wanted %s posts", got, want)
			}
		}
	})
}
