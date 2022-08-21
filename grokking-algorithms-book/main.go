package main

import (
	"fmt"

	ch4 "github.com/aliml92/algdata/grokking-algorithms-books/chapter04"
	ch6 "github.com/aliml92/algdata/grokking-algorithms-books/chapter06"
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
	
	graph := make(map[string][]string)
	graph["you"] = []string{"alice", "claire", "peggy"}
	graph["bob"] = []string{"anuj", "peggy"}
	graph["alice"] = []string{"peggy"}
	graph["claire"] = []string{"thom", "jonny"}
	graph["anuj"] = []string{}
	graph["peggy"] = []string{}
	graph["thom"] = []string{}
	graph["jonny"] = []string{}
	fmt.Println(ch6.HasSellerFor("you", graph))
}