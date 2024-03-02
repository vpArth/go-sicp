package stack

type Stack[T any] interface {
	Push(el T)
	Pop() (T, error)
}
