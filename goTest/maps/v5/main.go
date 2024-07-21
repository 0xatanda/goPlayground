package main

import "errors"

type Dic map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")

func (d Dic) Search(word string) (string, error) {
	def, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return def, nil
}

func (d Dic) Add(word, def string) {
	d[word] = def
}
