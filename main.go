package main

import "fmt"

type Node struct{}

type Queue struct{}

type Cache struct {
	Queue Queue
	Hash  Hash
}

type Hash map[string]*Node

func NewCache() {
}

func main() {
	fmt.Println("Start Cache...")
	listOfItems := []string{"Cat", "Dog", "Turtle", "Mouse", "Bat"}
	fmt.Println(listOfItems)
}
