package Binarytree

const MAXSIZE = 20

type QueueArray struct {
	Data			[MAXSIZE + 1]*BinaryTreeNode
	Front			int16
	Rear			int16
}

func (c *QueueArray)InitQueue () {
	c.Front, c.Rear = 0, 0
}

func (c *QueueArray)IsEmpty() bool {
	return c.Rear == c.Front
}

func (c *QueueArray)EnQueue(node *BinaryTreeNode) bool {
	if (c.Rear + 1) % MAXSIZE == c.Front {
		return false
	}
	c.Data[c.Rear] = node
	c.Rear = (c.Rear + 1) % MAXSIZE
	return true
}

func (c *QueueArray)SqQueue() (*BinaryTreeNode, bool) {
	if c.Rear == c.Front {
		return nil,false
	}
	x := c.Data[c.Front]
	c.Front = (c.Front + 1) % MAXSIZE

	return x,true
}

func (c *QueueArray)GetHead() (*BinaryTreeNode, bool) {
	if c.Rear == c.Front {
		return nil,false
	}
	x := c.Data[c.Front]

	return x,true
}

func (c *QueueArray)GetSize() (int16) {
	return (c.Rear + MAXSIZE - c.Front) % MAXSIZE
}