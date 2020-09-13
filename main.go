package main

import (
	"familyBinaryTree/binarytree"
	"fmt"
)

type myVisit struct {
}

func (c *myVisit) Visit (node *Binarytree.BinaryTreeNode)  {
	fmt.Printf("%s" ,(*node).Name)
}

func main() {
	tree := Binarytree.NewBinaryTree()
	(*tree).SetRoot(&Binarytree.BinaryTreeNode{"dad", "a", []string{"a1"}, []string{}, nil, nil})
	(*tree).BuildTreeFromConsole()

	if (*tree).AddMember("b","f") {
		fmt.Printf("true\n")
	} else {
		fmt.Printf("err\n")
	}
/*
	if (*tree).ChangeName("c","g") {
		fmt.Printf("true\n")
	} else {
		fmt.Printf("err\n")
	}
*/
	tree.BinaryTreePreOder(new(myVisit))
	fmt.Printf("\n")
	tree.BinaryTreeInOder(new(myVisit))
	fmt.Printf("\n")
	tree.BinaryTreePostOder(new(myVisit))
	fmt.Printf("\n")
	tree.BinaryLevelPostOder(new(myVisit))
}
