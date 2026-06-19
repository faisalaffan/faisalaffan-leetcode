package main

// LeetCode #1044: Longest Duplicate Substring
// https://leetcode.com/problems/longest-duplicate-substring/
// Difficulty: Hard
//
// Approach: Binary search on length + Rabin-Karp rolling hash.
//   - Binary search for the longest length L that has a duplicate substring.
//   - For a given length, use rolling hash (base 26, mod large int) to find duplicates.
//   - Use two moduli to minimize collision probability.

import "fmt"

func main() {
	fmt.Println(longestDupSubstring("banana")) // "ana"
	fmt.Println(longestDupSubstring("abcd"))   // ""
}

func longestDupSubstring(s string) string {
	n := len(s)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = int(s[i] - 'a')
	}

	mod1 := int64(1_000_000_007)
	mod2 := int64(1_000_000_009)
	base := int64(26)

	// verify an actual substring match (handles rare hash collisions)
	verify := func(i, j, length int) bool {
		for k := 0; k < length; k++ {
			if s[i+k] != s[j+k] {
				return false
			}
		}
		return true
	}

	// find duplicate substring of given length; returns start index or -1
	find := func(length int) int {
		if length == 0 {
			return -1
		}
		pow1 := int64(1)
		pow2 := int64(1)
		for i := 1; i < length; i++ {
			pow1 = (pow1 * base) % mod1
			pow2 = (pow2 * base) % mod2
		}

		hash1 := int64(0)
		hash2 := int64(0)
		for i := 0; i < length; i++ {
			hash1 = (hash1*base + int64(nums[i])) % mod1
			hash2 = (hash2*base + int64(nums[i])) % mod2
		}

		seen := make(map[[2]int64]int)
		seen[[2]int64{hash1, hash2}] = 0

		for i := length; i < n; i++ {
			hash1 = ((hash1 - pow1*int64(nums[i-length])%mod1 + mod1) % mod1)
			hash1 = (hash1*base + int64(nums[i])) % mod1

			hash2 = ((hash2 - pow2*int64(nums[i-length])%mod2 + mod2) % mod2)
			hash2 = (hash2*base + int64(nums[i])) % mod2

			key := [2]int64{hash1, hash2}
			if start, ok := seen[key]; ok {
				if verify(start, i-length+1, length) {
					return start
				}
			}
			seen[key] = i - length + 1
		}
		return -1
	}

	lo, hi := 1, n
	start := -1
	for lo <= hi {
		mid := (lo + hi) / 2
		if pos := find(mid); pos != -1 {
			start = pos
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	if start == -1 {
		return ""
	}
	return s[start : start+hi]
}
