# 1147 — Longest Chunked Palindrome Decomposition

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan string. Tugasmu adalah memeriksa apakah string tersebut palindrome — dibaca sama dari depan dan belakang. Abaikan non-alfanumerik dan case.

**Cara berpikir:** Two Pointer — kiri dan kanan. Skip non-alfanumerik. Bandingkan.

**Fungsi Solusi:** `func longestDecomposition(text string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Prefix Sum

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Prefix Sum** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1147: Longest Chunked Palindrome Decomposition
// https://leetcode.com/problems/longest-chunked-palindrome-decomposition/
// Difficulty: Hard
//
// Greedy two-pointer: try the shortest matching prefix/suffix pair. When a
// match is found, increment count by 2 and advance both pointers. Any
// unmatched remnant in the middle adds 1.

import "fmt"

func main() {
	fmt.Println(longestDecomposition("ghiabcdefhelloadamhelloabcdefghi"))
}

func longestDecomposition(text string) int {
	n := len(text)
	ans := 0
	l, r := 0, n-1

	for l <= r {
		found := false
		for length := 1; l+length-1 < r-length+1; length++ {
			if text[l:l+length] == text[r-length+1:r+1] {
				ans += 2
				l += length
				r -= length
				found = true
				break
			}
		}
		if !found {
			ans++
			break
		}
	}

	return ans
}
```
