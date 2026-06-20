# 3713 — Longest Balanced Substring I

## Deskripsi

**Soal:** [3713. Longest Balanced Substring I](https://leetcode.com/problems/longest-balanced-substring-i/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func longestBalancedSubstringI(s string) int`

## Solusi Go

```go
package main

// LeetCode #3713: Longest Balanced Substring I
// https://leetcode.com/problems/longest-balanced-substring-i/
// Difficulty: Medium
// Time: O(n^2) | Space: O(1)

import "fmt"

func longestBalancedSubstringI(s string) int {
	n := len(s)
	ans := 0
	for i := 0; i < n; i++ {
		var cnt [26]int
		distinct := 0
		maxFreq := 0
		for j := i; j < n; j++ {
			idx := s[j] - 'a'
			cnt[idx]++
			if cnt[idx] == 1 {
				distinct++
			}
			if cnt[idx] > maxFreq {
				maxFreq = cnt[idx]
			}
			if maxFreq*distinct == j-i+1 {
				if j-i+1 > ans {
					ans = j - i + 1
				}
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(longestBalancedSubstringI("abbac"))
	fmt.Println(longestBalancedSubstringI("zzabccy"))
	fmt.Println(longestBalancedSubstringI("aba"))
}
```
