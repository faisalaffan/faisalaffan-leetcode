package main

// LeetCode #1446: Consecutive Characters
// https://leetcode.com/problems/consecutive-characters/
// Difficulty: Easy
//
// LeetCode submission: func maxPower(s string) int

import "fmt"

func main() {
	fmt.Println(ConsecutiveCharacters("leetcode")) // 2
	fmt.Println(ConsecutiveCharacters("abbcccddddeeeeedcba")) // 5
	fmt.Println(ConsecutiveCharacters("triplepillooooow")) // 5
}

// Time: O(n), Space: O(1)
func ConsecutiveCharacters(s string) int {
	ans, cur := 1, 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			cur++
			if cur > ans {
				ans = cur
			}
		} else {
			cur = 1
		}
	}
	return ans
}
