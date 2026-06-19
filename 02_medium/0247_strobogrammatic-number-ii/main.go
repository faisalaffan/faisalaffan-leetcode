package main

// LeetCode #247: Strobogrammatic Number II
// https://leetcode.com/problems/strobogrammatic-number-ii/
// Difficulty: Medium [Paid]
// Time: O(5^(n/2)), Space: O(n)

import "fmt"

func findStrobogrammatic(n int) []string {
	return build(n, n)
}

func build(n, final int) []string {
	if n == 0 {
		return []string{""}
	}
	if n == 1 {
		return []string{"0", "1", "8"}
	}

	inner := build(n-2, final)
	result := []string{}

	for _, s := range inner {
		if n != final {
			result = append(result, "0"+s+"0")
		}
		result = append(result, "1"+s+"1")
		result = append(result, "6"+s+"9")
		result = append(result, "8"+s+"8")
		result = append(result, "9"+s+"6")
	}

	return result
}

func main() {
	fmt.Println(findStrobogrammatic(2))
	fmt.Println(findStrobogrammatic(1))
	fmt.Println(findStrobogrammatic(3))
}
