package queue

import "fmt"

type QueueSlice struct {
	data []interface{}
}

func (q *QueueSlice) Enqueue(value interface{}) {
	q.data = append(q.data, value)
}

func (q *QueueSlice) Dequeue() (value interface{}, containsValue bool) {
	if len(q.data) == 0 {
		return
	}
	containsValue = true
	value = q.data[0]
	q.data = q.data[1:]
	return
}

func (q *QueueSlice) DisplayElements() {
	fmt.Println(q.data...)
}
func (q *QueueSlice)IsEmpty()bool{
	return len(q.data) == 0 
}