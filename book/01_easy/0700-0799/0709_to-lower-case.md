# 0709 — To Lower Case

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func toLowerCase(s string) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n). Space: O(n).  |  **Ruang:** O(n).


## 💻 Solusi Go

```go
package main

// LeetCode #709: To Lower Case
// https://leetcode.com/problems/to-lower-case/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(toLowerCase("Hello"))  // "hello"
	fmt.Println(toLowerCase("LOVELY")) // "lovely"
	fmt.Println(toLowerCase("here"))   // "here"
}

// toLowerCase converts a string to lowercase.
// Time: O(n). Space: O(n).
func toLowerCase(s string) string {
	b := []byte(s)
	for i, c := range b {
		if c >= 'A' && c <= 'Z' {
			b[i] = c + 32
		}
	}
	return string(b)
}
```
