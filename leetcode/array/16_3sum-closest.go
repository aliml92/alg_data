package array

import (
	"sort"
)


func ThreeSumClosest(nums []int, target int) int {
	var closest, diff, d int
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if (i == 0 || (i > 0 && nums[i] != nums[i-1])) {
			low := i+1 
            high := len(nums)-1
			for (low < high) {
				sum := nums[i] + nums[low] + nums[high]
				if (sum < target) {
					d = target - sum
					low++
				} else if (sum > target) {
					d = sum - target
					high--
				} else {
					return sum
				}
				if (d < diff) || diff == 0 {
					diff = d
					closest = sum
				}
			}
		}
	}
	return closest

}


