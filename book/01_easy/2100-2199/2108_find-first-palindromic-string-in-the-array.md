# 2108 — Find First Palindromic String In The Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func FindFirstPalindromicStringInTheArray(words []string) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n * m), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2108: Find First Palindromic String in the Array
// https://leetcode.com/problems/find-first-palindromic-string-in-the-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindFirstPalindromicStringInTheArray([]string{"abc", "car", "ada", "racecar", "cool"})) // "ada"
	fmt.Println(FindFirstPalindromicStringInTheArray([]string{"notapalindrome", "racecar"}))             // "racecar"
	fmt.Println(FindFirstPalindromicStringInTheArray([]string{"def", "ghi"}))                             // ""
}

// Time: O(n * m), Space: O(1)
func FindFirstPalindromicStringInTheArray(words []string) string {
	for _, w := range words {
		if isPalindrome(w) {
			return w
		}
	}
	return ""
}

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
```
