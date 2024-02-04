package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	result := solutionPartTwo(file)
	fmt.Println("result:", result)

}

func solutionWithStringsPkg(file *os.File) int {
	var result int
	m := map[string]int{
		"red": 12,
		"green": 13,
		"blue": 14,
	}

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		i++
		line := scanner.Text()
		s := strings.Split(line, ":")[1]
		sets := strings.Split(s, ";")	
		for _, set := range sets {
			setItems := strings.Split(set, ",")
			for _, setItem := range setItems {
				setItem := strings.TrimSpace(setItem)
				pair := strings.Split(setItem, " ")
				color := pair[1]
				num, _ := strconv.Atoi(pair[0])
				if num > m[color] {
					// stop scanning the current line as it does meet the requirement
					goto stopLineScan
				} 	
			}
		}

		result += i

		stopLineScan:
			// stop scanning the current line as it does not meet the requirement
	}

	return result
}

func solutionWithoutStringsPkg(file *os.File) int {
	var result int
	m := map[byte]int{
		'r': 12,
		'g': 13,
		'b': 14,
	}

	ml := map[byte]int{
		'r': 3,
		'g': 5,
		'b': 4,
	}

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		i++
		line := scanner.Bytes()
		j := 0
		scanStarted := false
		for j < len(line) {
			b := line[j]
			if b == ':' {
				scanStarted = true
			}

			if scanStarted {
				if '0' <= b && b <= '9' {
					nums := []byte{b}
					for k := j+1; k < len(line); k++ {
						if '0' <= line[k] && line[k] <= '9' {
							nums = append(nums, line[k])
						} else {
							break	
						}  
					}
					
					var num int
					coef := 1
					for x:= len(nums)-1; x >=0; x-- {
						num += coef * int(nums[x]-'0')
						coef *= 10
					}

					colorFirstByte := line[j + len(nums) + 1]
					
					if num > m[colorFirstByte] {
						goto stopLineScan
					}

					j +=  len(nums) + ml[colorFirstByte]
				}				
			}
			
			j++
		}	
		
		result += i
		
		stopLineScan:
			// stop scanning the line as it does not meet the requirement
	}

	return result
}

func solutionPartTwo(file *os.File) int {
	var result int
	m := map[string]int{
		"red": 0,
		"green": 0,
		"blue": 0,
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Split(line, ":")[1]
		sets := strings.Split(s, ";")	
		for _, set := range sets {
			setItems := strings.Split(set, ",")
			for _, setItem := range setItems {
				setItem := strings.TrimSpace(setItem)
				pair := strings.Split(setItem, " ")
				color := pair[1]
				num, _ := strconv.Atoi(pair[0])
				currNum := m[color]
				if num > currNum {
					m[color] = num
				}
			}
		}

		p := 1
		for k, v := range m {
			p *= v

			// reset values for the next lines
			m[k] = 0
		}

		result += p

	}

	return result
}