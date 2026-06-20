# 2738 — Count Occurrences In Text

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func CountOccurrencesInText(text string, word string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2738: Count Occurrences in Text
// https://leetcode.com/problems/count-occurrences-in-text/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import (
	"fmt"
	"strings"
)

func CountOccurrencesInText(text string, word string) int {
	count := 0
	words := strings.Fields(text)
	for _, w := range words {
		if w == word {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(CountOccurrencesInText("hello world hello", "hello"))
	fmt.Println(CountOccurrencesInText("this is a test test this", "test"))
	fmt.Println(CountOccurrencesInText("unique", "none"))
}
```
