package main

// LeetCode #3179: Find the N-th Value After K Seconds
// https://leetcode.com/problems/find-the-n-th-value-after-k-seconds/
// Difficulty: Medium
// Time: O(n * k) | Space: O(n)

import "fmt"

func valueAfterKSeconds(n int, k int) int {
	const mod = 1_000_000_007
	arr := make([]int, n)
	for i := range arr {
		arr[i] = 1
	}

	for s := 0; s < k; s++ {
		for i := 1; i < n; i++ {
			arr[i] = (arr[i] + arr[i-1]) % mod
		}
	}
	return arr[n-1]
}

func main() {
	fmt.Println(valueAfterKSeconds(4, 5)) // Expected: 56
	fmt.Println(valueAfterKSeconds(5, 3)) // Expected: 35
}
