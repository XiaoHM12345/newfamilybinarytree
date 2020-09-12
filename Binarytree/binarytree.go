package Binarytree

type LinkedBinaryTree struct {
	root		*BinaryTreeNode
	treeSize	int16
}

func (c *LinkedBinaryTree) Empty() bool {
	return c.treeSize == 0
}

func (c *LinkedBinaryTree) Size() int16 {
	return c.treeSize
}

func (c *LinkedBinaryTree) BinaryTreePreOder() {
	PreOder(c.root)
}

func (c *LinkedBinaryTree) BinaryTreeInOder() {
	InOder(c.root)
}

func (c *LinkedBinaryTree) BinaryTreePostOder() {
	PostOder(c.root)
}

func NewBinaryTree() *LinkedBinaryTree {
	c := LinkedBinaryTree{nil, 0}
	return &c
}

func (c *LinkedBinaryTree)SetRoot(newNode *BinaryTreeNode) {
	c.root = newNode
}

func (c *LinkedBinaryTree)PtrRoot() *BinaryTreeNode {
	return c.root
}

func (c *LinkedBinaryTree)BuildTreeFromConsole() {
	BuildFromConsole(c.root, "noName")
}

func (c *LinkedBinaryTree)AddMember(parentName string, name string) bool {
	x := SearchCode(parentName, c.root)
	if x == nil {
		return false
	}
	if x.LeftChild == nil {
		(*x).LeftChild = &BinaryTreeNode{Name: name, Parent: parentName, LeftChild: nil, RightChild: nil}
	} else {
		x = x.LeftChild
		for x.RightChild != nil {
			x = x.RightChild
		}
		x.RightChild = &BinaryTreeNode{Name: name, Parent: parentName, LeftChild: nil, RightChild: nil}
	}
	return true
}