package engine

import (
	"bufio"
	"io"
)

type Dictionary interface {
	HasWord(word string) bool
}

type dictImpl struct {
	words map[string]struct{}
}

func NewDictionary(r io.Reader) Dictionary {
	dict := &dictImpl{
		words: make(map[string]struct{}),
	}
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		dict.words[sc.Text()] = struct{}{}
	}
	return dict
}

func (d *dictImpl) HasWord(word string) bool {
	_, ok := d.words[word]
	return ok
}
