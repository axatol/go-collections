package ds

import (
	"sync"
)

type subscriber[T any] struct {
	id       string
	delivery chan<- T
}

// thread-safe distributes events from an incoming channel into subscriber
// channels will be discarded, i.e. you will lose events if your subscriber
// can't keep up
type Fanout[T any] struct {
	lock        sync.RWMutex
	subscribers map[string]subscriber[T]
	incoming    <-chan T
}

// create a new Fanout, creates a goroutine to broadcast when an incoming
// event is received
func NewFanout[T any](incoming <-chan T) *Fanout[T] {
	f := Fanout[T]{incoming: incoming}

	go func() {
		// continuously broadcast events
		for event := range incoming {
			f.Broadcast(event)
		}

		// release subscribers once incoming is closed
		for _, subscriber := range f.subscribers {
			close(subscriber.delivery)
		}
	}()

	return &f
}

// add channel as a target for event delivery
func (f *Fanout[T]) Subscribe(id string, delivery chan<- T) {
	f.lock.Lock()
	defer f.lock.Unlock()

	f.subscribers[id] = subscriber[T]{id, delivery}
}

// remove channel from deliveries
func (f *Fanout[T]) Unsubscribe(id string) {
	f.lock.Lock()
	defer f.lock.Unlock()

	delete(f.subscribers, id)
}

// send an event to all subscribers
func (f *Fanout[T]) Broadcast(payload T) {
	f.lock.RLock()
	defer f.lock.RUnlock()

	for _, subscriber := range f.subscribers {
		select {
		case subscriber.delivery <- payload:
		default:
		}
	}
}
