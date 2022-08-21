package array

import (
	"sort"
)



func FourSum(nums []int, target int) [][]int{
	res := [][]int{} 
    sort.Ints(nums)
	for i := 0; i < len(nums)-3; i++  {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums) - 2; j++ {
			if j > i + 1 && nums[j] == nums[j-1] {
				continue
			}
			low := j + 1
			high := len(nums) - 1
			for low < high {
				sum := nums[i] + nums[j] + nums[low] + nums[high]
				if(sum == target) {
					a := []int{nums[i], nums[j], nums[low], nums[high]}
					res = append(res, a) 
					for low < high {
						if nums[low] == nums[low+1] {
							low++
						} else if nums[high] == nums[high-1] {
							high--
						} else {
							break
						}
					}
					low++
					high--
					// since we the slice is sorted, we should lower high
				} else if (sum < target) {
					low++
					// since we the slice is sorted, we should increase low
				} else {
					high--
				}
			}
		}
	}
	return res		
}




