# 0647 — Palindromic Substrings

## Deskripsi

**Soal:** [0647. Palindromic Substrings](https://leetcode.com/problems/palindromic-substrings/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #647: Palindromic Substrings
// https://leetcode.com/problems/palindromic-substrings/
// Difficulty: Medium
// Time: O(n^2)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(countSubstrings("abc"))
	fmt.Println(countSubstrings("aaa"))
}

func countSubstrings(s string) int {
	count := 0
	n := len(s)

	for center := 0; center < 2*n-1; center++ {
		left := center / 2
		right := left + center%2
		for left >= 0 && right < n && s[left] == s[right] {
			count++
			left--
			right++
		}
	}

	return count
}
```
