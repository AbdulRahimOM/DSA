package sort

import "DSA/DataStructures/heap"

func HeapSortAsc(array []int) {
	h := heap.NewMaxHeap()
	h.BuildHeap(array)

	for i := len(array)-1; i >= 0; i-- {
		array[0], array[i] = array[i], array[0]
		h.HeapifyDown(0, i-1)
	}
}

func HeapSortDesc(array []int) {
	h := heap.NewMinHeap()
	h.BuildHeap(array)

	for i := len(array)-1; i >= 0; i-- {
		array[0], array[i] = array[i], array[0]
		h.HeapifyDown(0, i-1)
	}
}
