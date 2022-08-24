package array

// type stack []int

// func (s stack) Push(v int) stack {
//     return append(s, v)
// }

// func (s stack) Pop() (stack, int) {
//     l := len(s)
//     return  s[:l-1], s[l-1]
// }




func Permute(nums []int) [][]int {
	var res [][]int
	res = dfs(nums, res)
	return res
}


func dfs(nums []int, res [][]int) [][]int{
	if len(nums) == 1 {
		i := [][]int{nums[:]}
		return i
	}

	for i:= 0; i < len(nums); i++ {
		n := nums[0]
		nums = nums[1:]
		perms := dfs(nums, res)
		var temp [][]int
		for _,perm := range perms {
			perm = append(perm, n)
			temp = append(temp, perm)	
		res = append(res, temp...)
		nums = append(nums, n)	
		}
	}
	return res
}