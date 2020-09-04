package main

type Dictionary map[string]string

const (
	ErrNotFound      = DictionaryErr("could not find the word you were looking for")
	ErrWordExists    = DictionaryErr("cannot add word because it already exists")
	ErrWordNotExists = DictionaryErr("cannot update word because it doesn't exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(key string) (string, error) {
	definition, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(key string, value string) error {
	_, ok := d[key]
	if !ok {
		d[key] = value
	} else {
		return ErrWordExists
	}
	return nil
}

func (d Dictionary) Update(key string, value string) error {
	_, err := d.Search(key)
	switch err {
	case ErrNotFound:
		return ErrWordNotExists
	case nil:
		d[key] = value
		return nil
	default:
		return err
	}
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}
