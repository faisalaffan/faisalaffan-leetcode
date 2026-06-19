package main

// LeetCode #3003: Maximize the Number of Partitions After Operations
// https://leetcode.com/problems/maximize-the-number-of-partitions-after-operations/
// Difficulty: Hard
//
// We have a string s and an integer k.
// We can change AT MOST ONE character in s to any lowercase letter.
// We then partition the string into the maximum number of substrings
// such that each substring has at most k distinct characters.
//
// Approach: DP + memoization (DFS with bitmask)
// For each position, we track the current partition's bitmask of distinct chars.
// We can either continue the current partition or start a new one.
// Additionally, we have one "change" operation we may use to replace s[i]
// with any other character to potentially increase partition count.

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

		// Option 1: Keep the character as is
		if bits.OnesCount(uint(newMask)) > k {
			// Need to start a new partition
			res = dfs(i+1, bit, changed) + 1
		} else {
			res = dfs(i+1, newMask, changed)
		}

		// Option 2: Use the change operation (only if not used yet)
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
	// Example 1: "accca", k=2 -> 3
	// Without change: "accca" -> distinct={a,c} -> 1 partition
	// With change: "accca" -> change 'c' at pos 2 to 'b' -> "acbca"
	//   partitions: "ac" (a,c) | "b" (b) | "ca" (c,a) -> 3
	fmt.Println("Test 1:", maxPartitionsAfterOperations("accca", 2))

	// Example 2: "aab", k=1 -> 3
	// Without change: "a" | "a" | "b" -> 3 (each partition has at most 1 distinct)
	fmt.Println("Test 2:", maxPartitionsAfterOperations("aab", 1))

	// Single char
	fmt.Println("Test 3:", maxPartitionsAfterOperations("a", 1))

	// All same characters
	fmt.Println("Test 4:", maxPartitionsAfterOperations("aaaa", 1))

	// All distinct
	fmt.Println("Test 5:", maxPartitionsAfterOperations("abcdef", 1))

	fmt.Println("Test 6:", maxPartitionsAfterOperations("aba", 1))
}
