# 2129 — Capitalize The Title

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func CapitalizeTheTitle(title string) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #2129: Capitalize the Title
// https://leetcode.com/problems/capitalize-the-title/
// Difficulty: Easy

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(CapitalizeTheTitle("capiTalIze tHe titLe")) // "Capitalize The Title"
	fmt.Println(CapitalizeTheTitle("First leTTER of EACH Word")) // "First Letter of Each Word"
	fmt.Println(CapitalizeTheTitle("i lOve leetcode"))           // "i Love Leetcode"
}

// Time: O(n), Space: O(n)
func CapitalizeTheTitle(title string) string {
	words := strings.Fields(title)
	for i, w := range words {
		lower := strings.ToLower(w)
		if len(lower) > 2 {
			runes := []rune(lower)
			runes[0] = unicode.ToUpper(runes[0])
			words[i] = string(runes)
		} else {
			words[i] = lower
		}
	}
	return strings.Join(words, " ")
}
```
