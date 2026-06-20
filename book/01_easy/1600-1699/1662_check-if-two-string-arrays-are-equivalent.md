# 1662 — Check If Two String Arrays Are Equivalent

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func ArrayStringsAreEqual(word1 []string, word2 []string) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(n) where n is total characters  |  **Ruang:** O(n) where n is total characters


## 💻 Solusi Go

```go
package main

// LeetCode #1662: Check If Two String Arrays Are Equivalent
// https://leetcode.com/problems/check-if-two-string-arrays-are-equivalent/
// Difficulty: Easy

import "fmt"
import "strings"

// Time: O(n), Space: O(n) where n is total characters
func ArrayStringsAreEqual(word1 []string, word2 []string) bool {
	return strings.Join(word1, "") == strings.Join(word2, "")
}

func main() {
	fmt.Println(ArrayStringsAreEqual([]string{"ab", "c"}, []string{"a", "bc"}))
	fmt.Println(ArrayStringsAreEqual([]string{"a", "cb"}, []string{"ab", "c"}))
	fmt.Println(ArrayStringsAreEqual([]string{"abc", "d", "defg"}, []string{"abcddefg"}))
}
```
