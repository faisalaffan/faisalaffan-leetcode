# 1897 — Redistribute Characters To Make All Strings Equal

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func MakeEqual(words []string) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n * len), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1897: Redistribute Characters to Make All Strings Equal
// https://leetcode.com/problems/redistribute-characters-to-make-all-strings-equal/
// Difficulty: Easy

import "fmt"

// Time: O(n * len), Space: O(1)
func MakeEqual(words []string) bool {
  // Alokasi slice
	freq := make([]int, 26)
	for _, w := range words {
  // Linear scan O(n)
		for i := 0; i < len(w); i++ {
			freq[w[i]-'a']++
		}
	}
	n := len(words)
	for _, count := range freq {
		if count%n != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(MakeEqual([]string{"abc", "aabc", "bc"}))
	fmt.Println(MakeEqual([]string{"ab", "a"}))
}
```
