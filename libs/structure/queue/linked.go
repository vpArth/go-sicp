package queue

import "sync"

var _ Queue[any] = (*Linked[any])(nil)

type linkedItem[T comparable] struct {
	value T
	next  *linkedItem[T]
}
type Linked[T comparable] struct {
	front *linkedItem[T]
	back  *linkedItem[T]

	lock sync.RWMutex
}

func NewLinked[T comparable]() *Linked[T] {
	return &Linked[T]{
		lock: sync.RWMutex{},
	}
}

func (q *Linked[T]) isEmpty() bool {
	return q.front == nil
}

func (q *Linked[T]) Enqueue(el T) {
	q.lock.Lock()
	defer q.lock.Unlock()

	node := &linkedItem[T]{el, nil}

	if q.isEmpty() {
		q.front = node
		q.back = node
		return
	}

	q.back.next = node
	q.back = node
}

func (q *Linked[T]) Dequeue() (T, error) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.isEmpty() {
		var zeroT T
		return zeroT, ErrorQueueEmpty
	}

	node := q.front.value

	q.front = q.front.next
	if q.isEmpty() {
		q.back = nil
	}

	return node, nil
}

func (q *Linked[T]) Peek() (T, error) {
	q.lock.RLock()
	defer q.lock.RUnlock()

	if q.isEmpty() {
		var zeroT T
		return zeroT, ErrorQueueEmpty
	}

	return q.front.value, nil
}
