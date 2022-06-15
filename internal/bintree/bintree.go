package bintree

import (
	"github.com/paulwizviz/go-algo/internal/basemdl"
)

// Node[N model.NumericType] of a binary tree
type Node[N basemdl.NumericType] interface {

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

// NewNode[N model.NumericType] is a callback to an implementation to instantiate a node
type NewNode[N basemdl.NumericType] func(key N) Node[N]

// InsertNode[N model.NumericType] an operation to create a binary tree
func InsertNode[N basemdl.NumericType](newNode NewNode[N], root Node[N], key N) Node[N] {

	if root == nil {
		root = newNode(key)
	}

	if root.Key() == key {
		root.AddCount(root.Count() + 1)
	}

	if root.Key() < key {
		root.SetRight(InsertNode(newNode, root.Right(), key))
	}

	if root.Key() > key {
		root.SetLeft(InsertNode(newNode, root.Left(), key))
	}

	return root
}
