package main

// LeetCode #504: Base 7
// https://leetcode.com/problems/base-7/
// Difficulty: Easy

import "fmt"

// Time: O(log n), Space: O(log n)
func BaseSeven(num int) string {
	if num == 0 {
		return "0"
	}
	negative := num < 0
	if negative {
		num = -num
	}
	var result []byte
	for num > 0 {
		result = append([]byte{byte('0' + num%7)}, result...)
		num /= 7
	}
	if negative {
		result = append([]byte{'-'}, result...)
	}
	return string(result)
}

func main() {
	fmt.Println(BaseSeven(100))
	fmt.Println(BaseSeven(-7))
	fmt.Println(BaseSeven(0))
}
