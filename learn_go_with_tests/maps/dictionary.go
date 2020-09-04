package main

type Dictionary map[string]string

const (
	ErrorNotFound = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(key string) (string, error) {
	definition, ok := d[key]
	if !ok {
		return "", ErrorNotFound
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

func (d Dictionary) Update(key string, value string) {
	d[key] = value
}
