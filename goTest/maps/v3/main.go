package main

import "errors"

var ErrNotFound = errors.New("could not find the word you are looking for")

type Dic map[string]string

func (d Dic) Search(word string) (string, error) {

	definition, ok := d[word]
	if !ok {
		return "", errors.New("could not find the word you were looking for")
	}
	return definition, nil
}
