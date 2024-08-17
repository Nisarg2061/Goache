package main

import "fmt"

func (c *Cache) Remove(n *Node) *Node {
	fmt.Println("Removing:", n.Value)

	left := n.Left
	right := n.Right

	left.Right = right
	right.Left = left
	c.Queue.Length -= 1
	delete(c.Hash, n.Value)

	return n
}

func (c *Cache) Add(n *Node) {
	fmt.Println("Adding:", n.Value)

	tmp := c.Queue.Head.Right

	c.Queue.Head.Right = n
	n.Left = c.Queue.Head
	n.Right = tmp
	tmp.Left = n

	c.Queue.Length += 1
	if c.Queue.Length > SIZE {
		c.Remove(c.Queue.Tail.Left)
	}
}

func (c *Cache) Display() {
	c.Queue.Display()
}
