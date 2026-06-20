# 0214 — Shortest Palindrome

## Deskripsi

**Soal:** [0214. Shortest Palindrome](https://leetcode.com/problems/shortest-palindrome/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

**Fungsi Solusi:** `func shortestPalindrome(s string) string`

## Solusi Go

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

  // Membuat slice untuk menyimpan hasil
	rev := make([]byte, n)
	for i := 0; i < n; i++ {
		rev[i] = s[n-1-i]
	}

	combined := s + "#" + string(rev)
  // Membuat slice untuk menyimpan hasil
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
