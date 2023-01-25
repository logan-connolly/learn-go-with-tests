package main

import (
	"bytes"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
)

func TestRender(t *testing.T) {
	samplePost := Post{
		Title:       "hello world",
		Body:        "This is a post",
		Description: "This is a description",
		Tags:        []string{"go", "tdd"},
	}

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}

		if err := Render(&buf, samplePost); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
}
