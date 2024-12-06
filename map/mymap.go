package mymap

type Dictionary map[string]string

type DictionaryErr string

func (d DictionaryErr) Error() string {
	return string(d)
}

var ErrWordNotFound = DictionaryErr("word has not been found")
var ErrKeyAlreadyExisting = DictionaryErr("key has already in the dictionary")
var ErrKeyWorldDoesntExist = DictionaryErr("the world doesn't exist")

func (d Dictionary) Search(word string) (string, error) {
	result, ok := d[word]
	if ok {
		return result, nil
	}
	return "", ErrWordNotFound

}

func (d Dictionary) Add(key, value string) error {

	_, err := d.Search(key)

	switch err {
	case ErrWordNotFound:
		// this means that the key was not found, we can assign the key.
		d[key] = value
	case nil:
		// this means that the key is already in the dictionary, we should assign the err.
		return ErrKeyAlreadyExisting
	default:
		return nil
	}

	return nil

}

func (d Dictionary) Update(key, value string) error {

	_, err := d.Search(key)

	switch err {
	case ErrWordNotFound:
		return ErrWordNotFound
	case nil:
		d[key] = value
		return nil
	default:
		return err
	}

}

func (d Dictionary) Delete(key string) error {

	_, err := d.Search(key)

	switch err {
	case ErrWordNotFound:
		return ErrKeyWorldDoesntExist
	case nil:
		delete(d, key)
	default:
		return err
	}

	return nil

}
