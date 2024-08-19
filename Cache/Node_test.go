package Cache

import (
	"strconv"
	"testing"
)

func Test_Add(t *testing.T) {
	c := NewCache(5)
	node1 := &Node{Value: "Check"}
	c.Add(node1)

	// Testing for the position of value added
	if c.Queue.Head.Right != node1 {
		t.Errorf("Expected node1 to be at the right of head got %v", c.Queue.Head.Right)
	}

	// Exciding limit
	for i := 0; i < SIZE; i++ {
		chechNode := &Node{Value: string("c" + strconv.Itoa(i))}
		c.Add(chechNode)
	}

	if c.Queue.Tail.Left.Value == "Check" {
		t.Errorf("Expected node1 to be removed but it didn't.\n")
	}
}
