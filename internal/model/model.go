package model

import (
	"math/rand"
	"time"
)

// Error represents related to model operations
type Error struct {
	message string
}

func (e *Error) Error() string {
	return e.message
}

func newError(message string) error {
	return &Error{
		message: message,
	}
}

// Identifier a unique time based value
type Identifier int64

// NewIdentifier instantiate an identifier
func NewIdentifier() Identifier {
	rand.Seed(time.Now().Local().UnixNano())
	return Identifier(time.Now().Unix() + rand.Int63n(256))
}

// NumericType is a generic type constrains consisting of Go numerics.
// NOTE: This is only for demonstration purpose only. There are a number
// of equivalent constraints in standard packages. One example is the
// `constraints` package. You can also use an alias to numeric constraints
// call `comparable`.
type NumericType interface {
	uint8 | uint16 | uint32 | uint64 | int8 | int16 | int32 | int64 | float32 | float64
}

// NumericCmp compares to numerical types.
// 1) LHS > RHS returns 1.
// 2) LHS < RHS returns -1
// 3) LHS == RHS returns 0
// NOTE: This demonstrate the use of generics with NumericType
// constraints.
func NumericCmp[N NumericType](lhs N, rhs N) int8 {

	if lhs-rhs == 0 {
		return 0
	}

	if lhs < rhs {
		return -1
	}

	return 1

}
