// Package trie provides the Trie data types and related operations
// The Trie is a structure for storing strings.
// WARNING: The implementations here are not intended for
// mission critical use cases. These are only for demonstrations.
package trie

import (
	"fmt"
	"unicode"
)

const (
	EOLCharacter rune = '\n'
)

// Error related to trie data structure and related operation
type Error struct {
	message string
}

func (e *Error) Error() string {
	return e.message
}

func newError(message string) *Error {
	return &Error{
		message: message,
	}
}

// Node of a trie
type Node interface {
	// InsertRune value of edge from this node
	InsertRune(data rune) Node

	// IsLeaf determines if this node
	IsLeaf() bool

	// MatchRune determines if a rune matches one of the edge.
	MatchedRune(data rune) (isEnd bool, isMatched bool)

	// NextNode return the nodes that is reference by the edge named 'ch'
	NextNode(ch rune) Node

	// Runes returns the edges connected to this node
	Runes() []rune
}

// This is not serializable
type defaultNode struct {
	edges map[rune]Node
}

func (d *defaultNode) InsertRune(data rune) Node {
	lrune := unicode.ToLower(data)
	node, ok := d.edges[lrune]
	if !ok {
		d.edges[lrune] = NewDefaultNode()
		return d.edges[lrune]
	}
	return node
}

func (d *defaultNode) IsLeaf() bool {
	return len(d.edges) == 0
}

func (d *defaultNode) MatchedRune(data rune) (isEnd bool, isMatched bool) {
	isMatched, isEnd = false, false
	nextNode, ok := d.edges[data]
	if !ok {
		return isEnd, isMatched
	}
	isEnd, isMatched = nextNode.IsLeaf(), true
	return isEnd, isMatched
}

func (d *defaultNode) NextNode(ch rune) (node Node) {
	node, ok := d.edges[ch]
	if !ok {
		return nil
	}
	return node
}

func (d *defaultNode) Runes() []rune {
	var runes []rune

	for r, _ := range d.edges {
		runes = append(runes, r)
	}

	return runes
}

// NewDefaultNode instantiate a default Trie node.
func NewDefaultNode() Node {
	return &defaultNode{
		edges: make(map[rune]Node),
	}
}

// Populate operation to populate a Trie
func Populate(trie Node, str string) {
	var node Node
	for index, ch := range str {
		if index == 0 {
			node = trie.InsertRune(ch)
		} else {
			node = node.InsertRune(ch)
		}
	}
}

// VerifyWord against a pre-populated trie
func VerifyWord(trie Node, word string) error {
	var isEnd, isMatched bool
	var err error
	var node Node = trie
	for _, ch := range word {
		isEnd, isMatched = node.MatchedRune(ch)
		if !isMatched {
			return newError(fmt.Sprintf("Trie search error. Unable to match criteria: %v", word))
		}
		node = node.NextNode(ch)
	}
	if !isEnd {
		err = newError(fmt.Sprintf("Trie search error. Unable to match criteria: %v", word))
	} else {
		err = nil
	}
	return err
}

// DeleteWord from a trie
func DeleteWord(trie Node, word string) error {
	var todelete map[rune]Node = make(map[rune]Node)
	var err error
	var isEnd, isMatched bool

	var node Node = trie
	for _, ch := range word {
		todelete[ch] = node
		isEnd, isMatched = node.MatchedRune(ch)
		if !isMatched {
			return newError(fmt.Sprintf("Trie search error. Unable to delete word: %v", word))
		}
		node = node.NextNode(ch)
	}
	if !isEnd {
		err = newError(fmt.Sprintf("Trie search error. Unable to delete word: %v", word))
	} else {
		for k, _ := range todelete {
			delete(todelete, k)
		}
		err = nil
	}
	return err
}
