package main

import (
	"fmt"

	ch4 "github.com/aliml92/algdata/grokking-algorithms-books/chapter04"
)


func main(){
	println("Recursive sum ")
	nums := []int{1,2,3,4,5}
	s := ch4.Sum(nums)
	fmt.Println(s)
	println("Merge Sort")
	nums = []int{10, 6, 2, 1, 5, 8, 3, 4, 7, 9}
	fmt.Println(ch4.MergeSort(nums))
	println("Quick Sort")
	nums = []int{10, 6, 2, 1, 5, 8, 3, 4, 7, 9}
	fmt.Println(ch4.QuickSort(nums)) 
}