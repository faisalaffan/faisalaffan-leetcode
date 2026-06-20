# 0214 — Shortest Palindrome

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan string. Tugasmu adalah memeriksa apakah string tersebut palindrome — dibaca sama dari depan dan belakang. Abaikan non-alfanumerik dan case.

**Cara berpikir:** Two Pointer — kiri dan kanan. Skip non-alfanumerik. Bandingkan.

**Fungsi Solusi:** `func shortestPalindrome(s string) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #214: Shortest Palindrome
// https://leetcode.com/problems/shortest-palindrome/
// Difficulty: Hard

import "fmt"

func shortestPalindrome(s string) string {
	n := len(s)
  // Edge case: input kosong
	if n == 0 {
		return ""
	}

	rev := make([]byte, n)
	for i := 0; i < n; i++ {
		rev[i] = s[n-1-i]
	}

	combined := s + "#" + string(rev)
  // Alokasi slice
	lps := make([]int, len(combined))

	for i := 1; i < len(combined); i++ {
		j := lps[i-1]
		for j > 0 && combined[i] != combined[j] {
			j = lps[j-1]
		}
		if combined[i] == combined[j] {
			j++
		}
		lps[i] = j
	}

	palLen := lps[len(lps)-1]
	suffix := rev[:n-palLen]
	return string(suffix) + s
}

func main() {
	fmt.Println(shortestPalindrome("aacecaaa"))
	fmt.Println(shortestPalindrome("abcd"))
}
```
