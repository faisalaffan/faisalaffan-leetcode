package main

// LeetCode #3773: Maximum Number of Equal Length Runs
// https://leetcode.com/problems/maximum-number-of-equal-length-runs/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func maximumNumberOfEqualLengthRuns(s string) int {
	cnt := make(map[int]int)
	maxCount := 0
	n := len(s)
	i := 0
	for i < n {
		j := i
		for j < n && s[j] == s[i] {
			j++
		}
		runLen := j - i
		cnt[runLen]++
		if cnt[runLen] > maxCount {
			maxCount = cnt[runLen]
		}
		i = j
	}
	return maxCount
}

func main() {
	fmt.Println(maximumNumberOfEqualLengthRuns("hello"))
	fmt.Println(maximumNumberOfEqualLengthRuns("aaabaaa"))
	fmt.Println(maximumNumberOfEqualLengthRuns("aabbcc"))
}
