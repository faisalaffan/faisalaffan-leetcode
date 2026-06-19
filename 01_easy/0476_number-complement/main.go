package main

// LeetCode #476: Number Complement
// https://leetcode.com/problems/number-complement/
// Difficulty: Easy

import "fmt"

// Time: O(1), Space: O(1)
func NumberComplement(num int) int {
	mask := ^0
	for num&mask != 0 {
		mask <<= 1
	}
	return ^num & ^mask
}

func main() {
	fmt.Println(NumberComplement(5))
	fmt.Println(NumberComplement(1))
	fmt.Println(NumberComplement(2))
}
