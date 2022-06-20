package queue

import "fmt"

type Test struct {
	content int
}

func Example_variableNew() {
	vq := NewVariable[uint8]()
	fmt.Println(vq)

	vq0 := NewVariable(1.1, 1.2) // default to float32
	fmt.Println(vq0)

	vq01 := NewVariable[float32](1.1, 1.2)
	fmt.Println(vq01)

	vq1 := NewVariable("abcd", "efgh")
	fmt.Println(vq1)

	vq2 := NewVariable('a', 'b')
	fmt.Println(vq2)

	vq3 := NewVariable[byte]('a', 'b')
	fmt.Println(vq3)

	vq4 := NewVariable(Test{content: 1})
	fmt.Println(vq4)

	// Output:
	// &{[]}
	// &{[1.1 1.2]}
	// &{[1.1 1.2]}
	// &{[abcd efgh]}
	// &{[97 98]}
	// &{[97 98]}
	// &{[{1}]}
}

func Example_variableEnqueue() {
	vq := NewVariable[Test]()
	fmt.Println(vq)

	vq.Enqueue(Test{content: 12})
	fmt.Println(vq)

	// Output:
	// &{[]}
	// &{[{12}]}
}

func Example_variableDequeue() {
	vq := NewVariable[Test]()
	fmt.Println(vq)

	result := vq.Dequeue()
	fmt.Println(result)

	vq.Enqueue(Test{content: 400})
	result = vq.Dequeue()
	fmt.Println(result)

	// Output:
	// &{[]}
	// <nil>
	// &{400}
}

func Example_fixedNew() {
	fq := NewFixed[uint16](2)
	fmt.Println(fq)

	// Output:
	// &{[] 2}
}

func Example_fixedCapacity() {
	fq := NewFixed[uint16](2)
	capcity := fq.Capcity()
	fmt.Println(capcity)

	// Output:
	// 2
}

func Example_fixedEnqueue() {
	fq := NewFixed[uint16](2)
	fmt.Println(fq)

	result := fq.Enqueue(12)
	fmt.Println(result)

	result = fq.Enqueue(13)
	fmt.Println(result)

	result = fq.Enqueue(40)
	fmt.Println(result)

	// Output:
	// &{[] 2}
	// 0xc000016318
	// 0xc00001631a
	// <nil>
}

func Example_fixedIsFull() {
	fq := NewFixed[uint16](2)
	fmt.Println(fq)

	fq.Enqueue(12)
	result := fq.IsFull()
	fmt.Println(result)

	fq.Enqueue(13)
	result = fq.IsFull()
	fmt.Println(result)

	// Output:
	// &{[] 2}
	// false
	// true
}

func Example_fixedDequeue() {
	fq := NewFixed[uint16](2)
	fmt.Println(fq)

	result := fq.Dequeue()
	fmt.Println(result)

	// Output:
	// &{[] 2}
	// <nil>
}
