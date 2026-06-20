# 0758 — Bold Words In String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func boldWords(words []string, s string) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n * L)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #758: Bold Words in String
// https://leetcode.com/problems/bold-words-in-string/
// Difficulty: Medium [Paid]
// Time: O(n * L)
// Space: O(n)

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(boldWords([]string{"ab", "bc"}, "aabcd"))
	fmt.Println(boldWords([]string{"abc", "123"}, "abcxyz123"))
}

func boldWords(words []string, s string) string {
	n := len(s)
	bold := make([]bool, n)

	for _, word := range words {
		start := 0
		for {
			idx := strings.Index(s[start:], word)
			if idx == -1 {
				break
			}
			pos := start + idx
			for i := pos; i < pos+len(word); i++ {
				bold[i] = true
			}
			start = pos + 1
		}
	}

	var result strings.Builder
	i := 0
	for i < n {
		if bold[i] {
			result.WriteString("<b>")
			for i < n && bold[i] {
				result.WriteByte(s[i])
				i++
			}
			result.WriteString("</b>")
		} else {
			result.WriteByte(s[i])
			i++
		}
	}

	return result.String()
}
```
