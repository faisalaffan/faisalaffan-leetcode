package main

// LeetCode #2405: Optimal Partition of String
// https://leetcode.com/problems/optimal-partition-of-string/
// Difficulty: Medium
// Time: O(n) | Space: O(1)
// Greedy: start new partition when duplicate char found.

import "fmt"

func main() {
	fmt.Println(partitionString("abacaba")) // 4
	fmt.Println(partitionString("ssssss"))  // 6
}

func partitionString(s string) int {
	ans := 1
	seen := make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		if seen[s[i]] {
			ans++
			seen = make(map[byte]bool)
		}
		seen[s[i]] = true
	}
	return ans
}
