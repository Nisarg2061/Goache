package main

import (
	"fmt"
)

func main() {
	fmt.Println("Start Cache...")
	cache := NewCache()
	listOfItems := []string{"Cat", "Dog", "Turtle", "Mouse", "Bat", "Dog", "Nisarg", "Fox", "Tiger"}

	for _, val := range listOfItems {
		cache.Check(val)
		cache.Display()
	}
}
