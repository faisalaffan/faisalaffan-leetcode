# 3744 — Find Kth Character In Expanded String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func findKthCharacterInExpandedString(s string, k int) byte`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #3744: Find Kth Character in Expanded String
// https://leetcode.com/problems/find-kth-character-in-expanded-string/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"strings"
)

func findKthCharacterInExpandedString(s string, k int) byte {
	words := strings.Fields(s)

	for _, word := range words {
		l := len(word)
		m := l * (l + 1) / 2

		if k == m {
			return ' '
		} else if k > m {
			k -= (m + 1)
			continue
		} else {
			cur := 0
			for i, ch := range word {
				cur += (i + 1)
				if k < cur {
					return byte(ch)
				}
			}
			return ' '
		}
	}
	return ' '
}

func main() {
	fmt.Printf("%c\n", findKthCharacterInExpandedString("hello world", 0))
	fmt.Printf("%c\n", findKthCharacterInExpandedString("hello world", 15))
	fmt.Printf("%c\n", findKthCharacterInExpandedString("hello world", 20))
}
```
