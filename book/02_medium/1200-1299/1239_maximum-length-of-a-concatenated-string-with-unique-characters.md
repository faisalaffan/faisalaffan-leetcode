# 1239 — Maximum Length Of A Concatenated String With Unique Characters

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func maxLength(arr []string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Backtracking, Bitmask

**Kompleksitas Waktu:** O(2^n) worst case  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Backtracking** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #1239: Maximum Length of a Concatenated String with Unique Characters
// https://leetcode.com/problems/maximum-length-of-a-concatenated-string-with-unique-characters/
// Difficulty: Medium

// Backtracking with bitmask. Each string is represented as a bitmask.
// Only concatenate if no character conflict.

// Time: O(2^n) worst case
// Space: O(n)

func maxLength(arr []string) int {
  // Alokasi slice integer
	masks := make([]int, 0)

	for _, s := range arr {
		mask := 0
		valid := true
		for _, ch := range s {
			bit := 1 << (ch - 'a')
			if mask&bit != 0 {
				valid = false
				break
			}
			mask |= bit
		}
		if valid {
			masks = append(masks, mask)
		}
	}

	maxLen := 0
	var backtrack func(idx, mask, length int)
	backtrack = func(idx, mask, length int) {
		if length > maxLen {
			maxLen = length
		}
		for i := idx; i < len(masks); i++ {
			if mask&masks[i] == 0 {
				backtrack(i+1, mask|masks[i], length+countBits(masks[i]))
			}
		}
	}

	backtrack(0, 0, 0)
	return maxLen
}

func countBits(mask int) int {
	count := 0
	for mask > 0 {
		count += mask & 1
		mask >>= 1
	}
	return count
}

func main() {
	fmt.Printf("%d (expected: 4)\n", maxLength([]string{"un", "iq", "ue"}))
	fmt.Printf("%d (expected: 6)\n", maxLength([]string{"cha", "r", "act", "ers"}))
	fmt.Printf("%d (expected: 26)\n", maxLength([]string{"abcdefghijklmnopqrstuvwxyz"}))
}
```
