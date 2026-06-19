package main

// LeetCode #2842: Count K-Subsequences of a String With Maximum Beauty
// https://leetcode.com/problems/count-k-subsequences-of-a-string-with-maximum-beauty/
// Difficulty: Hard
//
// Combinatorics. Beauty = sum of frequencies of characters in subsequence.
// To maximize beauty, pick the k most frequent distinct characters. Count ways
// = product over top k chars of (freq choose 1). Return mod 1e9+7.
// O(N + alphabet log alphabet) time, O(alphabet) space.

import (
	"fmt"
	"sort"
)

const mod2842 = 1000000007

func powMod2842(a, e int64) int64 {
	res := int64(1)
	a %= mod2842
	for e > 0 {
		if e&1 == 1 {
			res = (res * a) % mod2842
		}
		a = (a * a) % mod2842
		e >>= 1
	}
	return res
}

func countKSubsequencesWithMaxBeauty(s string, k int) int {
	if k > 26 {
		return 0
	}

	freq := make([]int, 26)
	for _, c := range s {
		freq[c-'a']++
	}

	freqs := make([]int, 0, 26)
	for _, f := range freq {
		if f > 0 {
			freqs = append(freqs, f)
		}
	}

	if len(freqs) < k {
		return 0
	}

	sort.Slice(freqs, func(i, j int) bool {
		return freqs[i] > freqs[j]
	})

	kthFreq := freqs[k-1]

	totalAtCutoff := 0
	for _, f := range freqs {
		if f == kthFreq {
			totalAtCutoff++
		}
	}

	greater := 0
	for _, f := range freqs {
		if f > kthFreq {
			greater++
		}
	}
	needFromCutoff := k - greater

	if needFromCutoff > totalAtCutoff {
		return 0
	}

	n := totalAtCutoff
	r := needFromCutoff
	if r > n-r {
		r = n - r
	}
	comb := int64(1)
	for i := 0; i < r; i++ {
		comb = comb * int64(n-i) % mod2842
		comb = comb * powMod2842(int64(i+1), mod2842-2) % mod2842
	}

	result := comb
	for _, f := range freqs {
		if f > kthFreq {
			result = result * int64(f) % mod2842
		}
	}
	for i := 0; i < needFromCutoff; i++ {
		result = result * int64(kthFreq) % mod2842
	}

	return int(result)
}

func main() {
	// Example: s="bcca", k=2 => 4
	fmt.Println(countKSubsequencesWithMaxBeauty("bcca", 2))

	// k=1
	fmt.Println(countKSubsequencesWithMaxBeauty("aabc", 1))

	// k > distinct chars
	fmt.Println(countKSubsequencesWithMaxBeauty("ab", 3))

	// k = distinct chars
	fmt.Println(countKSubsequencesWithMaxBeauty("abcd", 4))

	// All same char
	fmt.Println(countKSubsequencesWithMaxBeauty("aaaa", 1))
	fmt.Println(countKSubsequencesWithMaxBeauty("aaaa", 2))

	// Larger example with varying frequencies
	fmt.Println(countKSubsequencesWithMaxBeauty("abbcccddddeeeee", 3))

	// k=26 with many chars
	fmt.Println(countKSubsequencesWithMaxBeauty("abcdefghijklmnopqrstuvwxyz", 26))

	// k=0 (not valid per constraints but test edge)
	fmt.Println(countKSubsequencesWithMaxBeauty("abc", 0))

	// All same frequency, k equals number of chars
	fmt.Println(countKSubsequencesWithMaxBeauty("abc", 3))

	// Ties at cutoff
	fmt.Println(countKSubsequencesWithMaxBeauty("aabbccddee", 3))

	// Single character repeated many times
	fmt.Println(countKSubsequencesWithMaxBeauty("zzzzzzzzzz", 1))

	// k=2 with multiple chars at same frequency
	fmt.Println(countKSubsequencesWithMaxBeauty("aaabbbccc", 2))

	// Long string with uneven frequencies
	fmt.Println(countKSubsequencesWithMaxBeauty("thequickbrownfoxjumpsoverthelazydog", 5))
}
