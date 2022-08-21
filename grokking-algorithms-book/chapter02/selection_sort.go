package chapter02


import "fmt"


func selectionSort(nums []int) {
	n := len(nums)
	for i := 0; i < n - 1; i++ {
		min := i
		for j := i+1; j < n; j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		nums[i], nums[min] = nums[min], nums[i] 
	} 
}




func main(){
	sample := []int{3, 4, 5, 2, 1, 7, 8, -1, -3}
    fmt.Println(sample)
}