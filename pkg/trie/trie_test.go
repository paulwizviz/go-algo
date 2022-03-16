package trie

import (
	"fmt"
	"strconv"
)

func Example_isLeaf() {
	node := NewDefaultNode()
	fmt.Printf("Is a terminated edge? %v", node.InsertRune('a').IsLeaf())

	// Output:
	// Is a terminated edge? true
}

func Example_matchedRune() {
	trie := NewDefaultNode()
	node := trie.InsertRune('a')
	node.InsertRune('b')

	isEnd, isMatched := trie.MatchedRune('a')
	fmt.Printf("Matching first rune. isEnd: %v isMatched: %v\n", isEnd, isMatched)

	isEnd, isMatched = trie.MatchedRune('c')
	fmt.Printf("Not matching first rune. isEnd: %v isMatched: %v\n", isEnd, isMatched)

	isEnd, isMatched = node.MatchedRune('b')
	fmt.Printf("Matching terminated rune. isEnd: %v isMatched: %v\n", isEnd, isMatched)

	isEnd, isMatched = node.MatchedRune('c')
	fmt.Printf("Not matching terminated rune. isEnd: %v isMatched: %v", isEnd, isMatched)

	// Output:
	// Matching first rune. isEnd: false isMatched: true
	// Not matching first rune. isEnd: false isMatched: false
	// Matching terminated rune. isEnd: true isMatched: true
	// Not matching terminated rune. isEnd: false isMatched: false
}

func Example_populateTrie() {

	trie := NewDefaultNode()
	Populate(trie, "abc")

	var node Node = trie
	var res []string
	for {
		runes := node.NextRunes()
		str := strconv.QuoteRune(runes[0])
		res = append(res, str)
		node = node.NextNode(runes[0])
		if node.IsLeaf() {
			break
		}
	}

	fmt.Println(res)

	// Output:
	// ['a' 'b' 'c']
}

func Example_searchTrie() {

	trie := NewDefaultNode()
	Populate(trie, "ant")
	Populate(trie, "another")

	result := Search(trie, "ant")
	fmt.Println(result)

	result = Search(trie, "another")
	fmt.Println(result)

	result = Search(trie, "an")
	fmt.Println(result)

	// Output:
	// <nil>
	// <nil>
	// Trie search error. Unable to match criteria: an
}
