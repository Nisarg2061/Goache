package Cache

var SIZE int

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
