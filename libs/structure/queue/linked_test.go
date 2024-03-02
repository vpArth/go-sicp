package queue_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/vpArth/go-sicp/libs/structure/queue"
	"sync"
	"testing"
)

func TestLinked(t *testing.T) {
	t.Parallel()
	t.Run("EnDeQueue", func(t *testing.T) {
		t.Parallel()
		var (
			el, peek int
			err      error
		)
		q := queue.NewLinked[int]()
		peek, err = q.Peek()
		assert.ErrorIs(t, err, queue.ErrorQueueEmpty)
		q.Enqueue(1)
		q.Enqueue(2)
		q.Enqueue(3)

		el, _ = q.Dequeue()
		assert.Equal(t, 1, el)
		peek, _ = q.Peek()
		assert.Equal(t, 2, peek)
		el, _ = q.Dequeue()
		assert.Equal(t, 2, el)
		el, _ = q.Dequeue()
		assert.Equal(t, 3, el)
		_, err = q.Dequeue()
		assert.ErrorIs(t, err, queue.ErrorQueueEmpty)
	})

	t.Run("Concurrency", func(t *testing.T) {
		t.Parallel()
		q := queue.NewLinked[int]()

		const N = 1000

		var wgProduce sync.WaitGroup
		wgProduce.Add(N)
		for i := 0; i < N; i++ {
			go func(num int) {
				defer wgProduce.Done()
				q.Enqueue(num)
			}(i + 1)
		}
		wgProduce.Wait()

		sum := 0
		var (
			wgConsume sync.WaitGroup
			mux       sync.Mutex
		)
		wgConsume.Add(N)
		for i := 0; i < N; i++ {
			go func() {
				defer wgConsume.Done()
				el, _ := q.Dequeue()

				mux.Lock()
				sum += el
				mux.Unlock()
			}()
		}
		wgConsume.Wait()

		assert.Equal(t, N*(N+1)/2, sum)
	})
}
