package queue

type Variable[T any] interface {
	Dequeue() *T
	Enqueue(item T)
}

type defaultVariable[T any] struct {
	content []T
}

func (d *defaultVariable[T]) Dequeue() *T {
	if len(d.content) == 0 {
		return nil
	}
	item := d.content[0]
	d.content = d.content[1:]
	return &item
}

func (d *defaultVariable[T]) Enqueue(item T) {
	d.content = append(d.content, item)
}

func NewVariable[T any](seed ...T) Variable[T] {
	return &defaultVariable[T]{
		content: seed,
	}
}

type Fixed[T any] interface {
	Capcity() uint16
	Dequeue() *T
	Enqueue(item T) *T
	IsFull() bool
}

type defaultFixed[T any] struct {
	content []T
	capcity uint16
}

func (d defaultFixed[T]) Capcity() uint16 {
	return d.capcity
}

func (d *defaultFixed[T]) Dequeue() *T {
	if len(d.content) == 0 {
		return nil
	}
	item := d.content[0]
	d.content = d.content[1:]
	return &item
}

func (d *defaultFixed[T]) Enqueue(item T) *T {

	if len(d.content) == int(d.capcity) {
		return nil
	}
	d.content = append(d.content, item)
	return &d.content[len(d.content)-1]
}

func (d *defaultFixed[T]) IsFull() bool {
	if len(d.content) == int(d.capcity) {
		return true
	}
	return false
}

func NewFixed[T any](capcity uint16) Fixed[T] {
	return &defaultFixed[T]{
		content: make([]T, 0),
		capcity: capcity,
	}
}
