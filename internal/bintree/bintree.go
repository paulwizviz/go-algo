package bintree

import (
	"fmt"

	"github.com/paulwizviz/go-data/internal/model"
)

// Error represents a error related to the operations of cache
type Error struct {
	message string
}

func (e *Error) Error() string {
	return fmt.Sprintf("Binary tree. Error: %v", e.message)
}

func newError(message string) error {
	return &Error{
		message: message,
	}
}

// Node of a binary tree
type Node[N model.NumericType] interface {

	// Key returns the node's key
	Key() N

	// Count the number of duplicate keys
	Count() uint32

	// AddCount to number of duplicate keys
	AddCount(count uint32)

	// Left return the left child
	Left() Node[N]

	// SetLeft child
	SetLeft(left Node[N])

	// Right returns the right child
	Right() Node[N]

	// SetRight right child
	SetRight(left Node[N])
}

type NewNode[N model.NumericType] func(key N) Node[N]

func InsertNode[N model.NumericType](newNode NewNode[N], root Node[N], key N) {

	if root.Key() == key {
		root.AddCount(root.Count() + 1)
		return
	}

	if root.Key() < key {
		child := root.Right()
		if child == nil {
			child = newNode(key)
			child.AddCount(1)
			root.SetRight(child)
			return
		}
		InsertNode(newNode, root.Right(), key)
		return
	}

	child := root.Left()
	if child == nil {
		child = NewDefaultNode(key)
		child.AddCount(1)
		root.SetLeft(child)
		return
	}
	InsertNode(newNode, child, key)
	return

}
