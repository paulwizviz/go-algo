package bintree

import "fmt"

func Example_insertNode() {

	root := NewDefaultNode[uint8](6)
	root.AddCount(1)
	InsertNode(NewDefaultNode[uint8], root, 3)
	InsertNode(NewDefaultNode[uint8], root, 7)
	InsertNode(NewDefaultNode[uint8], root, 3)
	InsertNode(NewDefaultNode[uint8], root, 2)
	InsertNode(NewDefaultNode[uint8], root, 8)
	InsertNode(NewDefaultNode[uint8], root, 4)

	fmt.Printf("Level 0. Node: %v(%v)\n", root.Key(), root.Count())
	fmt.Printf("Level 1. Nodes: %v(%v) %v(%v)\n", root.Left().Key(), root.Left().Count(), root.Right().Key(), root.Right().Count())

	fmt.Printf("Level 2. Nodes: %v(%v) %v(%v) %v(nil) %v(%v)", root.Left().Left().Key(), root.Left().Left().Count(), root.Left().Right().Key(), root.Left().Right().Count(), root.Right().Left(), root.Right().Right().Key(), root.Right().Right().Count())

	// Output:
	//
	// Level 0. Node: 6(1)
	// Level 1. Nodes: 3(2) 7(1)
	// Level 2. Nodes: 2(1) 4(1) <nil>(nil) 8(1)
}
