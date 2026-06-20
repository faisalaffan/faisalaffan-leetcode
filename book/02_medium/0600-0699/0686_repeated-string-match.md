# 0686 — Repeated String Match

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func repeatedStringMatch(a string, b string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n * m) worst case  
**Kompleksitas Ruang:** O(n + m)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #686: Repeated String Match
// https://leetcode.com/problems/repeated-string-match/
// Difficulty: Medium
// Time: O(n * m) worst case
// Space: O(n + m)

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(repeatedStringMatch("abcd", "cdabcdab"))
	fmt.Println(repeatedStringMatch("a", "aa"))
	fmt.Println(repeatedStringMatch("abc", "wxyz"))
}

func repeatedStringMatch(a string, b string) int {
	maxRepeats := len(b)/len(a) + 3
	var sb strings.Builder

	for i := 1; i <= maxRepeats; i++ {
		sb.WriteString(a)
		if strings.Contains(sb.String(), b) {
			return i
		}
	}

	return -1
}
```
