package main

import (
	dsa "dsa_go/dsa_go"
	"fmt"
)

func main() {
	fmt.Println("\nLinkedList")
	dsa.ImplementOps()

	fmt.Println("\nQueue")
	dsa.ImplementQueue()

	fmt.Println("\nStack")
	dsa.ImplementStack()

	fmt.Println("\nTree")
	dsa.ImplementTree()
}
