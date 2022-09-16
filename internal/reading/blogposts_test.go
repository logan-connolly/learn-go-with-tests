package main

import (
	"testing"
	"testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte("hi")},
		"hello-world2.md": {Data: []byte("hola")},
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

	t.Run("test file name extracted", func(t *testing.T) {
		expected := []string{"hello world.md", "hello-world2.md"}
		for idx, want := range expected {
			got := posts[idx].name
			if want != got {
				t.Errorf("got %v posts, wanted %v posts", got, want)
			}
		}
	})
}
