package queue

func Create[T any](seed []T) []T {
	var instance []T
	instance = append(instance, seed...)
	return instance
}

func Enqueue[T any](queue []T, item T) []T {
	return append(queue, item)
}

func Dequeue[T any](queue []T) []T {
	return queue[1:]
}
