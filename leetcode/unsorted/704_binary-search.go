package main


func Search(nums []int, target int) (int, int) {
	low, high := 0, len(nums) - 1
	i := 0
	for low <= high {
		i++
		mid := low + (high - low) / 2
		if (nums[mid] == target) {
			return mid, i
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1, i
}

