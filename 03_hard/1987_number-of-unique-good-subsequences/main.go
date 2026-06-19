package main

// LeetCode #1987: Number of Unique Good Subsequences
// https://leetcode.com/problems/number-of-unique-good-subsequences/
// Difficulty: Hard
// Approach: DP tracking distinct good subsequences.
// "Good" means no leading zeros (except the single "0" itself).
// State:
//   zeroExists: whether we've seen the single "0" subsequence
//   ends1: number of good subsequences ending with '1'
//   ends0: number of good subsequences ending with '0' that start with '1' (no leading zero)
// For each '1': ends1 += ends1 + ends0 + 1
// For each '0': ends0 += ends1 + ends0; count "0" once

import "fmt"

const MOD1987 = 1000000007

func numberOfUniqueGoodSubsequences(binary string) int {
	zeroExists := false
	ends1 := 0
	ends0 := 0

	for _, ch := range binary {
		if ch == '1' {
			// Append '1' to all existing good subsequences, plus the single "1"
			ends1 = (ends1 + ends0 + 1) % MOD1987
		} else {
			// Append '0' to all good subsequences that already have a '1' (no leading zero)
			ends0 = (ends0 + ends1) % MOD1987
			if !zeroExists {
				zeroExists = true
			}
		}
	}

	ans := (ends1 + ends0) % MOD1987
	if zeroExists {
		ans = (ans + 1) % MOD1987
	}
	return ans
}

func main() {
	// Example: "001" -> 2 (subsequences: "0", "1")
	fmt.Println(numberOfUniqueGoodSubsequences("001"))

	// Additional tests
	fmt.Println(numberOfUniqueGoodSubsequences("101")) // "0","1","11","10","101" -> 5
	fmt.Println(numberOfUniqueGoodSubsequences("000")) // just "0" -> 1
	fmt.Println(numberOfUniqueGoodSubsequences("111")) // "1" -> 1
}
