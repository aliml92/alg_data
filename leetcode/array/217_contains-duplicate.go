package array


func ContainsDuplicate(nums []int) bool {
	set := make(map[int]bool)
	for _, i := range nums {
		_, ok := set[i]
		if ok {
			return true
		}
		set[i] = true
	}
	return false
}