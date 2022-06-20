package sorter

import "fmt"

func Example_bubble() {
	iSlice := []int{5, 1, 4, 2, 3}
	fmt.Println("Before sort: ", iSlice)
	result1 := BubbleOT(iSlice)
	fmt.Println("After sort: ", result1)

	sSlice := []string{"c", "a", "z", "f"}
	fmt.Println("Before sort: ", sSlice)
	result2 := BubbleOT(sSlice)
	fmt.Println("After sort: ", result2)

	// Output:
	// Before sort:  [5 1 4 2 3]
	// After sort:  [1 2 3 4 5]
	// Before sort:  [c a z f]
	// After sort:  [a c f z]
}
