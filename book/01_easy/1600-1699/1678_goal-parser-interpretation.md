# 1678 — Goal Parser Interpretation

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func Interpret(command string) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #1678: Goal Parser Interpretation
// https://leetcode.com/problems/goal-parser-interpretation/
// Difficulty: Easy

import (
	"fmt"
	"strings"
)

// Time: O(n), Space: O(n)
func Interpret(command string) string {
	command = strings.ReplaceAll(command, "()", "o")
	command = strings.ReplaceAll(command, "(al)", "al")
	return command
}

func main() {
	fmt.Println(Interpret("G()(al)"))
	fmt.Println(Interpret("G()()()()(al)"))
	fmt.Println(Interpret("(al)G(al)()()G"))
}
```
