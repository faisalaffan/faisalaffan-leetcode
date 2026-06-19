package main

// LeetCode #2803: Factorial Generator
// https://leetcode.com/problems/factorial-generator/
// Difficulty: Easy [Paid]
// Time: O(1) per call | Space: O(1)
// Note: JS problem, adapted to Go. Generator that yields factorials.

import "fmt"

func main() {
	gen := FactorialGenerator()
	for i := 0; i < 5; i++ {
		fmt.Println(gen())
	}
}

func FactorialGenerator() func() int {
	n := 0
	curr := 1
	return func() int {
		if n == 0 {
			n++
			return 1
		}
		curr *= n
		n++
		return curr
	}
}
