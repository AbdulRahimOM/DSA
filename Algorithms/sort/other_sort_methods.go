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

