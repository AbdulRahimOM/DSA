package queue

import "fmt"

type queueList struct {
	head *node
	tail *node
}

type node struct {
	data interface{}
	next *node
}

func NewQueueList() *queueList {
	return &queueList{
		head: nil,
		tail: nil,
	}
}

func (q *queueList) Enqueue(value interface{}) {
	newNode := &node{data: value}

	if q.tail == nil {
		q.head = newNode
		q.tail = q.head
	} else {
		q.tail.next = newNode
		q.tail = q.tail.next
	}
}

func (q *queueList) Dequeue() (interface{}, bool) {
	if q.IsEmpty() {
		return nil, false
	}

	value := q.head.data
	if q.head == q.tail {
		q.tail = nil
	}
	q.head = q.head.next
	return value, true
}

func (q *queueList) DisplayElements() {
	current := q.head
	for current != nil {
		fmt.Print(current.data, ",")
		current = current.next
	}
	fmt.Println()
}

func (q *queueList) IsEmpty() bool {
	return q == nil
}
