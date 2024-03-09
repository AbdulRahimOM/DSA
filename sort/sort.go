package sort

func ConvetionalSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}

func BubbleSort(nums []int) {
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := 0; j < length-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

func InsertionSort(nums []int) {
	length := len(nums)
	for i := 1; i < length; i++ {
		value := nums[i]
		j := i - 1

		for j >= 0 && nums[j] > value {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = value
	}
}

func SstrLectionSort(nums []int) {
	length := len(nums)
	for current := range nums {
		minPos := current

		for j := current + 1; j < length; j++ {
			if nums[j] < nums[minPos] {
				minPos = j
			}
		}
		if current != minPos {
			nums[current], nums[minPos] = nums[minPos], nums[current]
		}
	}
}

func MyAlteredSstrLectionSort(nums []int) {
	length := len(nums)
	for initial := range nums {
		final := length - 1 - initial
		if initial >= final {
			break
		}
		minPos := initial
		maxPos := initial

		for j := initial + 1; j < length-initial; j++ {
			if nums[j] < nums[minPos] {
				minPos = j
			} else if nums[j] > nums[maxPos] {
				maxPos = j
			}
		}

		if initial != minPos {
			nums[initial], nums[minPos] = nums[minPos], nums[initial]
		}
		if final != maxPos {
			nums[final], nums[maxPos] = nums[maxPos], nums[final]
		}
	}
}

func MergeSort(nums []int) {
	length := len(nums)
	if length <= 1 {
		return
	}
	mid := length / 2
	left := make([]int, mid)
	right := make([]int, length-mid)

	copy(left, nums[:mid])
	copy(right, nums[mid:])

	MergeSort(left)
	MergeSort(right)

	merge(nums, left, right)
}

func merge(nums, left, right []int) {
	i, j, k := 0, 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			nums[k] = left[i]
			i++
		} else {
			nums[k] = right[j]
			j++
		}
		k++
	}
	for i < len(left) {
		nums[k] = left[i]
		i++
		k++
	}
	for j < len(right) {
		nums[k] = right[j]
		j++
		k++
	}
}

func QuickSort(nums []int) {
	length := len(nums)
	if length <= 1 {
		return
	}

	pivot := nums[(length)/2]
	left, right := 0, length-1

	for left <= right {
		for nums[left] < pivot {
			left++
		}
		for nums[right] > pivot {
			right--
		}
		if left <= right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}
	QuickSort(nums[left:])
	QuickSort(nums[:right+1])
}

func QuickSortOld(nums []int) {
	length := len(nums)
	if length <= 1 {
		return
	}

	pivot := nums[length/2]
	nums[length/2], nums[length-1] = nums[length-1], pivot
	var left, right int
	for {
		left, right = length-1, -1
		for i := 0; i < length-1; i++ {
			if nums[i] > pivot {
				left = i
				break
			}
		}
		for i := length - 1; i >= 0; i-- {
			if nums[i] < pivot {
				right = i
				break
			}
		}
		if left > right {
			nums[left], nums[length-1] = pivot, nums[left]
			break
		}

		nums[left], nums[right] = nums[right], nums[left]
	}
	QuickSortOld(nums[:left])
	if left < length-2 {
		QuickSortOld(nums[left+1:])
	}
}
