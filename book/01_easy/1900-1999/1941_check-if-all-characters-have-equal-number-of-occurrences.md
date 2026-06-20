# 1941 — Check If All Characters Have Equal Number Of Occurrences

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func CheckIfAllCharactersHaveEqualNumberOfOccurrences(s string) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n), Space: O(1) (max 26 chars)  |  **Ruang:** O(1) (max 26 chars)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1941: Check if All Characters Have Equal Number of Occurrences
// https://leetcode.com/problems/check-if-all-characters-have-equal-number-of-occurrences/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CheckIfAllCharactersHaveEqualNumberOfOccurrences("abacbc")) // true
	fmt.Println(CheckIfAllCharactersHaveEqualNumberOfOccurrences("aaabb"))  // false
}

// Time: O(n), Space: O(1) (max 26 chars)
func CheckIfAllCharactersHaveEqualNumberOfOccurrences(s string) bool {
  // HashMap: O(1) lookup
	freq := make(map[byte]int)
  // Linear scan O(n)
	for i := 0; i < len(s); i++ {
		freq[s[i]]++
	}

	var target int
	for _, v := range freq {
		target = v
		break
	}
	for _, v := range freq {
		if v != target {
			return false
		}
	}
	return true
}
```
