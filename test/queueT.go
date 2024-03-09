package test

import (
	"DSA/DataStructures/queue"
	"fmt"
)

func TestQueue() {
	queue := &queue.QueueSlice{}

	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)
	queue.Enqueue(40)

	fmt.Print("Elements: ")
	queue.DisplayElements()

	value, ok := queue.Dequeue()
	if !ok {
		fmt.Println("Queue is  empty")
	}
	fmt.Println("Dequeued element is", value)

	fmt.Print("Elements: ")
	queue.DisplayElements()
}
