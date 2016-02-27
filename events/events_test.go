package events

import (
	"testing"

	"github.com/samalba/dockerclient"
)

func TestEventHandler(t *testing.T) {
	tChan := make(chan *dockerclient.Event)

	errChan := make(chan error)

	go func() {
		for err := range errChan {
			t.Fatal(err)
		}
	}()

	h, err := NewEventHandler(tChan)
	if err != nil {
		t.Fatal(err)
	}

	testEvent := &dockerclient.Event{
		Type: "testevent",
	}

	go h.Handle(testEvent, errChan, nil)

	v := <-tChan

	if v.Type != "testevent" {
		t.Fatalf("unexpected event type %s", v.Type)
	}
}
