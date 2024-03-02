package queue

type Queue[T comparable] interface {
	Enqueue(T)
	Dequeue() (T, error)
	Peek() (T, error)
}
