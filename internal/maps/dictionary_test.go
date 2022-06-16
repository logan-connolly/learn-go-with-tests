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
		err := d.Add(TEST_WORD, TEST_DEFINITION)

		assertError(t, err, nil)
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

func TestUpdate(t *testing.T) {
	t.Run("update word", func(t *testing.T) {
		new_test_definition := "My new definition"
		d := Dictionary{TEST_WORD: TEST_DEFINITION}
		err := d.Update(TEST_WORD, new_test_definition)

		assertError(t, err, nil)
		assertDefinition(t, d, TEST_WORD, new_test_definition)
	})

	t.Run("target word does not exist", func(t *testing.T) {
		d := Dictionary{TEST_WORD: TEST_DEFINITION}
		err := d.Update("badword", "this will not be updated")

		if err == nil {
			t.Error("Expected to get an error")
		}

		assertError(t, err, ErrorWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete word", func(t *testing.T) {
		d := Dictionary{TEST_WORD: TEST_DEFINITION}
		d.Delete(TEST_WORD)

		_, err := d.Search(TEST_WORD)
		assertError(t, err, ErrorWordNotFound)
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
