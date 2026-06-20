# 2311 — Longest Binary Subsequence Less Than Or Equal To K

## Deskripsi

**Soal:** [2311. Longest Binary Subsequence Less Than Or Equal To K](https://leetcode.com/problems/longest-binary-subsequence-less-than-or-equal-to-k/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func longestSubsequence(s string, k int) int`

## Solusi Go

```go
package main

// LeetCode #2311: Longest Binary Subsequence Less Than or Equal to K
// https://leetcode.com/problems/longest-binary-subsequence-less-than-or-equal-to-k/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func longestSubsequence(s string, k int) int {
	val := 0
	count := 0
	pow := 1

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '0' {
			count++
		} else if s[i] == '1' {
			if val+pow <= k {
				val += pow
				count++
			}
		}
		if pow <= k {
			pow <<= 1
		}
	}
	return count
}

func main() {
	// Test case 1
	fmt.Println(longestSubsequence("1001010", 5))
	// Expected: 5

	// Test case 2
	fmt.Println(longestSubsequence("0010101010110101001001011010100010101010101010111111", 93))
	// Expected: 44
}
```
