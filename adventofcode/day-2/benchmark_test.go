package main

import (
	"os"
	"testing"
)


func BenchmarkSolutionWithStringsPkg(b *testing.B) {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for i := 0; i < b.N; i++ {
		result := solutionWithStringsPkg(file)
		_ = result
	}
}

func BenchmarkSolutionWithoutStringsPkg(b *testing.B) {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for i := 0; i < b.N; i++ {
		result := solutionWithoutStringsPkg(file)
		_ = result
	}
}
