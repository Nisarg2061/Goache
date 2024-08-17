package main

const SIZE = 5

type (
	Node struct {
		Value string
		Left  *Node
		Right *Node
	}
	Queue struct {
		Head   *Node
		Tail   *Node
		Length int
	}
	Cache struct {
		Queue Queue
		Hash  Hash
	}
	Hash map[string]*Node
)
