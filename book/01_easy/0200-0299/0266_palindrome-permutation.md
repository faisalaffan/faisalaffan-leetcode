# 0266 — Palindrome Permutation

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu adalah memeriksa apakah string tersebut palindrome — dibaca sama dari depan dan belakang. Abaikan non-alfanumerik dan case.

**Cara berpikir:** Two Pointer — kiri dan kanan. Skip non-alfanumerik. Bandingkan.

**Fungsi Solusi:** `func CanPermutePalindrome(s string) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(1) (fixed 256 chars)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #266: Palindrome Permutation
// https://leetcode.com/problems/palindrome-permutation/
// Difficulty: Easy [Paid]

import "fmt"

// Time: O(n) | Space: O(1) (fixed 256 chars)
func CanPermutePalindrome(s string) bool {
  // HashMap: O(1) lookup
	count := make(map[rune]int)
	for _, c := range s {
		count[c]++
	}
	oddCount := 0
	for _, v := range count {
		if v%2 == 1 {
			oddCount++
		}
	}
	return oddCount <= 1
}

func main() {
	fmt.Println(CanPermutePalindrome("code"))
	fmt.Println(CanPermutePalindrome("aab"))
	fmt.Println(CanPermutePalindrome("carerac"))
}
```
