# 0409 — Longest Palindrome

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu adalah memeriksa apakah string tersebut palindrome — dibaca sama dari depan dan belakang. Abaikan non-alfanumerik dan case.

**Cara berpikir:** Two Pointer — kiri dan kanan. Skip non-alfanumerik. Bandingkan.

**Fungsi Solusi:** `func LongestPalindrome(s string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #409: Longest Palindrome
// https://leetcode.com/problems/longest-palindrome/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func LongestPalindrome(s string) int {
	count := [128]int{}
	for _, c := range s {
		count[c]++
	}
	length, odd := 0, 0
	for _, c := range count {
		length += (c / 2) * 2
		if c%2 == 1 {
			odd = 1
		}
	}
	return length + odd
}

func main() {
	fmt.Println(LongestPalindrome("abccccdd"))
	fmt.Println(LongestPalindrome("a"))
	fmt.Println(LongestPalindrome("bb"))
}
```
