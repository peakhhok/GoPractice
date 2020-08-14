package indic

import "errors"

//Dictionary type
type Dictionary map[string]string

var (
	errNotFound   = errors.New("Not Found")
	errWordExists = errors.New("That word already exists")
	errCantUpdate = errors.New("You can't Update")
)

//Search word
func (d Dictionary) Search(word string) (string, error) {
	value, exist := d[word]
	if exist {
		return value, nil

	}
	return "", errNotFound

}

//Add word
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}

	return nil
}

//Update text
func (d Dictionary) Update(word, definition string) error {

	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case errNotFound:
		return errCantUpdate
	}
	return nil

}

//Delete Word
func (d Dictionary) Delete(word string) error {

	_, err := d.Search(word)
	switch err {
	case nil:
		delete(d, word)
	case errWordExists:
		return errNotFound

	}
	return nil
}
