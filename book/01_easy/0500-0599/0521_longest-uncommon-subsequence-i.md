# 0521 — Longest Uncommon Subsequence I

## Deskripsi

**Soal:** [0521. Longest Uncommon Subsequence I](https://leetcode.com/problems/longest-uncommon-subsequence-i/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(1), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func LongestUncommonSubsequenceI(a, b string) int`

## Solusi Go

```go
package main

// LeetCode #521: Longest Uncommon Subsequence I
// https://leetcode.com/problems/longest-uncommon-subsequence-i/
// Difficulty: Easy

import "fmt"

// Time: O(1), Space: O(1)
func LongestUncommonSubsequenceI(a, b string) int {
	if a == b {
		return -1
	}
	if len(a) > len(b) {
		return len(a)
	}
	return len(b)
}

func main() {
	fmt.Println(LongestUncommonSubsequenceI("aba", "cdc"))
	fmt.Println(LongestUncommonSubsequenceI("aaa", "bbb"))
	fmt.Println(LongestUncommonSubsequenceI("aaa", "aaa"))
}
```
