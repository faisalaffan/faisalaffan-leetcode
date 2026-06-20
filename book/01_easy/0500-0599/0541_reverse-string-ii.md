# 0541 — Reverse String Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func ReverseStringIi(s string, k int) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #541: Reverse String II
// https://leetcode.com/problems/reverse-string-ii/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(n)
func ReverseStringIi(s string, k int) string {
	b := []byte(s)
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(b); i += 2 * k {
		lo, hi := i, i+k-1
		if hi >= len(b) {
			hi = len(b) - 1
		}
		for lo < hi {
			b[lo], b[hi] = b[hi], b[lo]
			lo++
			hi--
		}
	}
	return string(b)
}

func main() {
	fmt.Println(ReverseStringIi("abcdefg", 2))
	fmt.Println(ReverseStringIi("abcd", 2))
}
```
