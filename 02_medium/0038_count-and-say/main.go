package main

// LeetCode #38: Count and Say
// https://leetcode.com/problems/count-and-say/
// Difficulty: Medium

import "fmt"

func countAndSay(n int) string {
	curr := "1"

	for i := 2; i <= n; i++ {
		var next []byte
		count := 1
		for j := 1; j < len(curr); j++ {
			if curr[j] == curr[j-1] {
				count++
			} else {
				next = append(next, byte('0'+count), curr[j-1])
				count = 1
			}
		}
		next = append(next, byte('0'+count), curr[len(curr)-1])
		curr = string(next)
	}

	return curr
}

func main() {
	// Test case 1
	fmt.Println(countAndSay(4)) // "1211"
	fmt.Println(countAndSay(1)) // "1"
	fmt.Println(countAndSay(5)) // "111221"
}

// Time: O(2^n) | Space: O(2^n)
