# 1177 — Can Make Palindrome From Substring

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu adalah memeriksa apakah string tersebut palindrome — dibaca sama dari depan dan belakang. Abaikan non-alfanumerik dan case.

**Cara berpikir:** Two Pointer — kiri dan kanan. Skip non-alfanumerik. Bandingkan.

**Fungsi Solusi:** `func canMakePaliQueries(s string, queries [][]int) []bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, Prefix Sum

**Waktu:** O(n * 26 + m) where m = len(queries)  |  **Ruang:** O(n * 26) — prefix sums per character

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #1177: Can Make Palindrome from Substring
// https://leetcode.com/problems/can-make-palindrome-from-substring/
// Difficulty: Medium

// For each query [left, right, k], check if we can rearrange
// substring s[left..right] into a palindrome with at most k replacements.
// A palindrome can have at most 1 odd count character.
// We need floor(oddCount/2) <= k replacements.

// Time: O(n * 26 + m) where m = len(queries)
// Space: O(n * 26) — prefix sums per character

func canMakePaliQueries(s string, queries [][]int) []bool {
	n := len(s)
	// prefix[i][c] = count of char c in s[0:i]
  // Alokasi slice
	prefix := make([][26]int, n+1)

	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i]
		prefix[i+1][s[i]-'a']++
	}

	result := make([]bool, len(queries))
	for idx, q := range queries {
		l, r, k := q[0], q[1]+1, q[2]
		oddCount := 0
		for c := 0; c < 26; c++ {
			count := prefix[r][c] - prefix[l][c]
			if count%2 == 1 {
				oddCount++
			}
		}
		result[idx] = oddCount/2 <= k
	}
	return result
}

func main() {
	fmt.Printf("%v (expected: [true false])\n",
		canMakePaliQueries("abcda", [][]int{{3, 3, 0}, {1, 2, 0}}))

	fmt.Printf("%v (expected: [true])\n",
		canMakePaliQueries("abcda", [][]int{{0, 3, 1}}))

	fmt.Printf("%v (expected: [true true false])\n",
		canMakePaliQueries("abcdd", [][]int{{0, 4, 1}, {0, 2, 1}, {1, 2, 0}}))
}
```
