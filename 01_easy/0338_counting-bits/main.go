package main

// LeetCode #338: Counting Bits
// https://leetcode.com/problems/counting-bits/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(n)
func CountingBits(n int) []int {
	ans := make([]int, n+1)
	for i := 1; i <= n; i++ {
		ans[i] = ans[i>>1] + (i & 1)
	}
	return ans
}

func main() {
	fmt.Println(CountingBits(2))
	fmt.Println(CountingBits(5))
	fmt.Println(CountingBits(0))
}
