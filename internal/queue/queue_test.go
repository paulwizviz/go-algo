package queue

import "fmt"

func Example_create() {
	instance := Create[string](nil)
	fmt.Println(instance)

	instance = Create([]string{"abc", "def"})
	fmt.Println(instance)

	// Output:
	// []
	// [abc def]
}

func Example_enqueue() {
	qString := Create[string](nil)
	qString = Enqueue(qString, "abc")
	fmt.Println(qString)

	// Output:
	// [abc]
}

func Example_dequeue() {
	qString := Create([]string{"abc", "def"})
	fmt.Println(qString)

	qString = Dequeue(qString)
	fmt.Println(qString)

	// Output:
}
