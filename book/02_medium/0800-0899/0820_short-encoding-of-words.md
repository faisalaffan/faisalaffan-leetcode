# 0820 — Short Encoding Of Words

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func ShortEncodingOfWords(words []string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n * L^2)  |  **Ruang:** O(n * L) where L is average word length

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #820: Short Encoding of Words
// https://leetcode.com/problems/short-encoding-of-words/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(ShortEncodingOfWords([]string{"time", "me", "bell"}))
	fmt.Println(ShortEncodingOfWords([]string{"t"}))
	fmt.Println(ShortEncodingOfWords([]string{"me", "time"}))
}

// Time: O(n * L^2) | Space: O(n * L) where L is average word length
func ShortEncodingOfWords(words []string) int {
  // HashMap: O(1) lookup
	set := make(map[string]bool)
	for _, word := range words {
		set[word] = true
	}

	for _, word := range words {
		for i := 1; i < len(word); i++ {
			delete(set, word[i:])
		}
	}

	ans := 0
	for word := range set {
		ans += len(word) + 1
	}
	return ans
}
```
