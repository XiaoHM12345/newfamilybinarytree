package main

import (
	"familyBinaryTree/binarytree"
	"fmt"
)

func main() {
	tree := Binarytree.NewBinaryTree()
	(*tree).SetRoot(&Binarytree.BinaryTreeNode{"dad", "a", []string{"a1"}, []string{}, nil, nil})
	(*tree).BuildTreeFromConsole()
	if (*tree).AddMember("b","f") {
		fmt.Printf("true\n")
	} else {
		fmt.Printf("err\n")
	}

	if (*tree).ChangeName("c","g") {
		fmt.Printf("true\n")
	} else {
		fmt.Printf("err\n")
	}

	tree.BinaryTreePreOder()
}
