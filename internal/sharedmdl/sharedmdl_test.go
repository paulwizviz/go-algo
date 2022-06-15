package sharedmdl

import (
	"fmt"
)

func Example_numericCmp_int8() {
	lhs := int8(1)
	rhs := int8(1)
	fmt.Printf("lhs(%v) equal rhs(%v). Result: %v\n", lhs, rhs, NumericCmp(lhs, rhs))

	lhs1 := int32(-1)
	rhs1 := int32(1)
	fmt.Printf("lhs(%v) < rhs(%v). Result: %v\n", lhs1, rhs1, NumericCmp(lhs1, rhs1))

	lhs2 := int64(1)
	rhs2 := int64(0)
	fmt.Printf("lhs(%v) > rhs(%v). Result: %v\n", lhs2, rhs2, NumericCmp(lhs2, rhs2))

	// Output:
	// lhs(1) equal rhs(1). Result: 0
	// lhs(-1) < rhs(1). Result: -1
	// lhs(1) > rhs(0). Result: 1

}

func Example_numericCmp_float() {
	lhs := float32(1.0)
	rhs := float32(1.0)
	fmt.Printf("lhs(%v) equal rhs(%v). Result: %v\n", lhs, rhs, NumericCmp(lhs, rhs))

	lhs1 := float64(-1.1)
	rhs1 := 1.1
	fmt.Printf("lhs(%v) < rhs(%v). Result: %v\n", lhs1, rhs1, NumericCmp(lhs1, rhs1))

	lhs2 := 1.0
	rhs2 := 0.1
	fmt.Printf("lhs(%v) > rhs(%v). Result: %v\n", lhs2, rhs2, NumericCmp(lhs2, rhs2))

	// Output:
	// lhs(1) equal rhs(1). Result: 0
	// lhs(-1.1) < rhs(1.1). Result: -1
	// lhs(1) > rhs(0.1). Result: 1

}
