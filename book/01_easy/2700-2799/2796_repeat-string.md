# 2796 — Repeat String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func RepeatString(s string, n int) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2796: Repeat String
// https://leetcode.com/problems/repeat-string/
// Difficulty: Easy [Paid]
// Time: O(n) | Space: O(n)
// Note: JS problem, adapted to Go. Repeats string n times.

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(RepeatString("abc", 3))
	fmt.Println(RepeatString("x", 5))
}

func RepeatString(s string, n int) string {
	return strings.Repeat(s, n)
}
```
