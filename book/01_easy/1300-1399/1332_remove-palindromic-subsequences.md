# 1332 — Remove Palindromic Subsequences

## Deskripsi

**Soal:** [1332. Remove Palindromic Subsequences](https://leetcode.com/problems/remove-palindromic-subsequences/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func removePalindromeSub(s string) int`

## Solusi Go

```go
package main

// LeetCode #1332: Remove Palindromic Subsequences
// https://leetcode.com/problems/remove-palindromic-subsequences/
// Difficulty: Easy
//
// LeetCode submission: func removePalindromeSub(s string) int

import "fmt"

func main() {
	fmt.Println(RemovePalindromicSubsequences("ababa"))  // 1 (already palindrome)
	fmt.Println(RemovePalindromicSubsequences("abb"))    // 2
	fmt.Println(RemovePalindromicSubsequences("baabb"))  // 2
}

// Time: O(n), Space: O(1)
func RemovePalindromicSubsequences(s string) int {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return 2
		}
		i++
		j--
	}
	return 1
}
```
