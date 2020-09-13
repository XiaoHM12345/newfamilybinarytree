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

func (c *LinkedBinaryTree) BinaryTreePreOder(visit Visit) {
	PreOder(c.root, visit)
}

func (c *LinkedBinaryTree) BinaryTreeInOder(visit Visit) {
	InOder(c.root, visit)
}

func (c *LinkedBinaryTree) BinaryTreePostOder(visit Visit) {
	PostOder(c.root, visit)
}

func (c *LinkedBinaryTree) BinaryLevelPostOder(visit Visit) {
	LevelOrder(c.root, visit)
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
	x := SearchNode(parentName, c.root)
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

func (c *LinkedBinaryTree)ChangeName(oldName string, newName string) bool {
	var queue QueueArray
	queue.InitQueue()
	temp := c.root

	if c.root == nil {
		return false
	}

	for temp != nil {
		//visit node
		if oldName == (*temp).Name {
			(*temp).Name = newName
		} else if oldName == (*temp).Parent {
			(*temp).Parent = newName
		}

		if (*temp).LeftChild != nil{
			queue.EnQueue((*temp).LeftChild)
		}
		if (*temp).RightChild != nil{
			queue.EnQueue((*temp).RightChild)
		}

		temp,_ = queue.SqQueue()
	}
	return true
}


