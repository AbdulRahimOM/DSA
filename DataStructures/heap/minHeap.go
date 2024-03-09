package heap

import (
	"DSA/tools"
	"fmt"
)

type MinHeap struct {
	array []int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{}
}

func (heap *MinHeap) PrintAsTree() {
	tools.PrintArrayAsBinaryTree(heap.array)
}

func (h *MinHeap) Insert(value int) {
	h.array = append(h.array, value)
	h.heapifyUp(len(h.array) - 1)
}
func (h *MinHeap) heapifyUp(index int) {
	for index > 0 {
		parentIndex := (index - 1) / 2
		if h.array[index] >= h.array[parentIndex] {
			break
		}
		h.array[index], h.array[parentIndex] = h.array[parentIndex], h.array[index]
		index = parentIndex
	}
}
func (h *MinHeap) ExtractMin() (int, error) {
	if len(h.array) == 0 {
		return 0, fmt.Errorf("heap is empty")
	}
	minValue := h.array[0]
	lastIndex := len(h.array) - 1
	h.array[0] = h.array[lastIndex]
	h.array = h.array[:lastIndex]
	h.heapifyDown(0, lastIndex-1)
	return minValue, nil
}
func (h *MinHeap) Delete(value int) bool {
	if len(h.array) == 0 {
		return false
	}
	for i := range h.array {
		if h.array[i] == value {
			lastIndex := len(h.array) - 1
			h.array[i] = h.array[lastIndex]
			h.array = h.array[:lastIndex]
			h.heapifyDown(i, lastIndex-1)
			return true
		}
	}
	return false
}

func (h *MinHeap) heapifyDown(parent int, lastIndex int) {
	leftChildIndex := 2*parent + 1
	rightChildIndex := 2*parent + 2
	var minIndex int

	if rightChildIndex <= lastIndex && h.array[rightChildIndex] < h.array[parent] {
		if h.array[leftChildIndex] < h.array[rightChildIndex] {
			minIndex = leftChildIndex
		} else {
			minIndex = rightChildIndex
		}
	} else if leftChildIndex <= lastIndex && h.array[leftChildIndex] < h.array[parent] {
		minIndex = leftChildIndex
	} else {
		return
	}
	h.array[parent], h.array[minIndex] = h.array[minIndex], h.array[parent]
	h.heapifyDown(minIndex, lastIndex)
}

func (h *MinHeap) BuildHeap(arr []int) {
	h.array = arr
	lastIndex := len(h.array) - 1
	for i := (lastIndex) / 2; i >= 0; i-- {
		h.heapifyDown(i, lastIndex)
	}
}


func HeapSortDesc(array []int) {
	h := &MinHeap{
		array: array,
	}
	lastIndex := len(array) - 1
	for i := (lastIndex - 1) / 2; i >= 0; i-- {
		h.heapifyDown(i, lastIndex)
	}

	for i := lastIndex; i >= 0; i-- {
		array[0], array[i] = array[i], array[0]

		h.heapifyDown(0, i-1)
	}

}