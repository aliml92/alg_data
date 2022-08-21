package chapter04


func MergeSort(nums []int) []int {
	n := len(nums)
	if n < 2 {    // base case
		return nums
	}
	first  := MergeSort(nums[:n/2])
	second := MergeSort(nums[n/2:])
	return merge(first, second)
}

func merge(a []int, b []int) []int{
	var f []int
	for 0 < len(a) && 0 < len(b) {
		if a[0] <= b[0] {
			f = append(f, a[0])
			a = a[1:]
		} else {
			f = append(f, b[0])
			b = b[1:]
		}
	}
	for 0 < len(a) {
		f = append(f, a[0])
		a = a[1:]
	}
	for 0 < len(b) {
		f = append(f, b[0])
		b = b[1:]
	}
	return f 	 
}