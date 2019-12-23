package pubsub

import (
	"fmt"
	"sync"
)

type Channel string

const (
	Global = Channel("GLOBAL")
)

type PubSub struct {
	mu            sync.RWMutex
	subscriptions map[Channel][]*subscription
	latestID      int
}

type Message struct {
	Channel Channel
	Payload *Payload
}

func New() *PubSub {
	return &PubSub{
		subscriptions: map[Channel][]*subscription{
			Global: []*subscription{},
		},
	}
}

func (ps *PubSub) NewChannel(name string) (Channel, error) {
	ch := Channel(name)

	ps.mu.Lock()
	defer ps.mu.Unlock()

	if _, ok := ps.subscriptions[ch]; ok {
		return "", fmt.Errorf("channel %q already exists", name)
	}
	ps.subscriptions[ch] = []*subscription{}
	return ch, nil
}

func (ps *PubSub) Publish(c Channel, p *Payload) error {
	ps.mu.RLock()
	defer ps.mu.RUnlock()
	subs, ok := ps.subscriptions[c]
	if !ok {
		return fmt.Errorf("channel %q doesn't exist", c)
	}
	for _, sub := range subs {
		sub.in <- &Message{
			Channel: c,
			Payload: p,
		}
	}
	return nil
}

func (ps *PubSub) Subscribe(channels ...Channel) (*Subscriber, error) {
	in, out, done := make(chan *Message), make(chan *Message), make(chan struct{})

	ps.mu.Lock()
	defer ps.mu.Unlock()

	subID := ps.nextSubscriptionID()

	// First, check all the requested channels exist.
	for _, c := range channels {
		if _, ok := ps.subscriptions[c]; !ok {
			return nil, fmt.Errorf("channel %q doesn't exist", c)
		}
	}

	sub := &subscription{
		id:   subID,
		in:   in,
		out:  out,
		done: done,
	}
	go sub.listen()

	// Then, add the subscriptions
	for _, c := range channels {
		ps.subscriptions[c] = append(ps.subscriptions[c], sub)
	}

	return &Subscriber{
		msgs:    out,
		done:    make(chan struct{}),
		closeFn: ps.removeSubscription(subID, channels),
	}, nil
}

func (ps *PubSub) removeSubscription(id subscriptionID, channels []Channel) func() {
	return func() {
		ps.mu.Lock()
		defer ps.mu.Unlock()

		for _, c := range channels {
			for i, sub := range ps.subscriptions[c] {
				if sub.id == id {
					ss := ps.subscriptions[c]
					if i < len(ss)-1 {
						copy(ss[i:], ss[i+1:])
					}
					ss[len(ss)-1] = nil
					ss = ss[:len(ss)-1]
					ps.subscriptions[c] = ss
					break
				}
			}
		}
	}
}

// nextSubscriptionID returns the next available subscription ID. This should only
// be called from functions that already hold the mutex.
func (ps *PubSub) nextSubscriptionID() subscriptionID {
	ps.latestID++
	return subscriptionID(ps.latestID)
}

type subscriptionID int

type subscription struct {
	id   subscriptionID
	in   chan *Message
	out  chan *Message
	done chan struct{}

	mu       sync.Mutex
	draining bool
	queue    []*Message
}

func (s *subscription) drain() {
	var msg *Message
outer:
	for {
		s.mu.Lock()
		if len(s.queue) == 0 {
			s.mu.Unlock()
			break
		}
		msg, s.queue = s.queue[0], s.queue[1:]
		s.mu.Unlock()

		select {
		case s.out <- msg:
			// Message was sent.
		case <-s.done:
			// We're done.
			break outer
		}
	}

	s.mu.Lock()
	s.draining = false
	s.mu.Unlock()
}

func (s *subscription) listen() {
	for {
		select {
		case msg := <-s.in:
			s.mu.Lock()
			s.queue = append(s.queue, msg)
			if !s.draining {
				s.draining = true
				go s.drain()
			}
			s.mu.Unlock()
		case <-s.done:
			return
		}
	}
}

type Subscriber struct {
	msgs    chan *Message
	done    chan struct{}
	closeFn func()
}

// Next will block until a message is available, or the channel is closed.
func (s *Subscriber) Next() (*Message, bool) {
	select {
	case m := <-s.msgs:
		return m, false
	case <-s.done:
		return nil, true
	}
}

func (s *Subscriber) Close() {
	s.closeFn()
	close(s.done)
}
