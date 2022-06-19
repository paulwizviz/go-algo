package queue

type numericType interface {
	uint8 | uint16 | uint32 | uint64 | int8 | int16 | int32 | int64 | float32 | float64
}

type charSliceType interface {
	rune | byte
}

type primitiveType interface {
	numericType | charSliceType | string
}

type Queue[T primitiveType] interface {
}

type variableQueue[T primitiveType] struct {
	content []T
}

func (d *variableQueue[T]) dequeue() *T {
	if len(d.content) == 0 {
		return nil
	}
	item := d.content[0]
	d.content = d.content[1:]
	return &item
}

func (d *variableQueue[T]) enqueue(item T) *T {
	d.content = append(d.content, item)
	return &d.content[len(d.content)-1]
}

func NewVariableQueue[T primitiveType](seed ...T) Queue[T] {
	return &variableQueue[T]{
		content: seed,
	}
}

type fixedQueue[T primitiveType] struct {
	content []T
	size    uint16
}

func (f *fixedQueue[T]) dequeue() *T {
	if len(f.content) == 0 {
		return nil
	}
	item := f.content[0]
	f.content = f.content[1:]
	return &item
}

func (f *fixedQueue[T]) enqueue(item T) *T {

	if len(f.content) == int(f.size) {
		return nil
	}
	f.content = append(f.content, item)
	return &f.content[len(f.content)-1]
}

func NewFixQueue[T primitiveType](size uint16) Queue[T] {
	return &fixedQueue[T]{
		content: make([]T, 0),
		size:    size,
	}
}

func Dequeue[T primitiveType](queue Queue[T]) (Queue[T], *T) {
	switch v := queue.(type) {
	case *variableQueue[T]:
		item := v.dequeue()
		return v, item
	case *fixedQueue[T]:
		item := v.dequeue()
		return v, item
	default:
		return nil, nil
	}
}

func Enqueue[T primitiveType](queue Queue[T], item T) (Queue[T], *T) {
	switch v := queue.(type) {
	case *variableQueue[T]:
		item := v.enqueue(item)
		return v, item
	case *fixedQueue[T]:
		item := v.enqueue(item)
		return v, item
	default:
		return nil, nil
	}
}
