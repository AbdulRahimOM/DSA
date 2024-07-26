package sort

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
