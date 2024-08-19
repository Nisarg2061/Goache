package Cache

func NewCache(s int) Cache {
	SIZE = s
	return Cache{Queue: NewQueue(), Hash: Hash{}}
}

func (c *Cache) Check(value string) {
	node := &Node{}
	if val, ok := c.Hash[value]; ok {
		node = c.Remove(val)
	} else {
		node = &Node{Value: value}
	}
	c.Add(node)
	c.Hash[value] = node
}
