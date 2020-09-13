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

func PreOder(node *BinaryTreeNode, visit Visit) {
	if node != nil{
		//visit node
		//test
		//fmt.Printf("%s\n", node.Name)
		visit.Visit(node)
		//test
		PreOder(node.LeftChild, visit)
		PreOder(node.RightChild, visit)
	}
}

func InOder(node *BinaryTreeNode, visit Visit) {
	if node != nil{
		InOder(node.LeftChild, visit)
		//visit node
		//test
		//fmt.Printf("%s\n", node.Name)
		visit.Visit(node)
		//test
		InOder(node.RightChild, visit)
	}
}

func PostOder(node *BinaryTreeNode, visit Visit) {
	if node != nil{
		PostOder(node.LeftChild, visit)
		PostOder(node.RightChild, visit)
		//visit node
		//test
		//fmt.Printf("%s\n", node.Name)
		visit.Visit(node)
		//test
	}
}

func LevelOrder(node *BinaryTreeNode, visit Visit) {
	var queue QueueArray
	queue.InitQueue()

	temp := node

	for temp != nil {
		//visit node
		visit.Visit(temp)

		if (*temp).LeftChild != nil{
			queue.EnQueue((*temp).LeftChild)
		}
		if (*temp).RightChild != nil{
			queue.EnQueue((*temp).RightChild)
		}

		temp,_ = queue.SqQueue()
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

func SearchNode(name string, node *BinaryTreeNode) *BinaryTreeNode {
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

		if (*temp).LeftChild != nil{
			queue.EnQueue((*temp).LeftChild)
		}
		if (*temp).RightChild != nil{
			queue.EnQueue((*temp).RightChild)
		}

		temp,ok = queue.SqQueue()
		if ok == false {
			return nil
		}
	}

	return res
}
