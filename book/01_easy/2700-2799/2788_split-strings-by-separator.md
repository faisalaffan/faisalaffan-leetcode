# 2788 — Split Strings By Separator

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func SplitStringsBySeparator(words []string, separator byte) []string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #2788: Split Strings by Separator
// https://leetcode.com/problems/split-strings-by-separator/
// Difficulty: Easy
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(SplitStringsBySeparator([]string{"one.two.three", "four.five", "six"}, '.'))
	fmt.Println(SplitStringsBySeparator([]string{"$easy$", "$problem$"}, '$'))
}

func SplitStringsBySeparator(words []string, separator byte) []string {
	result := []string{}
	for _, w := range words {
		parts := strings.Split(w, string(separator))
		for _, p := range parts {
			if p != "" {
				result = append(result, p)
			}
		}
	}
	return result
}
```
