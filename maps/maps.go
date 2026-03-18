package maps

const (
	ErrNotFound   = DictionaryErr("definition with that name could not be found")
	ErrWordExists = DictionaryErr("definition with this word already exists in dictionary")
)

type DictionaryErr string
type Dictionary map[string]string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	found, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return found, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}
