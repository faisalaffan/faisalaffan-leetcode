# 2255 — Count Prefixes Of A Given String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func CountPrefixesOfAGivenString(words []string, s string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Prefix Sum

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Prefix Sum** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2255: Count Prefixes of a Given String
// https://leetcode.com/problems/count-prefixes-of-a-given-string/
// Difficulty: Easy
// Time O(n * m) | Space O(1)

import "fmt"

func main() {
	fmt.Println(CountPrefixesOfAGivenString([]string{"a", "b", "c", "ab", "bc", "abc"}, "abc"))           // 3
	fmt.Println(CountPrefixesOfAGivenString([]string{"a", "a"}, "aa"))                                      // 2
	fmt.Println(CountPrefixesOfAGivenString([]string{"feh", "w", "w", "l", "w", "o", "w", "o", "w"}, "w")) // 0
}

func CountPrefixesOfAGivenString(words []string, s string) int {
	count := 0
	for _, w := range words {
		if len(w) <= len(s) && s[:len(w)] == w {
			count++
		}
	}
	return count
}
```
