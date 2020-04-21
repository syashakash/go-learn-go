package maps

import "errors"

var (
	ErrNotFound = errors.New("could not find the word you were looking for")
	ErrWordExists = errors.New("")
	)



type Dictionary map[string]string

func (dict Dictionary) Search(word string) (string, error) {
	definition, ok := dict[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (dict Dictionary) Add(word, definition string) error {
	_, err := dict.Search(word)
	switch err {
	case ErrNotFound:
		dict[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}
