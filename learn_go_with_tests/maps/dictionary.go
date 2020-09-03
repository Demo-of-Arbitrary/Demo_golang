package main

type Dictionary map[string]string

func (d Dictionary) Search(key string) string {
	return d[key]
}

func Search(dictionary map[string]string, query string) string {
	return dictionary[query]
}
