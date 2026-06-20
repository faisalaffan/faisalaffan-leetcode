# 1641 — Count Sorted Vowel Strings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func CountVowelStrings(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(N), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1641: Count Sorted Vowel Strings
// https://leetcode.com/problems/count-sorted-vowel-strings/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(CountVowelStrings(1))
	fmt.Println(CountVowelStrings(2))
	fmt.Println(CountVowelStrings(33))
}

func CountVowelStrings(n int) int {
	// Time: O(N), Space: O(1)
	// Combinatorics: C(n+4, 4) = (n+4)*(n+3)*(n+2)*(n+1)/24
	return (n + 4) * (n + 3) * (n + 2) * (n + 1) / 24
}
```
