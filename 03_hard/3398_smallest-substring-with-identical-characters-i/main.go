package main

// LeetCode #3398: Smallest Substring With Identical Characters I
// https://leetcode.com/problems/smallest-substring-with-identical-characters-i/
// Difficulty: Hard
//
// Binary search on max run length. Greedy check with runLen/(L+1) flips.

import "fmt"

func main() {
	fmt.Println(SmallestSubstringWithIdenticalCharactersI("000111", 1))
}

func SmallestSubstringWithIdenticalCharactersI(s string, k int) int {
	if k == 1 {
		return 1
	}

	n := len(s)
	lo, hi := 1, n
	for lo < hi {
		mid := lo + (hi-lo)/2
		if feasible(s, k, mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func feasible(s string, k, limit int) bool {
	cnt := 0
	i := 0
	n := len(s)
	for i < n {
		j := i
		for j < n && s[j] == s[i] {
			j++
		}
		runLen := j - i
		cnt += runLen / (limit + 1)
		if cnt > k {
			return false
		}
		i = j
	}
	return cnt <= k
}
