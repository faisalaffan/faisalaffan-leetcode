# 0058 — Length Of Last Word

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func LengthOfLastWord(s string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #58: Length of Last Word
// https://leetcode.com/problems/length-of-last-word/
// Difficulty: Easy

import "fmt"

// Time: O(n) | Space: O(1)
func LengthOfLastWord(s string) int {
	length := 0
	i := len(s) - 1
	for i >= 0 && s[i] == ' ' {
		i--
	}
	for i >= 0 && s[i] != ' ' {
		length++
		i--
	}
	return length
}

func main() {
	fmt.Println(LengthOfLastWord("Hello World"))
	fmt.Println(LengthOfLastWord("   fly me   to   the moon  "))
	fmt.Println(LengthOfLastWord("luffy is still joyboy"))
}
```
