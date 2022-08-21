package array


func SearchRange(nums []int, target int) []int {
	var res [2]int
	n := len(nums)
	res[0] = findLeft(nums, n, target)
	res[1] = findRight(nums, n, target)
	return res[:]
}


func findLeft(nums []int, n int, target int) int {
	left, right := 0, n-1
	for left <= right {
		mid := left + (right - left) / 2
		if (mid == 0 || nums[mid-1] < target) && nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}


func findRight(nums []int, n int, target int) int {
	left, right := 0, n-1
	for left <= right {
		mid := left + (right - left) / 2
		if (mid == n - 1 || nums[mid+1] > target) && nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
