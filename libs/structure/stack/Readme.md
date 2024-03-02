Actually, stack in Go can be implemented by slice itself.  

```go
var stack T[]
var value T
// Push
stack = append(stack, value)
// Pop
stack, value = stack[:len(stack)-1], stack[len(stack)-1]
```

This package exists just for learning purpose
_________________

### Slice implemetation of Stack interface

	Push(el T)
    adds el into stack

	Pop() (T, error)
    tries pop el from stack, gives an «empty stack» error, if empty
    
    Concurrency safety is implemented by simple mutex lock around of mutations
