# 0833 — Find And Replace In String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func FindAndReplaceInString(s string, indices []int, sources []string, targets []string) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n + m) where n = len(s), m = len(indices)  |  **Ruang:** O(n + m)


## 💻 Solusi Go

```go
package main

// LeetCode #833: Find And Replace in String
// https://leetcode.com/problems/find-and-replace-in-string/
// Difficulty: Medium

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(FindAndReplaceInString("abcd", []int{0, 2}, []string{"a", "cd"}, []string{"eee", "ffff"}))
	fmt.Println(FindAndReplaceInString("abcd", []int{0, 2}, []string{"ab", "ec"}, []string{"eee", "ffff"}))
	fmt.Println(FindAndReplaceInString("jjievdtjfb", []int{4, 6}, []string{"md", "tjfb"}, []string{"foe", "oov"}))
}

// Time: O(n + m) where n = len(s), m = len(indices) | Space: O(n + m)
func FindAndReplaceInString(s string, indices []int, sources []string, targets []string) string {
	n := len(s)
  // Alokasi slice
	replace := make([]int, n)
  // Range loop
	for i := range replace {
		replace[i] = -1
	}

	for k, idx := range indices {
		if idx+len(sources[k]) <= n && s[idx:idx+len(sources[k])] == sources[k] {
			replace[idx] = k
		}
	}

	var sb strings.Builder
	for i := 0; i < n; {
		if replace[i] >= 0 {
			sb.WriteString(targets[replace[i]])
			i += len(sources[replace[i]])
		} else {
			sb.WriteByte(s[i])
			i++
		}
	}

	return sb.String()
}
```
