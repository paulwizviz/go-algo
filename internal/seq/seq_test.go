package seq

import "fmt"

func Example_newNumericQueue() {
	nq := NewNumericQueue[uint8](nil)
	fmt.Println(nq)

	// Output:
	// &{[]}
}

func Example_enqueueNumericQueue() {
	nq := NewNumericQueue[uint8](nil)
	nq.Enqueue(1)
	fmt.Println(nq)
	nq.Enqueue(2)
	fmt.Println(nq)

	// Output:
	// &{[1]}
	// &{[1 2]}
}

func Example_dequeueNumericQueue() {
	nq := NewNumericQueue[uint8](nil)
	result, err := nq.Dequeue()
	fmt.Println(result, err)

	nq.Enqueue(1)
	nq.Enqueue(2)
	fmt.Println(nq)
	result, err = nq.Dequeue()
	fmt.Println(result, err)

	// Output:
	// 0 empty queue
	// &{[1 2]}
	// 1 <nil>
}
