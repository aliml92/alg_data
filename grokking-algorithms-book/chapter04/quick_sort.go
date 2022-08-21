package chapter04


func QuickSort(nums []int) []int {
	if len(nums) < 2 {   // base case
		return nums
	} else {
		pivot := nums[0]
		less, greater := split(nums[1:], pivot) 
		i := append(QuickSort(less),  pivot)
		i = append(i, QuickSort(greater)...)
		return i
	}

}



func split(nums []int, pivot int) ([]int, []int) {
	less, greater := []int{}, []int{}
	for _, v := range nums {
		if v <= pivot {
			less = append(less, v)
		} else {
			greater = append(greater, v)
		}
	}
	return less, greater
}