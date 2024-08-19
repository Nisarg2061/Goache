# Goache

## Description
Goache is a simple implementation of a Least Recently Used (LRU) Cache in Go, utilizing a Linked List data structure. 
This project is designed to demonstrate the practical application of data structures and algorithms in building a cache system.

## Concepts Used
- Data Structures and Algorithms (DSA)
- Linked Lists
- Caching Mechanisms

## Objectives
- LRU Cache: Implement a cache that follows the Least Recently Used (LRU) eviction policy.
- Item Reordering: On cache hit, the item is moved to the front of the list (head).
- Order Maintenance: Maintain the order of items based on their usage.
- Efficient Operations: Deletion occurs at the tail (least recently used), while addition occurs at the head (most recently used).

## Installation
To install Goache, use `go get`:
```
go get github.com/Nisarg2061/Goache
```


## Example Usage
Let's start by using the Goache cache system in a simple Go application:
```go
package main

import "github.com/Nisarg2061/Goache/Cache"

func main() {
	c := Cache.NewCache(5)
	c.Check("String")
	c.Display()
}
```
- Here, we create a new cache instance with a maximum capacity of 5 items. The `NewCache` function initializes the cache and sets the limit on the number of items it can hold. Once the limit is reached, the least recently used item will be evicted to make space for new items. 
- The `Check` method moves the item to the front of the list, marking it as the most recently used. If the item already exists in the cache, it updates its position.
- The `Display` method outputs the current state of the cache, showing its contents and the order of items. This helps you verify that the cache is functioning as expected and provides a way to inspect what’s currently stored in the cache.
## License
BSD licensed. See the LICENSE file for details.
