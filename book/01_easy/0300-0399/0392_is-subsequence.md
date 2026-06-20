# 0392 — Is Subsequence

## Deskripsi

**Soal:** [0392. Is Subsequence](https://leetcode.com/problems/is-subsequence/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n+m), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func IsSubsequence(s, t string) bool`

## Solusi Go

```go
package main

// LeetCode #392: Is Subsequence
// https://leetcode.com/problems/is-subsequence/
// Difficulty: Easy

import "fmt"

// Time: O(n+m), Space: O(1)
func IsSubsequence(s, t string) bool {
	i := 0
	for j := 0; i < len(s) && j < len(t); j++ {
		if s[i] == t[j] {
			i++
		}
	}
	return i == len(s)
}

func main() {
	fmt.Println(IsSubsequence("abc", "ahbgdc"))
	fmt.Println(IsSubsequence("axc", "ahbgdc"))
	fmt.Println(IsSubsequence("", "ahbgdc"))
}
```
