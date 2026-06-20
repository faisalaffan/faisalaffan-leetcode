# 3498 — Reverse Degree Of A String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func ReverseDegreeOfAString(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n). Space: O(1).  
**Kompleksitas Ruang:** O(1).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3498: Reverse Degree of a String
// https://leetcode.com/problems/reverse-degree-of-a-string/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(ReverseDegreeOfAString("abc"))
	fmt.Println(ReverseDegreeOfAString("zaba"))
}

// ReverseDegreeOfAString computes the sum of (position_in_reversed_alphabet * (i+1)) for each character.
// Reverse: a=26, b=25, ..., z=1.
// Time: O(n). Space: O(1).
func ReverseDegreeOfAString(s string) int {
	sum := 0
	for i, ch := range s {
		revPos := 26 - int(ch-'a')
		sum += revPos * (i + 1)
	}
	return sum
}
```
