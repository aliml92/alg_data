package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	defer func() {
		if r := recover(); r != nil {
			printStackTrace()
			fmt.Println("Recovered from panic:", r)
		}
	}()

	result := dayOnePartTwo(file)

	fmt.Printf("result: %d", result)

}

func dayOnePartOne(file *os.File) int {
	var result int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var digits []byte
		for _, b := range scanner.Bytes() {
			if '0' <= b && b <= '9' {
				digits = append(digits, b)
			}
		}
		l := len(digits)
		if l > 0 {
			left := digits[0]
			right := digits[l-1]
			result = result + int(left-'0') * 10  + int(right-'0')
		} 
	}

	return result
}

func dayOnePartTwo(file *os.File) int {
	var res int
	
	m := map[byte][]string{
		'o': {"one"},
		't': {"two", "three"},
		'f': {"four", "five"},
		's': {"six", "seven"},
		'e': {"eight"}, 
		'n': {"nine"},
	}

	md := map[string]int{
		"one": 1,
		"two": 2,
		"three": 3,
		"four": 4,
		"five": 5,
		"six": 6,
		"seven": 7,
		"eight": 8,
		"nine": 9,
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		
		var result []int
		
		text := scanner.Text()
		i := 0
		for i < len(text) {
			b := text[i]

			// check if we found the first letter of a digit word
			digits := m[b]
			if digits != nil {
				for _, digit := range digits {
					if i+len(digit) <= len(text) {
						target := text[i:i+len(digit)]
						if target == digit {
							result = append(result, md[target])
							i++
							goto exit
						}
					}
				}
			}

			// check if we found a digit
			if '0' <= b && b <= '9' {
				result = append(result, int(b-'0'))
			}
			
			i++

			exit:
				// exit immediately in case digit word detected 
		}
		
		// we are done checking the whole line
		// get the first and last digit
		l := len(result)
		if l > 0 {
			left := result[0]
			right := result[l-1]
			res = res + left * 10 + right
		}
	}

	return res
}


func printStackTrace() {
	buf := make([]byte, 1<<16)
	stackSize := runtime.Stack(buf, false)
	fmt.Printf("=== BEGIN STACK TRACE ===\n%s\n=== END STACK TRACE ===\n", buf[:stackSize])
}
