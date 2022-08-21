package chapter04



func Sum(nums []int) int {
	if len(nums) == 1 {    // base case
		return nums[0]
	} else {               // recursive part 
		return nums[0] + Sum(nums[1:])		
	}
}
