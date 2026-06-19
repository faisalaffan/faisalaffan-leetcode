package main

// LeetCode #2949: Count Beautiful Substrings II
// https://leetcode.com/problems/count-beautiful-substrings-ii/
// Difficulty: Hard
//
// Approach: Prefix sum + modulo arithmetic + hash map.
// A substring is beautiful if:
//   1. vowels == consonants (balanced)
//   2. (vowels * consonants) % k == 0, i.e., v^2 % k == 0 (since v == c)
//
// For condition 2: if v^2 % k == 0, then v % p == 0 where p is derived
// from factorization of k: for each prime factor q^e of k, p gets q^ceil(e/2).
//
// Track (prefix_diff, (index+1) % p) as state in hash map.
// Each matching pair (l, r) with same state forms a beautiful substring.

import "fmt"

func isVowel(ch byte) bool {
	switch ch {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	}
	return false
}

func beautifulSubstrings(s string, k int) int64 {
	// Compute p: smallest such that p^2 % (4k) == 0
	k4 := k * 4
	p := 1
	for i := 2; i*i <= k4; i++ {
		if k4%i == 0 {
			cnt := 0
			for k4%i == 0 {
				k4 /= i
				cnt++
			}
			for j := 0; j < (cnt+1)/2; j++ {
				p *= i
			}
		}
	}
	if k4 > 1 {
		p *= k4
	}

	// State: (diff, idx_mod_p) -> count
	type state struct {
		diff int
		mod  int
	}
	counts := make(map[state]int64)
	counts[state{diff: 0, mod: 0}] = 1

	var ans int64
	diff := 0
	for i := 0; i < len(s); i++ {
		if isVowel(s[i]) {
			diff++
		} else {
			diff--
		}
		st := state{diff: diff, mod: (i + 1) % p}
		ans += counts[st]
		counts[st]++
	}
	return ans
}

func main() {
	// Example: s="baeyh", k=2 -> 5
	fmt.Println(beautifulSubstrings("baeyh", 2))

	// Empty vowels
	fmt.Println(beautifulSubstrings("bcdf", 1))

	// Simple case
	fmt.Println(beautifulSubstrings("ab", 1))
}
