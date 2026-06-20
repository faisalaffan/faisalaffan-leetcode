# 1668 — Maximum Repeating Substring

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func MaxRepeating(sequence string, word string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n * m), Space: O(n) where n = len(sequence), m = len(word)  |  **Ruang:** O(n) where n = len(sequence), m = len(word)


## 💻 Solusi Go

```go
package main

// LeetCode #1668: Maximum Repeating Substring
// https://leetcode.com/problems/maximum-repeating-substring/
// Difficulty: Easy

import (
	"fmt"
	"strings"
)

// Time: O(n * m), Space: O(n) where n = len(sequence), m = len(word)
func MaxRepeating(sequence string, word string) int {
	k := 0
	repeated := word
	for strings.Contains(sequence, repeated) {
		k++
		repeated += word
	}
	return k
}

func main() {
	fmt.Println(MaxRepeating("ababc", "ab"))
	fmt.Println(MaxRepeating("ababc", "ba"))
	fmt.Println(MaxRepeating("ababc", "ac"))
}
```
