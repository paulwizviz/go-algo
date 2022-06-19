package queue

import "fmt"

func Example_newVariableQueue() {
	vq := NewVariableQueue[uint8]()
	fmt.Println(vq)

	vq = NewVariableQueue[uint8](1, 2, 3)
	fmt.Println(vq)

	vq1 := NewVariableQueue('a')
	fmt.Println(vq1)

	vq2 := NewVariableQueue[byte]('a')
	fmt.Println(vq2)

	// Output:
	// &{[]}
	// &{[1 2 3]}
	// &{[97]}
	// &{[97]}
}

func Example_newFixedQueue() {
	fq := NewFixQueue[uint8](1)
	fmt.Println(fq)

	// Output:
}

func Example_enqueueVariable() {
	vq := NewVariableQueue[uint8]()
	vq, item := Enqueue(vq, 1)
	vq, item = Enqueue(vq, 2)
	fmt.Println(vq, *item)

	// Output:
	// &{[1 2]} 2
}

func Example_dequeueVariable() {
	vq := NewVariableQueue[uint8](1, 2, 3)
	fmt.Println(vq)

	vq, item := Dequeue(vq)
	fmt.Println(vq, *item)

	vq, item = Dequeue(vq)
	fmt.Println(vq, *item)

	vq, item = Dequeue(vq)
	fmt.Println(vq, *item)

	vq, item = Dequeue(vq)
	fmt.Println(vq, item)

	// Output:
	// &{[1 2 3]}
	// &{[2 3]} 1
	// &{[3]} 2
	// &{[]} 3
	// &{[]} <nil>
}

func Example_enqueueFixed() {
	fq := NewFixQueue[uint8](3)
	fmt.Println(fq)

	fq, item := Enqueue(fq, 1)
	fmt.Println(item)

	fq, item = Enqueue(fq, 1)
	fmt.Println(item)

	fq, item = Enqueue(fq, 1)
	fmt.Println(item)

	_, item = Enqueue(fq, 1)
	fmt.Println(item)

	// Output:
	// 	&{[] 3}
	// 0xc0000a4238
	// 0xc0000a4239
	// 0xc0000a423a
	// <nil>
}

func Example_dequeFixed() {
	fq := NewFixQueue[uint8](3)
	fmt.Println(fq)

	fq, item := Dequeue(fq)
	fmt.Println(item)

	fq, item = Enqueue(fq, 1)
	if item != nil {
		fmt.Println(*item)
	}

	fq, item = Dequeue(fq)
	if item != nil {
		fmt.Println(*item)
	}

	// Output:
	// &{[] 3}
	// <nil>
	// 1
	// 1
}
