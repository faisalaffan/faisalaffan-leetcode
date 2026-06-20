# 1332 — Remove Palindromic Subsequences

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func removePalindromeSub(s string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

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
