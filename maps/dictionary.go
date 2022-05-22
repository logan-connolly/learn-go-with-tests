package main

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

var (
	ErrorWordNotFound = DictionaryErr("Word not found.")
	ErrorWordExists   = DictionaryErr("Word already exists.")
)

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrorWordNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	if _, err := d.Search(word); err == nil {
		return ErrorWordExists
	}

	d[word] = definition
	return nil
}
