package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	result := solution(file)
	fmt.Println("result:", result)

}

func solution(f *os.File) int {
	var result int
	var twoD [][]byte

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {	
		twoD = append(twoD, scanner.Bytes())
	}
	
	rowLen := len(twoD)
	colLen := len(twoD[0])


	checkPartNum := func (i, j, numLen int) bool {
		// check left
		if j-1 >=0 && twoD[i][j-1] != '.'  { 
			return true
		}

		// check right
		if j+ numLen < colLen && twoD[i][j+numLen] != '.' {
			return true
		}

		// check top
		if i-1 >=0 {
			for x := j; x < j+numLen; x++ {
				if twoD[i-1][x] != '.' && ( twoD[i-1][x] < '0' || '9' < twoD[i-1][x]) {
					return true
				}
			}

			if j-1 >= 0 && (twoD[i-1][j-1] != '.' && ( twoD[i-1][j-1] < '0' || '9' < twoD[i-1][j-1])) {
				return true
			} 

			if j+numLen < colLen && (twoD[i-1][j+numLen] != '.' && ( twoD[i-1][j+numLen] < '0' || '9' < twoD[i-1][j+numLen])) {
				return true
			} 
		} 

		// check bottom
		if i+1 < rowLen {
			for x := j; x < j+numLen; x++ {
				if twoD[i+1][x] != '.' && ( twoD[i+1][x] < '0' || '9' < twoD[i+1][x]) {
					return true
				}
			}

			if j-1 >= 0 && (twoD[i+1][j-1] != '.' && ( twoD[i+1][j-1] < '0' || '9' < twoD[i+1][j-1])) {
				return true
			} 

			if j+numLen < colLen && (twoD[i+1][j+numLen] != '.' && ( twoD[i+1][j+numLen] < '0' || '9' < twoD[i+1][j+numLen])) {
				return true
			} 
		}

		return false

	}

	for i:=0; i < rowLen; i++ {

		j := 0
		for j < colLen {
			item := twoD[i][j]
			if '0' <= item && item <= '9' {
				nums := []byte{item}	
				for k := j+1; k < colLen; k++ {
					nextItem := twoD[i][k] 
					if '0' <= nextItem && nextItem <= '9' {
						nums = append(nums, nextItem)
					} else {
						break
					}
				}

				// check if this nums found is a 'part number'
				if checkPartNum(i, j, len(nums)) {
					// construct num out of nums of byte slice
					var num int
					coef := 1
					for x:= len(nums)-1; x >=0; x-- {
						num += coef * int(nums[x]-'0')
						coef *= 10
					}

					result += num
				}

				j += len(nums) 
			}

			j++
		} 
	}

	return result
}