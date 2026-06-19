package main

// LeetCode #3351: Sum of Good Subsequences
// https://leetcode.com/problems/sum-of-good-subsequences/
// Difficulty: Hard
//
// A good subsequence is one where the absolute difference between consecutive
// elements is exactly 1. Single-element subsequences are trivially good.
// Return the sum of all elements across all good subsequences (mod 1e9+7).
//
// Approach: DP with hash map. For each value x, track:
// - cnt[x]: number of good subsequences ending with x
// - sum[x]: sum of all elements of good subsequences ending with x
// When processing x, it extends subsequences ending with x-1 or x+1.

import "fmt"

func main() {
	// Example 1
	fmt.Println(sumOfGoodSubsequences([]int{1, 2, 1}))
	// Example 2
	fmt.Println(sumOfGoodSubsequences([]int{3, 4, 5}))
	// Example 3
	fmt.Println(sumOfGoodSubsequences([]int{1, 1, 1}))
	// Edge: single element
	fmt.Println(sumOfGoodSubsequences([]int{5}))
	// Edge: alternating values
	fmt.Println(sumOfGoodSubsequences([]int{1, 3, 5}))
}

const mod = 1000000007

func sumOfGoodSubsequences(nums []int) int {
	cnt := make(map[int]int64)
	sum := make(map[int]int64)
	var ans int64

	for _, x := range nums {
		var newCnt int64 = 1 // the subsequence [x] alone
		var newSum int64

		// Extend subsequences ending with x-1 or x+1
		if c, ok := cnt[x-1]; ok {
			newCnt = (newCnt + c) % mod
			newSum = (newSum + sum[x-1]) % mod
		}
		if c, ok := cnt[x+1]; ok {
			newCnt = (newCnt + c) % mod
			newSum = (newSum + sum[x+1]) % mod
		}

		// Each subsequence that ends with x includes x as its last element
		newSum = (newSum + int64(x)*newCnt) % mod

		// Update DP maps
		cnt[x] = (cnt[x] + newCnt) % mod
		sum[x] = (sum[x] + newSum) % mod

		// Add to answer
		ans = (ans + newSum) % mod
	}

	return int(ans)
}
