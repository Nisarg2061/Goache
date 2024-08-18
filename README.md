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

## Example Usage
```go
package main

import "github.com/Nisarg2061/Goache/Cache"

func main() {
	c := Cache.NewCache()
	c.Check("String")
	c.Display()
}
```

## Installation
TO install Goache, use `go get`:
```
go get github.com/Nisarg2061/Goache
```

## Steps to Run the project locally
1. Clone the repository:
```
git clone https://github.com/Nisarg2061/Goache.git
```
2. Navigate to the project directory:
```
cd Goache
```
3. Run the project using Make:
```
make run
```
