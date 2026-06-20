package main

// LeetCode #3655: XOR After Range Multiplication Queries II
// https://leetcode.com/problems/xor-after-range-multiplication-queries-ii/
// Difficulty: Hard
//
// Given array nums and queries [l, r, k, v], multiply nums[i] by v for all
// positions i = l, l+k, l+2k, ... <= r. Return XOR of final array.
//
// Approach: Use difference array to track multiplier per position,
// apply queries efficiently with batch processing.

import "fmt"

func main() {
	// Example 1
	fmt.Println(xorAfterQueries([]int{1, 2, 3, 4}, [][]int{{0, 3, 1, 2}}))
	// Example 2
	fmt.Println(xorAfterQueries([]int{5, 3, 7}, [][]int{{0, 2, 1, 3}, {1, 1, 1, 2}}))
	// Edge: single element
	fmt.Println(xorAfterQueries([]int{10}, [][]int{{0, 0, 1, 5}}))
}

const MOD = 1000000007

func xorAfterQueries(nums []int, queries [][]int) int {
	n := len(nums)
	// Track multiplier per position
	mult := make([]int64, n)
	for i := range mult {
		mult[i] = 1
	}

	for _, q := range queries {
		l, r, k, v := q[0], q[1], q[2], q[3]
		for i := l; i <= r; i += k {
			mult[i] = (mult[i] * int64(v)) % MOD
		}
	}

	var result int64
	for i := 0; i < n; i++ {
		val := (int64(nums[i]) * mult[i]) % MOD
		result ^= val
	}
	return int(result)
}
