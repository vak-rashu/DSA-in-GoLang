package dsa

import "fmt"

type stack struct {
	slice []int
}

func (s *stack) Push(value int) {
	s.slice = append(s.slice, value)
	fmt.Println("Value Pushed.")
}

func (s *stack) Pop() {
	len := len(s.slice)
	lastIndex := len - 1
	pop := s.slice[lastIndex]
	s.slice = s.slice[:lastIndex]
	fmt.Printf("Removed element is %v", pop)
}

func ImplementStack() {
	Stack := stack{}
	Stack.Push(1)
	Stack.Push(2)
	Stack.Push(3)
	Stack.Pop()
	Stack.Pop()

}
