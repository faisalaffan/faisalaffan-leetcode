package main

// LeetCode #12: Integer to Roman
// https://leetcode.com/problems/integer-to-roman/
// Difficulty: Medium

import "fmt"

func intToRoman(num int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	result := ""
	for i := 0; i < len(values); i++ {
		for num >= values[i] {
			result += symbols[i]
			num -= values[i]
		}
	}

	return result
}

func main() {
	// Test case 1
	fmt.Println(intToRoman(3749)) // "MMMDCCXLIX"

	// Test case 2
	fmt.Println(intToRoman(58)) // "LVIII"

	// Test case 3
	fmt.Println(intToRoman(1994)) // "MCMXCIV"
}

// Time: O(1) | Space: O(1)
