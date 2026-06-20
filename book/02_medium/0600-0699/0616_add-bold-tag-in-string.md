# 0616 — Add Bold Tag In String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func AddBoldTag(s string, words []string) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n * L) where n = len(s), L = total length of all words  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #616: Add Bold Tag in String
// https://leetcode.com/problems/add-bold-tag-in-string/
// Difficulty: Medium [Paid]
// Time: O(n * L) where n = len(s), L = total length of all words
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(AddBoldTag("abcxyz123", []string{"abc", "123"}))
	fmt.Println(AddBoldTag("aaabbcc", []string{"aaa", "aab", "bc"}))
}

func AddBoldTag(s string, words []string) string {
	n := len(s)
	bold := make([]bool, n)

	for _, word := range words {
		for i := 0; i <= n-len(word); i++ {
			if s[i:i+len(word)] == word {
				for j := i; j < i+len(word); j++ {
					bold[j] = true
				}
			}
		}
	}

	result := ""
	i := 0
	for i < n {
		if bold[i] {
			result += "<b>"
			for i < n && bold[i] {
				result += string(s[i])
				i++
			}
			result += "</b>"
		} else {
			result += string(s[i])
			i++
		}
	}

	return result
}
```
