package engine

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

const (
	// How far to scale up the total number of tiles. For example, if there were
	// going to be 18 'E' tiles, a TileScalingFactor of 10 means there would be
	// 180 tiles
	TileScalingFactor = 1000
)

type Bunch struct {
	rand  *rand.Rand
	mu    sync.Mutex
	tiles FreqList // List of tiles that are left
	count int      // Number of tiles left
}

func (b *Bunch) Tile() Letter {
	if b.count <= 0 {
		return '0' // return an invalid rune character so we know clearly something is wrong
	}
	b.mu.Lock()
	defer b.mu.Unlock()
	n := b.rand.Intn(b.count)
	for i, freq := range b.tiles {
		if freq > n {
			b.count--
			return b.tiles.DecIndex(i)
		}
		n -= freq
	}
	// should never happen
	log.Printf("reached bottom of tile(): %+v", b)
	return '0' // return an invalid rune character so we know clearly something is wrong
}

var distribution = map[int][]Letter{
	2:  []Letter{'j', 'k', 'q', 'x', 'z'},
	3:  []Letter{'b', 'c', 'f', 'h', 'm', 'p', 'v', 'w', 'y'},
	4:  []Letter{'g'},
	5:  []Letter{'l'},
	6:  []Letter{'d', 's', 'u'},
	8:  []Letter{'n'},
	9:  []Letter{'t', 'r'},
	11: []Letter{'o'},
	12: []Letter{'i'},
	13: []Letter{'a'},
	18: []Letter{'e'},
}

func NewBunch() *Bunch {
	b := &Bunch{
		rand: rand.New(rand.NewSource(time.Now().UTC().UnixNano())),
	}
	for freq, letters := range distribution {
		for _, letter := range letters {
			scaledFreq := freq * TileScalingFactor
			b.tiles.Set(letter, scaledFreq)
			b.count += scaledFreq
		}
	}
	return b
}
