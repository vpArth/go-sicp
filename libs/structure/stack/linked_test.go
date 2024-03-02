package stack_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/vpArth/go-sicp/libs/structure/stack"
	"sync"
	"testing"
)

func TestLinked(t *testing.T) {
	t.Parallel()
	t.Run("PushPop", func(t *testing.T) {
		t.Parallel()
		s := stack.NewLinked[int]()
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

	t.Run("Concurrency", func(t *testing.T) {
		t.Parallel()

		s := stack.NewLinked[int]()
		const N = 100

		var wgPush sync.WaitGroup
		wgPush.Add(N)
		for i := 0; i < N; i++ {
			go func(num int) {
				defer wgPush.Done()
				s.Push(num)
			}(i + 1)
		}
		wgPush.Wait()

		sum := 0
		var (
			wgPop sync.WaitGroup
			mux   sync.Mutex
		)

		wgPop.Add(N)
		for i := 0; i < N; i++ {
			go func() {
				defer wgPop.Done()
				el, _ := s.Pop()

				mux.Lock()
				sum += el
				mux.Unlock()
			}()
		}
		wgPop.Wait()

		assert.Equal(t, N*(N+1)/2, sum)
	})
}
