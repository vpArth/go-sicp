package stack

import (
	"errors"
)

var _ Stack[any] = (*Slice[any])(nil)

type Slice[T any] struct {
	elements []T
}

func NewSlice[T any]() *Slice[T] {
	stack := &Slice[T]{}
	stack.elements = make([]T, 0)

	return stack
}

func (s *Slice[T]) Push(el T) {
	s.elements = append(s.elements, el)
}

func (s *Slice[T]) Pop() (el T, err error) {
	if len(s.elements) == 0 {
		return el, errors.New("empty stack")
	}
	el, s.elements = s.elements[len(s.elements)-1], s.elements[:len(s.elements)-1]

	return el, nil
}
