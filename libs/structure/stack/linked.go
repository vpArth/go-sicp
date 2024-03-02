package stack

import (
	"sync"
)

var _ Stack[any] = (*Linked[any])(nil)

type linkedItem[T any] struct {
	value T
	next  *linkedItem[T]
}

type Linked[T any] struct {
	head *linkedItem[T]

	lock sync.Mutex
}

func NewLinked[T any]() *Linked[T] {
	return &Linked[T]{
		lock: sync.Mutex{},
	}
}

func (s *Linked[T]) Push(el T) {
	s.lock.Lock()
	defer s.lock.Unlock()

	node := &linkedItem[T]{el, nil}

	node.next = s.head
	s.head = node
}

func (s *Linked[T]) Pop() (T, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.head == nil {
		var zeroT T
		return zeroT, ErrorStackEmpty
	}
	result := (*s.head).value

	s.head = s.head.next

	return result, nil
}
