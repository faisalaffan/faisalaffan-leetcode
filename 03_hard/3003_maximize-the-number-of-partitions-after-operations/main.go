package main

// LeetCode #3003: Maximize the Number of Partitions After Operations
// https://leetcode.com/problems/maximize-the-number-of-partitions-after-operations/
// Difficulty: Hard
//
// Approach: DP + memoization (DFS with bitmask)
// For each position, we can either continue the current partition or start
// a new one. Additionally, we have one "change" operation that lets us
// replace s[i] with any other character. We use a bitmask to track which
// distinct characters are in the current partition.

import (
	"fmt"
	"math/bits"
)

func maxPartitionsAfterOperations(s string, k int) int {
	n := len(s)
	memo := make(map[[3]int]int)
	var dfs func(i, mask int, changed bool) int
	dfs = func(i, mask int, changed bool) int {
		if i == n {
			return 1
		}
		key := [3]int{i, mask, 0}
		if changed {
			key[2] = 1
		}
		if v, ok := memo[key]; ok {
			return v
		}
		bit := 1 << (s[i] - 'a')
		newMask := mask | bit
		res := 0
		if bits.OnesCount(uint(newMask)) > k {
			res = dfs(i+1, bit, changed) + 1
		} else {
			res = dfs(i+1, newMask, changed)
		}
		if !changed {
			for j := 0; j < 26; j++ {
				newMask2 := mask | (1 << j)
				if bits.OnesCount(uint(newMask2)) > k {
					candidate := dfs(i+1, 1<<j, true) + 1
					if candidate > res {
						res = candidate
					}
				} else {
					candidate := dfs(i+1, newMask2, true)
					if candidate > res {
						res = candidate
					}
				}
			}
		}
		memo[key] = res
		return res
	}
	return dfs(0, 0, false)
}

func main() {
	// Example: "accca", k=2 -> 3
	fmt.Println(maxPartitionsAfterOperations("accca", 2))
	// Example: "aab", k=1 -> 3 (with change: "aab" -> "aac" gives partitions "a","a","c")
	fmt.Println(maxPartitionsAfterOperations("aab", 1))
	// Single char
	fmt.Println(maxPartitionsAfterOperations("a", 1))
}
