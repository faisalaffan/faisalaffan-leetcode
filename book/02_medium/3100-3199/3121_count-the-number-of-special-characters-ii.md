# 3121 — Count The Number Of Special Characters Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func numberOfSpecialChars(word string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(26) = O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3121: Count the Number of Special Characters II
// https://leetcode.com/problems/count-the-number-of-special-characters-ii/
// Difficulty: Medium
// Time: O(n) | Space: O(26) = O(1)

import "fmt"

func numberOfSpecialChars(word string) int {
  // Alokasi slice
	firstLower := make([]int, 26)
  // Alokasi slice
	lastUpper := make([]int, 26)
  // Range loop
	for i := range firstLower {
		firstLower[i] = -1
		lastUpper[i] = -1
	}

	for i, ch := range word {
		if ch >= 'a' && ch <= 'z' {
			idx := ch - 'a'
			if firstLower[idx] == -1 {
				firstLower[idx] = i
			}
		} else {
			idx := ch - 'A'
			lastUpper[idx] = i
		}
	}

	ans := 0
	for i := 0; i < 26; i++ {
		if firstLower[i] != -1 && lastUpper[i] != -1 && firstLower[i] > lastUpper[i] {
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println(numberOfSpecialChars("aaAbcBC"))  // Expected: 3
	fmt.Println(numberOfSpecialChars("abc"))       // Expected: 0
	fmt.Println(numberOfSpecialChars("AbBCab"))    // Expected: 0
}
```
