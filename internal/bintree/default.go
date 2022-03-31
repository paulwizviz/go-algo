package bintree

import "github.com/paulwizviz/go-data/internal/model"

// default implementation of a binary node using pointers.
// NOTE: This is not serializable.
type defaultNode[N model.NumericType] struct {
	count uint32
	key   N
	left  Node[N]
	right Node[N]
}

func (p *defaultNode[N]) Key() N {
	return p.key
}

func (d *defaultNode[N]) Count() uint32 {
	return d.count
}

func (d *defaultNode[N]) AddCount(count uint32) {
	d.count = count
}

func (p *defaultNode[N]) Left() Node[N] {
	return p.left
}

func (p *defaultNode[N]) SetLeft(left Node[N]) {
	p.left = left
}

func (p *defaultNode[N]) Right() Node[N] {
	return p.right
}

func (p *defaultNode[N]) SetRight(right Node[N]) {
	p.right = right
}

// NewDefaultNode instantiate a reference to adefault implementation
// of a Node
func NewDefaultNode[N model.NumericType](key N) Node[N] {
	return &defaultNode[N]{
		key: key,
	}
}