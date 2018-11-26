package banana

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

func NewDictionary(r io.Reader) (Dictionary, error) {
	dict := &dictImpl{
		words: make(map[string]struct{}),
	}
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		dict.words[sc.Text()] = struct{}{}
	}

	if err := sc.Err(); err != nil {
		return nil, err
	}

	return dict, nli
}

func (d *dictImpl) HasWord(word string) bool {
	_, ok := d.words[word]
	return ok
}
