package ds

import (
	"sync"
)

// an element of the map
type AsyncMapItem[T any] struct {
	ID        string `json:"id"`
	Completed bool   `json:"completed"`
	Failed    bool   `json:"failed"`
	Data      T      `json:"data"`
}

// type of possible actions that may be emitted
type AsyncMapEventAction string

const (
	AddedEventAction     AsyncMapEventAction = "added"
	CompletedEventAction AsyncMapEventAction = "completed"
	FailedEventAction    AsyncMapEventAction = "failed"
	RemovedEventAction   AsyncMapEventAction = "removed"
)

// the event emitted to subscribers
type AsyncMapEvent[T any] struct {
	Action AsyncMapEventAction `json:"action"`
	Item   AsyncMapItem[T]     `json:"item"`
}

// a thread-safe map with subscribable events
type AsyncMap[T any] struct {
	mutex  sync.RWMutex
	items  map[string]AsyncMapItem[T]
	events chan AsyncMapEvent[T]
}

func NewAsyncMap[T any](initial ...map[string]AsyncMapItem[T]) *AsyncMap[T] {
	items := map[string]AsyncMapItem[T]{}
	if len(initial) == 1 {
		items = initial[0]
	}

	return &AsyncMap[T]{
		mutex:  sync.RWMutex{},
		items:  items,
		events: make(chan AsyncMapEvent[T], 1),
	}
}

func (q *AsyncMap[T]) emit(item AsyncMapEvent[T]) {
	select {
	case q.events <- item:
	default:
	}
}

// returns the event channel
//
// use Fanout if you want multiple subscribers
func (q *AsyncMap[T]) Subscribe() <-chan AsyncMapEvent[T] {
	return q.events
}

// adds an element, no effect if an element already exists
func (q *AsyncMap[T]) Add(id string, data T) {
	if q.Get(id) != nil {
		return
	}

	q.Set(id, data)
}

// adds an element, overwriting if element already exists
func (q *AsyncMap[T]) Set(id string, data T) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	item := AsyncMapItem[T]{
		ID:        id,
		Completed: false,
		Failed:    false,
		Data:      data,
	}

	q.items[id] = item
	q.emit(AsyncMapEvent[T]{AddedEventAction, item})
}

// sets the item as failed
func (q *AsyncMap[T]) SetFailed(id string) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	item, ok := q.items[id]
	if !ok {
		return
	}

	item.Failed = true
	q.items[id] = item
	q.emit(AsyncMapEvent[T]{FailedEventAction, item})
}

// sets the item as completed
func (q *AsyncMap[T]) SetCompleted(id string) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	item, ok := q.items[id]
	if !ok {
		return
	}

	item.Completed = true
	q.items[id] = item
	q.emit(AsyncMapEvent[T]{CompletedEventAction, item})
}

// remove the item
func (q *AsyncMap[T]) Remove(id string) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	item, ok := q.items[id]
	if !ok {
		return
	}

	q.emit(AsyncMapEvent[T]{RemovedEventAction, item})
	delete(q.items, id)
}

// returns a list of all items currently in the map
func (q *AsyncMap[T]) Entries() []AsyncMapItem[T] {
	q.mutex.RLock()
	defer q.mutex.RUnlock()

	entries := make([]AsyncMapItem[T], len(q.items))
	i := 0
	for _, item := range q.items {
		entries[i] = item
		i += 1
	}

	return entries
}

// retrieves value by id
func (q *AsyncMap[T]) Get(id string) *AsyncMapItem[T] {
	q.mutex.RLock()
	defer q.mutex.RUnlock()

	item, ok := q.items[id]
	if !ok {
		return nil
	}

	return &item
}
