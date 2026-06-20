# 0320 — Generalized Abbreviation

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func generateAbbreviations(word string) []string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Backtracking

**Waktu:** O(n * 2^n)  |  **Ruang:** O(n * 2^n)

> 🎓 **Fresh Grad Tips:** Kuasai **Backtracking** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #320: Generalized Abbreviation
// https://leetcode.com/problems/generalized-abbreviation/
// Difficulty: Medium [Paid]
// Time: O(n * 2^n) | Space: O(n * 2^n)

import (
	"fmt"
	"strconv"
)

func generateAbbreviations(word string) []string {
	result := []string{}
	backtrack(word, 0, 0, "", &result)
	return result
}

func backtrack(word string, index int, count int, cur string, result *[]string) {
	if index == len(word) {
		if count > 0 {
			cur += strconv.Itoa(count)
		}
		*result = append(*result, cur)
		return
	}

	// Abbreviate current character (increase count)
	backtrack(word, index+1, count+1, cur, result)

	// Keep current character
	if count > 0 {
		cur += strconv.Itoa(count)
	}
	cur += string(word[index])
	backtrack(word, index+1, 0, cur, result)
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", generateAbbreviations("word"))
	// Expected: ["word", "1ord", "w1rd", "2rd", "wo1d", "1o1d", "w2d", "3d", "wor1", "1or1", "w1r1", "2r1", "wo2", "1o2", "w3", "4"]

	// Test case 2: Empty string
	fmt.Println("Test 2:", generateAbbreviations(""))
	// Expected: [""]

	// Test case 3: Single character
	fmt.Println("Test 3:", generateAbbreviations("a"))
	// Expected: ["a", "1"]
}
```
