package model

import (
	"math/rand"
	"time"
)

// Identifier a unique time based value
type Identifier int64

func NewIdentifier() Identifier {
	rand.Seed(time.Now().Local().UnixNano())
	return Identifier(time.Now().Unix() + rand.Int63n(256))
}
