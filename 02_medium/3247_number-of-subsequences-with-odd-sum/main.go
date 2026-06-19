package main

// LeetCode #3247: Number of Subsequences with Odd Sum
// https://leetcode.com/problems/number-of-subsequences-with-odd-sum/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func subsequenceCount(nums []int) int {
	const mod = 1000000007
	n := len(nums)
	oddCount := 0
	for _, v := range nums {
		if v%2 != 0 {
			oddCount++
		}
	}

	if oddCount == 0 {
		return 0
	}

	pow := 1
	for i := 0; i < n-1; i++ {
		pow = (pow * 2) % mod
	}
	return pow
}

func main() {
	fmt.Println(subsequenceCount([]int{1, 2, 3})) // Expected: 4
	fmt.Println(subsequenceCount([]int{2, 4, 6})) // Expected: 0
	fmt.Println(subsequenceCount([]int{1}))        // Expected: 1
}
