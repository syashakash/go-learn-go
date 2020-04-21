package maps

const (
	ErrNotFound = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
	)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}


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

func (dict Dictionary) Update(word, definition string) {
	dict[word] = definition
}