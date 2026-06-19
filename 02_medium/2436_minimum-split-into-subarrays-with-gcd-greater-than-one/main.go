package main

// LeetCode #2436: Minimum Split Into Subarrays With GCD Greater Than One
// https://leetcode.com/problems/minimum-split-into-subarrays-with-gcd-greater-than-one/
// Difficulty: Medium
// Time: O(n) | Space: O(1)
// Greedy: start new subarray when current GCD becomes 1.

import "fmt"

func main() {
	fmt.Println(minimumSplits([]int{12, 6, 3, 14, 8})) // 2
	fmt.Println(minimumSplits([]int{4, 12, 6, 14}))    // 1
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func minimumSplits(nums []int) int {
	ans := 1
	cur := 0
	for _, v := range nums {
		cur = gcd(cur, v)
		if cur == 1 {
			ans++
			cur = v
		}
	}
	return ans
}
