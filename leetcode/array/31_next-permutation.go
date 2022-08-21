package array

import "sort"



func NextPermutation(nums []int) {
	n := len(nums)
	var pivot, swap int
	// find the pivot
	for i := n - 1; i >= 0; i-- {
		if nums[i-1] < nums[i] {
			pivot = i
			break
		}
	}
	if pivot == 0 {
		sort.Ints(nums)
		return
	}	
	// then find the swap which is the first number greater than the pivot
	swap = n - 1
	for nums[pivot-1] >= nums[swap] {
		swap--
	}	

	// swap
	nums[swap], nums[pivot-1] = nums[pivot-1], nums[swap]	

	// reverse from the pivot to the end
	sort.Ints(nums[pivot:])
}