# 0696 — Count Binary Substrings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func countBinarySubstrings(s string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n). Space: O(1).  |  **Ruang:** O(1).


## 💻 Solusi Go

```go
package main

// LeetCode #696: Count Binary Substrings
// https://leetcode.com/problems/count-binary-substrings/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(countBinarySubstrings("00110011")) // 6
	fmt.Println(countBinarySubstrings("10101"))    // 4
	fmt.Println(countBinarySubstrings("00110"))    // 3
}

// countBinarySubstrings counts substrings that have equal numbers of 0s and 1s.
// Time: O(n). Space: O(1).
func countBinarySubstrings(s string) int {
	prev, curr, result := 0, 1, 0
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			curr++
		} else {
			prev = curr
			curr = 1
		}
		if prev >= curr {
			result++
		}
	}
	return result
}
```
