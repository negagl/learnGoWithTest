package maps

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")
var ErrAlreadyExists = errors.New("word already exists in dictionary")

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	if err == nil {
		return ErrAlreadyExists
	}

	d[word] = definition

	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	if err == nil {
		d[word] = definition
		return nil
	}

	return ErrNotFound
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
