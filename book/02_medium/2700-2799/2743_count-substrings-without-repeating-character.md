# 2743 — Count Substrings Without Repeating Character

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func CountSubstringsWithoutRepeatingCharacter(s string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Two Pointer

**Waktu:** O(n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2743: Count Substrings Without Repeating Character
// https://leetcode.com/problems/count-substrings-without-repeating-character/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func CountSubstringsWithoutRepeatingCharacter(s string) int {
  // HashMap: O(1) lookup
	lastPos := make(map[byte]int)
	left := 0
	count := 0

	for right := 0; right < len(s); right++ {
		if pos, ok := lastPos[s[right]]; ok && pos >= left {
			left = pos + 1
		}
		lastPos[s[right]] = right
		count += right - left + 1
	}

	return count
}

func main() {
	fmt.Println(CountSubstringsWithoutRepeatingCharacter("abcabc"))
	fmt.Println(CountSubstringsWithoutRepeatingCharacter("aaaa"))
	fmt.Println(CountSubstringsWithoutRepeatingCharacter(""))
}
```
