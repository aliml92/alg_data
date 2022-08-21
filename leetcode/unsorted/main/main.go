package main

import (
	"fmt"

	"github.com/aliml92/algdata/leetcode/unsorted"
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
	beginWord := "red"
	endWord := "tax"
	wordList := []string{"ted","tex","red","tax","tad","den","rex","pee"}
	r := unsorted.LadderLength(beginWord, endWord, wordList)
	fmt.Println(r)
}