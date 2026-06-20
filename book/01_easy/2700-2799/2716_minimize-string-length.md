# 2716 — Minimize String Length

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimizeStringLength(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2716: Minimize String Length
// https://leetcode.com/problems/minimize-string-length/
// Difficulty: Easy
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(MinimizeStringLength("aaabc"))
	fmt.Println(MinimizeStringLength("cbbd"))
}

func MinimizeStringLength(s string) int {
	seen := map[rune]bool{}
	for _, c := range s {
		seen[c] = true
	}
	return len(seen)
}
```
