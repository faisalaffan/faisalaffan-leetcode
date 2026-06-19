package main

// LeetCode #2380: Time Needed to Rearrange a Binary String
// https://leetcode.com/problems/time-needed-to-rearrange-a-binary-string/
// Difficulty: Medium
// Time: O(n) | Space: O(1)
// Each "01" becomes "10" per second. Equivalent to: each 1 moves right past 0s,
// and the total time is max over each 1 of (position - index_in_final_position).

import "fmt"

func main() {
	fmt.Println(secondsToRemoveOccurrences("0110101")) // 4
	fmt.Println(secondsToRemoveOccurrences("11100"))   // 0
	fmt.Println(secondsToRemoveOccurrences("001011"))  // 3
}

func secondsToRemoveOccurrences(s string) int {
	ans, zeros := 0, 0
	for _, ch := range s {
		if ch == '0' {
			zeros++
		} else if zeros > 0 {
			ans = max(ans+1, zeros)
		}
	}
	return ans
}
