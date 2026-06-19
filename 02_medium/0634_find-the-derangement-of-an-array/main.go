package main

// LeetCode #634: Find the Derangement of An Array
// https://leetcode.com/problems/find-the-derangement-of-an-array/
// Difficulty: Medium [Paid]
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(FindDerangement(3))
	fmt.Println(FindDerangement(4))
}

func FindDerangement(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 0
	}

	const mod = 1_000_000_007
	a, b := 0, 1 // D(1)=0, D(2)=1

	for i := 3; i <= n; i++ {
		c := ((i - 1) * (a + b)) % mod
		a, b = b, c
	}

	return b
}
