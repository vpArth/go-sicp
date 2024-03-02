package stack_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/vpArth/go-sicp/libs/structure/stack"
	"testing"
)

func TestSlice(t *testing.T) {
	t.Run("PushPop", func(t *testing.T) {
		s := stack.NewSlice[int]()
		s.Push(1)
		s.Push(2)
		s.Push(3)

		var el int

		el, _ = s.Pop()
		assert.Equal(t, 3, el)
		el, _ = s.Pop()
		assert.Equal(t, 2, el)
		el, _ = s.Pop()
		assert.Equal(t, 1, el)
		_, err := s.Pop()
		assert.EqualError(t, err, "empty stack")
	})
}
