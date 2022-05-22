package main

import "testing"

var TEST_WORD = "test"
var TEST_DEFINITION = "this is just a test"

func TestSearch(t *testing.T) {

	t.Run("known word", func(t *testing.T) {
		d := Dictionary{TEST_WORD: TEST_DEFINITION}

		assertDefinition(t, d, TEST_WORD, TEST_DEFINITION)
	})

	t.Run("unknown word", func(t *testing.T) {
		d := Dictionary{}

		_, err := d.Search("unknown")
		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assertError(t, err, ErrorWordNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("add word", func(t *testing.T) {
		d := Dictionary{}
		d.Add(TEST_WORD, TEST_DEFINITION)

		assertDefinition(t, d, TEST_WORD, TEST_DEFINITION)
	})

	t.Run("add word that already exists", func(t *testing.T) {
		d := Dictionary{TEST_WORD: TEST_DEFINITION}
		err := d.Add(TEST_WORD, TEST_DEFINITION)

		if err == nil {
			t.Error("Expected to get an error")
		}

		assertError(t, err, ErrorWordExists)
	})
}

func assertDefinition(t *testing.T, d Dictionary, word, definition string) {
	t.Helper()

	got, err := d.Search(word)

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if definition != got {
		t.Errorf("got %q want %q", definition, got)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
