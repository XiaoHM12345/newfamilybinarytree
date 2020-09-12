package Binarytree

import (
	"fmt"
)

type BinaryTreeNode struct {
	Parent 		string
	Name		string
	Wife		[]string
	Daughter	[]string
	LeftChild 	*BinaryTreeNode
	RightChild 	*BinaryTreeNode
}

func PreOder(node *BinaryTreeNode) {
	if node != nil{
		//visit node
		//test
		fmt.Printf("%s\n", node.Name)
		//test
		PreOder(node.LeftChild)
		PreOder(node.RightChild)
	}
}

func InOder(node *BinaryTreeNode) {
	if node != nil{
		InOder(node.LeftChild)
		//visit node
		//test
		fmt.Printf("%s\n", node.Name)
		//test
		InOder(node.RightChild)
	}
}

func PostOder(node *BinaryTreeNode) {
	if node != nil{
		PostOder(node.LeftChild)
		PostOder(node.RightChild)
		//visit node
		//test
		fmt.Printf("%s\n", node.Name)
		//test
	}
}

func BuildFromConsole(node *BinaryTreeNode, parent string) {
	var name string
	var answer string

	fmt.Printf("input name\n")
	fmt.Scanf("%s", &name)
	node.Name = name
	node.Parent = parent

	fmt.Printf("if this node names %s has child?\n", node.Name)
	fmt.Scanf("%s", &answer)
	if answer == "Y" || answer == "y" {
		fmt.Printf("enter the child name of %s\n", node.Name)
		node.LeftChild = &BinaryTreeNode{}
		BuildFromConsole(node.LeftChild, node.Name)
	}

	fmt.Printf("if this node names %s has brother?\n", node.Name)
	fmt.Scanf("%s", &answer)
	if answer == "Y" || answer == "y" {
		fmt.Printf("enter the brother name of %s\n", node.Name)
		node.RightChild = &BinaryTreeNode{}
		BuildFromConsole(node.RightChild, parent)
	}
}

func SearchCode(name string, node *BinaryTreeNode) *BinaryTreeNode {
	var queue QueueArray
	queue.InitQueue()
	var res *BinaryTreeNode
	res = nil
	var ok bool

	temp := node

	for temp != nil {
		//visit node
		if name == (*temp).Name {
			res = temp
			break
		}

		if (*node).LeftChild != nil{
			queue.EnQueue((*temp).LeftChild)
		}
		if (*node).RightChild != nil{
			queue.EnQueue((*temp).RightChild)
		}

		temp,ok = queue.SqQueue()
		if ok == false {
			return nil
		}
	}

	return res
}