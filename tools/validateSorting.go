package tools

func CheckOrder(nums []int) bool {
	length := len(nums)
	if length <= 1 {
		return true
	}
	max := nums[0]
	for i := 1; i < length; i++ {
		if nums[i] < max {
			return false
		} else {
			max = nums[i]
		}
	}
	return true
}
