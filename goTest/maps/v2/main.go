package main

type Dic map[string]string

func (d Dic) Search(word string) string {
	return d[word]
}
