package heap

import (
	"DSA/tools"
	"fmt"
)

type MaxHeap struct {
	array []int
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{}
}
func (h *MaxHeap) Insert(value int) {
	h.array = append(h.array, value)
	h.heapifyUp(len(h.array) - 1)
}

func (h *MaxHeap) heapifyUp(index int) {
	for index > 0 {
		parentIndex := (index - 1) / 2
		if h.array[index] <= h.array[parentIndex] {
			break
		}
		h.array[index], h.array[parentIndex] = h.array[parentIndex], h.array[index]
		index = parentIndex
	}
}
func (h *MaxHeap) ExtractMax() (int, error) {
	if len(h.array) == 0 {
		return 0, fmt.Errorf("heap is empty")
	}
	lastIndex := len(h.array) - 1
	maxValue := h.array[0]
	h.array[0] = h.array[lastIndex]
	h.array = h.array[:lastIndex]
	h.heapifyDown(0, lastIndex-1)
	return maxValue, nil
}
func (h *MaxHeap) Delete(value int) bool {
	if len(h.array) == 0 {
		return false
	}
	for i := range h.array {
		if h.array[i] == value {
			lastIndex := len(h.array)
			h.array[i] = h.array[lastIndex]
			h.array = h.array[:lastIndex]
			h.heapifyDown(i, lastIndex-1)
			return true
		}
	}
	return false
}

func (h *MaxHeap) BuildHeap(arr []int) {
	h.array = arr
	lastIndex := len(h.array) - 1
	for i := (lastIndex) / 2; i >= 0; i-- {
		h.heapifyDown(i, lastIndex)
	}
}
func (h *MaxHeap) PrintAsTree() {
	tools.PrintArrayAsBinaryTree(h.array)
}


func HeapSortAsc(array []int) {
	h := &MaxHeap{
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

func (h *MaxHeap) heapifyDown(parent int, indexLimit int) {
	var maxIndex int
	left := 2*parent + 1
	right := 2*parent + 2

	if right <= indexLimit && h.array[right] > h.array[parent] {
		if h.array[left] > h.array[right] {
			maxIndex = left
		} else {
			maxIndex = right
		}
	} else if left <= indexLimit && h.array[left] > h.array[parent] {
		maxIndex = left
	} else {
		return
	}
	
	h.array[parent], h.array[maxIndex] = h.array[maxIndex], h.array[parent]
	h.heapifyDown(maxIndex, indexLimit)
}