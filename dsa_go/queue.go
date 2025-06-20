package dsa

import "fmt"

type queue struct {
	slice []int
}

func (q *queue) Enqueue(num int) {
	q.slice = append(q.slice, num)
}

func (q *queue) Dequeue() {
	pop := q.slice[0]
	q.slice = q.slice[1:]
	fmt.Println("Removed element is", pop)
	fmt.Println(*q)
}

func ImplementQueue() {
	Queue := queue{}
	Queue.Enqueue(1)
	Queue.Enqueue(2)
	Queue.Enqueue(3)
	Queue.Dequeue()
}
