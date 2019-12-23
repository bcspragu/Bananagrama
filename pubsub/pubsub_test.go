package pubsub

import (
	"fmt"
	"testing"
)

func TestGlobal(t *testing.T) {
	ps := New()

	s, err := ps.Subscribe(Global)
	if err != nil {
		t.Fatalf("ps.Subscribe: %v", err)
	}

	for i := 0; i < 3; i++ {
		err := ps.Publish(Global, &Payload{
			Type:       PayloadTypePlayerMove,
			PlayerMove: &PlayerMove{Name: fmt.Sprintf("player %d", i), Word: "word"},
		})
		if err != nil {
			t.Fatalf("Publish: %v", err)
		}
	}

	// Read the three messages we're expecting.
	for i := 0; i < 3; i++ {
		msg, done := s.Next()
		if done {
			t.Fatalf("message %d was done, expected message", i)
		}
		compareMessage(t, msg, Global, fmt.Sprintf("player %d", i))
	}

	// Close the channel
	s.Close()

	if _, done := s.Next(); !done {
		t.Error("channel was closed, expected done")
	}
}

func TestMultipleChannels(t *testing.T) {
	ps := New()

	c1, err := ps.NewChannel("c1")
	if err != nil {
		t.Fatalf("failed to create c1: %v", err)
	}

	c2, err := ps.NewChannel("c2")
	if err != nil {
		t.Fatalf("failed to create c2: %v", err)
	}

	s1, err := ps.Subscribe(Global, c1)
	if err != nil {
		t.Fatalf("ps.Subscribe: %v", err)
	}
	defer s1.Close()

	s2, err := ps.Subscribe(Global, c2)
	if err != nil {
		t.Fatalf("ps.Subscribe: %v", err)
	}
	defer s2.Close()

	publish := func(ch Channel, name string) {
		if err != nil {
			return
		}

		err = ps.Publish(ch, &Payload{
			Type:       PayloadTypePlayerMove,
			PlayerMove: &PlayerMove{Name: name, Word: "word"},
		})
	}

	publish(Global, "global")
	publish(c1, "c1")
	publish(c2, "c2")

	if err != nil {
		t.Fatalf("Publish: %v", err)
	}

	confirmMessages := func(s *Subscriber, ch Channel, name string) {
		t.Helper()
		msg, done := s.Next()
		if done {
			t.Fatal("subscriber unexpectedly done")
		}
		compareMessage(t, msg, ch, name)
	}

	confirmMessages(s1, Global, "global")
	confirmMessages(s1, c1, "c1")

	confirmMessages(s2, Global, "global")
	confirmMessages(s2, c2, "c2")
}

func compareMessage(t *testing.T, got *Message, wantChannel Channel, wantName string) {
	t.Helper()

	if got.Channel != wantChannel {
		t.Errorf("got channel %q, want %q", got.Channel, wantChannel)
	}
	p := got.Payload
	if p == nil {
		t.Error("no payload given")
		return
	}
	if p.Type != PayloadTypePlayerMove {
		t.Errorf("got payload type %d, want %d", p.Type, PayloadTypePlayerMove)
		return
	}
	if p.PlayerMove.Name != wantName {
		t.Errorf("got name %q, want %q", p.PlayerMove.Name, wantName)
	}
	if p.PlayerMove.Word != "word" {
		t.Errorf("got word %q, want %q", p.PlayerMove.Word, "word")
	}
}
