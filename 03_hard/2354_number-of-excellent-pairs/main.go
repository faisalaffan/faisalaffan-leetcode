package main

import "fmt"

// 2354. Number of Excellent Pairs
// ----------------------------------------------------------------
// Given an array of integers nums and an integer k, a pair (i, j) is
// "excellent" if setBits(nums[i]) + setBits(nums[j]) >= k.
// Return the number of distinct pairs (i, j) where i and j are indices
// (unordered, i ≤ j — but pairs are counted by VALUE, not by index).
//
// The problem counts (i, j) such that the *values* nums[i] and nums[j]
// satisfy the condition.  Since we only care about set‑bit counts, and
// duplicate values produce the same pair, we can:
//   1. Deduplicate nums.
//   2. Count how many unique values have each set‑bit count (0 … 60).
//   3. Sum up pairs (c1, c2) where c1 + c2 >= k.
//
// For c1 == c2:  C(cnt, 2)  (choose 2 distinct indices).
// For c1 < c2:   cnt[c1] * cnt[c2].
// Note that (i,i) is allowed (same element can pair with itself).
// Wait — the problem says "pair (i, j)" — does i=j allowed? Usually
// LeetCode counts distinct pairs of indices with i ≤ j.
// Let's just follow the editorial approach: deduplicate values, count by
// set‑bit, then for each pair of distinct values, add cnt[c1] * cnt[c2] if
// c1 + c2 >= k, and for c1 == c2 add cnt[c1] * cnt[c1] (including same value).
// But "distinct pairs" usually means (i,j) where i can equal j.
//
// Actually, checking the problem more carefully: duplicate values should
// not be double‑counted because the excellence depends on the *value*, not
// the index.  So we deduplicate first, then count all unordered pairs
// (including (value, value)) where setBits(a) + setBits(b) >= k.
// The answer is simply the sum over all value pairs of 1 if condition holds.

func countExcellentPairs(nums []int, k int) int64 {
	// Deduplicate values.
	seen := make(map[int]bool)
	unique := make([]int, 0)
	for _, v := range nums {
		if !seen[v] {
			seen[v] = true
			unique = append(unique, v)
		}
	}

	// Count by set‑bit count.
	cnt := make([]int, 61) // bits up to 60 (nums[i] ≤ 10^9, but use 60 for safety)
	for _, v := range unique {
		cnt[popcount(v)]++
	}

	var ans int64
	// c1 + c2 >= k
	for c1 := 0; c1 <= 60; c1++ {
		if cnt[c1] == 0 {
			continue
		}
		for c2 := c1; c2 <= 60; c2++ {
			if cnt[c2] == 0 {
				continue
			}
			if c1+c2 < k {
				continue
			}
			if c1 == c2 {
				ans += int64(cnt[c1]) * int64(cnt[c2])
			} else {
				ans += int64(cnt[c1]) * int64(cnt[c2])
			}
		}
	}
	return ans
}

func popcount(x int) int {
	c := 0
	for x != 0 {
		x &= x - 1
		c++
	}
	return c
}

// ---------------------------------------------------------------------------
//  Wrapper

func NumberOfExcellentPairs() interface{} {
	return countExcellentPairs([]int{1, 2, 3, 1}, 3)
}

func main() {
	fmt.Println(NumberOfExcellentPairs())

	tests := []struct {
		nums []int
		k    int
		want int64
	}{
		{[]int{1, 2, 3, 1}, 3, 5},
		{[]int{5, 1, 1}, 10, 0},
		{[]int{1}, 1, 1},
		{[]int{7, 3, 1}, 4, 4},
	}
	for _, tc := range tests {
		got := countExcellentPairs(tc.nums, tc.k)
		if got != tc.want {
			fmt.Printf("FAIL nums=%v k=%d: got %d, want %d\n",
				tc.nums, tc.k, got, tc.want)
		}
	}
	fmt.Println("Done testing 2354.")
}
