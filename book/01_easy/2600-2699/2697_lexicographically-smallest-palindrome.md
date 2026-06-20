# 2697 — Lexicographically Smallest Palindrome

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu adalah memeriksa apakah string tersebut palindrome — dibaca sama dari depan dan belakang. Abaikan non-alfanumerik dan case.

**Cara berpikir:** Two Pointer — kiri dan kanan. Skip non-alfanumerik. Bandingkan.

**Fungsi Solusi:** `func LexicographicallySmallestPalindrome(s string) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #2697: Lexicographically Smallest Palindrome
// https://leetcode.com/problems/lexicographically-smallest-palindrome/
// Difficulty: Easy
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(LexicographicallySmallestPalindrome("egcfe"))
	fmt.Println(LexicographicallySmallestPalindrome("abcd"))
}

func LexicographicallySmallestPalindrome(s string) string {
	runes := []rune(s)
	i, j := 0, len(runes)-1
	for i < j {
		if runes[i] < runes[j] {
			runes[j] = runes[i]
		} else {
			runes[i] = runes[j]
		}
		i++
		j--
	}
	return string(runes)
}
```
