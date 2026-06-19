package main

// LeetCode #264: Ugly Number II
// https://leetcode.com/problems/ugly-number-ii/
// Difficulty: Medium
// Time: O(n), Space: O(n)

import "fmt"

func nthUglyNumber(n int) int {
	ugly := make([]int, n)
	ugly[0] = 1

	p2, p3, p5 := 0, 0, 0

	for i := 1; i < n; i++ {
		next := min(ugly[p2]*2, min(ugly[p3]*3, ugly[p5]*5))
		ugly[i] = next

		if next == ugly[p2]*2 {
			p2++
		}
		if next == ugly[p3]*3 {
			p3++
		}
		if next == ugly[p5]*5 {
			p5++
		}
	}

	return ugly[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(nthUglyNumber(10))
	fmt.Println(nthUglyNumber(1))
	fmt.Println(nthUglyNumber(7))
}
