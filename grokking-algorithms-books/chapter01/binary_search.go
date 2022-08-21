package chapter01

// time complexity  O(logn)
// space complexity O(1)
func search(nums []int, target int) int {
	low, high := 0, len(nums) - 1
	i := 0
	for low <= high {
		i++
		mid := low + (high - low) / 2
		if (nums[mid] == target) {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}