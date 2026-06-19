package main

// LeetCode #43: Multiply Strings
// https://leetcode.com/problems/multiply-strings/
// Difficulty: Medium

import "fmt"

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	m, n := len(num1), len(num2)
	result := make([]byte, m+n)

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			prod := (num1[i]-'0')*(num2[j]-'0') + result[i+j+1]
			result[i+j+1] = prod % 10
			result[i+j] += prod / 10
		}
	}

	if result[0] == 0 {
		result = result[1:]
	}

	for i := range result {
		result[i] += '0'
	}

	return string(result)
}

func main() {
	// Test case 1
	fmt.Println(multiply("2", "3")) // "6"

	// Test case 2
	fmt.Println(multiply("123", "456")) // "56088"

	// Test case 3
	fmt.Println(multiply("999", "999")) // "998001"
}

// Time: O(m*n) | Space: O(m+n)
