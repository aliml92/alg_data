package array

func RemoveDuplicates(nums []int) int{
	if len(nums) == 0 {
		return 0
	}
	res := 1
	for i, j := 0, 1; j < len(nums); {
		if nums[i] != nums[j] {
			if nums[i+1] != nums[j] {
				nums[i+1] = nums[j]
			}
			res += 1
			i += 1
			j += 1 
		} else {
			j += 1
		}
	}
	return res
}