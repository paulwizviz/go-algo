package sorter

import (
	"golang.org/x/exp/constraints"
)

// BubbleOT is a bubble sort implementation of ordered types.
func BubbleOT[T constraints.Ordered](collection []T) []T {
	swapped := make([]T, len(collection))
	copy(swapped, collection)
	for range collection {
		for index := range swapped {
			if index+1 < len(collection)-1 {
				if swapped[index] > swapped[index+1] {
					swapped[index], swapped[index+1] = swapped[index+1], swapped[index]
				}
			}
			if index+1 == len(collection)-1 {
				if swapped[index] > swapped[index+1] {
					swapped[index], swapped[index+1] = swapped[index+1], swapped[index]
				}
			}
		}
	}

	return swapped
}
