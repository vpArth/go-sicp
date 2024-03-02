package stack

import (
	"errors"
	"sync"
)

var _ Stack[any] = (*Slice[any])(nil)

type Slice[T any] struct {
	elements []T

	lock sync.Mutex
}

func NewSlice[T any]() *Slice[T] {
	stack := &Slice[T]{
		elements: make([]T, 0),
		lock:     sync.Mutex{},
	}

	return stack
}

func (s *Slice[T]) Push(el T) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.elements = append(s.elements, el)
}

func (s *Slice[T]) Pop() (el T, err error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if len(s.elements) == 0 {
		return el, errors.New("empty stack")
	}
	el, s.elements = s.elements[len(s.elements)-1], s.elements[:len(s.elements)-1]

	return el, nil
}
