# 0966 — Vowel Spellchecker

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func spellchecker(wordlist []string, queries []string) []string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n * L)  |  **Ruang:** O(n * L)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #966: Vowel Spellchecker
// https://leetcode.com/problems/vowel-spellchecker/
// Difficulty: Medium

import (
	"fmt"
	"strings"
)

// Time: O(n * L) | Space: O(n * L)
func spellchecker(wordlist []string, queries []string) []string {
  // HashMap: O(1) lookup
	exact := make(map[string]bool)
  // HashMap: O(1) lookup
	lower := make(map[string]string)
  // HashMap: O(1) lookup
	vowel := make(map[string]string)

	for _, w := range wordlist {
		exact[w] = true
		lo := strings.ToLower(w)
		if _, ok := lower[lo]; !ok {
			lower[lo] = w
		}
		vw := devowel(lo)
		if _, ok := vowel[vw]; !ok {
			vowel[vw] = w
		}
	}

	ans := make([]string, len(queries))
	for i, q := range queries {
		if exact[q] {
			ans[i] = q
		} else if w, ok := lower[strings.ToLower(q)]; ok {
			ans[i] = w
		} else if w, ok := vowel[devowel(strings.ToLower(q))]; ok {
			ans[i] = w
		} else {
			ans[i] = ""
		}
	}
	return ans
}

func devowel(s string) string {
	var sb strings.Builder
	for _, ch := range s {
		if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
			sb.WriteByte('*')
		} else {
			sb.WriteRune(ch)
		}
	}
	return sb.String()
}

func main() {
	fmt.Println(spellchecker([]string{"KiTe", "kite", "hare", "Hare"}, []string{"kite", "Kite", "KiTe", "Hare", "HARE", "Hear", "hear", "keti", "keet", "keto"}))
	fmt.Println(spellchecker([]string{"yellow", "wood"}, []string{"Yello", "wood", "yellow"}))
}
```
