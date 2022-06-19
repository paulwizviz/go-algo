package seq

import "fmt"

var (
	EmptyQueueErr = fmt.Errorf("empty queue")
)

type numericType interface {
	uint8 | uint16 | uint32 | uint64 | int8 | int16 | int32 | int64 | float32 | float64
}

type NumericQueue[T numericType] interface {
	Enqueue(item T)
	Dequeue() (T, error)
}

type variableNQ[T numericType] struct {
	content []T
}

func (n *variableNQ[T]) Enqueue(item T) {
	n.content = append(n.content, item)
}

func (n *variableNQ[T]) Dequeue() (T, error) {
	if len(n.content) == 0 {
		return 0, EmptyQueueErr
	}
	item := n.content[0]
	n.content = n.content[1:]
	return item, nil
}

func NewNumericQueue[T numericType](seed []T) NumericQueue[T] {
	return &variableNQ[T]{
		content: seed,
	}
}
