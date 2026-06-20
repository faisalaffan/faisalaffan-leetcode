# 1763 — Longest Nice Substring

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func LongestNiceSubstring(s string) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n^2), Space: O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #1763: Longest Nice Substring
// https://leetcode.com/problems/longest-nice-substring/
// Difficulty: Easy

import "fmt"
import "unicode"

// Time: O(n^2), Space: O(n)
func LongestNiceSubstring(s string) string {
	result := ""
  // Linear scan O(n)
	for i := 0; i < len(s); i++ {
		lower := 0
		upper := 0
		for j := i; j < len(s); j++ {
			ch := rune(s[j])
			if unicode.IsUpper(ch) {
				upper |= 1 << (unicode.ToLower(ch) - 'a')
			} else {
				lower |= 1 << (ch - 'a')
			}
			if lower == upper && j-i+1 > len(result) {
				result = s[i : j+1]
			}
		}
	}
	return result
}

func main() {
	fmt.Println(LongestNiceSubstring("YazaAay"))
	fmt.Println(LongestNiceSubstring("Bb"))
	fmt.Println(LongestNiceSubstring("c"))
}
```
