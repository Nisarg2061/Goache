package main

import (
	"fmt"

	"github.com/Nisarg2061/Goache/Cache"
)

func main() {
	fmt.Println("Start Cache...")
	cache := Cache.NewCache()
	listOfItems := []string{"Cat", "Dog", "Turtle", "Mouse", "Bat", "Dog", "Nisarg", "Fox", "Tiger"}

	for _, val := range listOfItems {
		cache.Check(val)
		cache.Display()
	}
}
