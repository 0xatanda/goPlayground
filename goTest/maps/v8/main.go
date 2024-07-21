package main

type Dic map[string]string

type DicErr string

const (
	ErrNotFound          = DicErr("could not find what you were looking for")
	ErrWordExists        = DicErr("cannot add word because it already exists")
	ErrWordDoesNotExists = DicErr("cannot update word because it does not exist")
)

func (e DicErr) Error() string {
	return string(e)
}

func (d Dic) Search(word string) (string, error) {
	def, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return def, nil
}

func (d Dic) Add(word, def string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = def
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (d Dic) Update(word, def string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExists
	case nil:
		d[word] = def
	default:
		return err
	}
	return nil
}

func (d Dic) Delete(word string) {
	delete(d, word)
}
