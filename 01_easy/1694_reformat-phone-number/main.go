package main

// LeetCode #1694: Reformat Phone Number
// https://leetcode.com/problems/reformat-phone-number/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(n)
func ReformatNumber(number string) string {
	digits := make([]byte, 0, len(number))
	for i := 0; i < len(number); i++ {
		if number[i] >= '0' && number[i] <= '9' {
			digits = append(digits, number[i])
		}
	}
	var result []byte
	i := 0
	for len(digits)-i > 4 {
		result = append(result, digits[i], digits[i+1], digits[i+2], '-')
		i += 3
	}
	remaining := len(digits) - i
	if remaining == 4 {
		result = append(result, digits[i], digits[i+1], '-', digits[i+2], digits[i+3])
	} else {
		result = append(result, digits[i:]...)
	}
	return string(result)
}

func main() {
	fmt.Println(ReformatNumber("1-23-45 6"))
	fmt.Println(ReformatNumber("123 4-567"))
	fmt.Println(ReformatNumber("123 4-5678"))
}
