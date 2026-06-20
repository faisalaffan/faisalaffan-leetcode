# 0937 — Reorder Data In Log Files

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func reorderLogFiles(logs []string) []string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #937: Reorder Data in Log Files
// https://leetcode.com/problems/reorder-data-in-log-files/
// Difficulty: Medium

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

// Time: O(n log n) | Space: O(n)
func reorderLogFiles(logs []string) []string {
	sort.SliceStable(logs, func(i, j int) bool {
		ci := strings.SplitN(logs[i], " ", 2)
		cj := strings.SplitN(logs[j], " ", 2)
		idI, restI := ci[0], ci[1]
		idJ, restJ := cj[0], cj[1]

		iDig := unicode.IsDigit(rune(restI[0]))
		jDig := unicode.IsDigit(rune(restJ[0]))

		if iDig && jDig {
			return false
		}
		if !iDig && !jDig {
			if restI != restJ {
				return restI < restJ
			}
			return idI < idJ
		}
		return !iDig
	})
	return logs
}

func main() {
	fmt.Println(reorderLogFiles([]string{"dig1 8 1 5 1", "let1 art can", "dig2 3 6", "let2 own kit dig", "let3 art zero"}))
	fmt.Println(reorderLogFiles([]string{"a1 9 2 3 1", "g1 act car", "zo4 4 7", "ab1 off key dog", "a8 act zoo"}))
}
```
