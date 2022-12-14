package observers

import (
	"errors"
	"fmt"
	"sort"
)

type Event string

type (
	Listener interface {
		Handle(payload interface{})
		Listen() Event
		Priority() int
	}

	Dispatcher struct {
		Events map[Event][]Listener
	}
)

func (e Event) String() string {
	return string(e)
}

func (d *Dispatcher) Register(listeners ...Listener) {
	d.Events = make(map[Event][]Listener)
	sort.Slice(listeners, func(i, j int) bool {
		return listeners[i].Priority() > listeners[j].Priority()
	})

	for _, listener := range listeners {
		if _, ok := d.Events[listener.Listen()]; !ok {
			d.Events[listener.Listen()] = make([]Listener, 0, len(listeners))
		}

		d.Events[listener.Listen()] = append(d.Events[listener.Listen()], listener)
	}
}

func (d *Dispatcher) Dispatch(event Event, payload interface{}) error {
	if _, ok := d.Events[event]; !ok {
		return errors.New(fmt.Sprintf("event %s not registered", event.String()))
	}

	for _, listener := range d.Events[event] {
		listener.Handle(payload)
	}

	return nil
}
