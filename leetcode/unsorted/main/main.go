package main

import (
	"fmt"

	"github.com/aliml92/algdata/leetcode/array"
	// "github.com/aliml92/algdata/leetcode/array"
)


func main(){
	// beginWord := "hit"
	// endWord := "cog"
	// wordList := []string{"hot","dot","dog","lot","log","cog"}
	// r := unsorted.LadderLength(beginWord, endWord, wordList)
	// fmt.Println(r)
	// beginWord = "hot"
	// endWord = "dog"
	// wordList = []string{"hot","dog"}
	// r = unsorted.LadderLength(beginWord, endWord, wordList)
	// fmt.Println(r)
	// beginWord := "red"
	// endWord := "tax"
	// wordList := []string{"ted","tex","red","tax","tad","den","rex","pee"}
	// r := unsorted.LadderLength(beginWord, endWord, wordList)
	// fmt.Println(r)
	nums := []int{1,2,3}
	// n := nums[0]
	// fmt.Println(n)
	// nums = nums[1:]
	// a := [][]int{nums[:]}
	// fmt.Println(a)
	// var temp [][]int
	// for _, perm := range a {
	// 	perm = append(perm, n)
	// 	temp = append(temp, perm)
	// }
	// a = temp
	// fmt.Println(a)
	fmt.Println(array.Permute(nums))
}