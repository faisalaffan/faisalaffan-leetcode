# 3330 — Find The Original Typed String I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func FindTheOriginalTypedStringI(word string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n). Space: O(1).  
**Kompleksitas Ruang:** O(1).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3330: Find the Original Typed String I
// https://leetcode.com/problems/find-the-original-typed-string-i/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindTheOriginalTypedStringI("aabbccdd"))
	fmt.Println(FindTheOriginalTypedStringI("aaaa"))
	fmt.Println(FindTheOriginalTypedStringI("abc"))
}

// FindTheOriginalTypedStringI counts possible original strings where adjacent equal characters could be merged.
// Time: O(n). Space: O(1).
func FindTheOriginalTypedStringI(word string) int {
	count := 1
	streak := 1
	for i := 1; i < len(word); i++ {
		if word[i] == word[i-1] {
			streak++
		} else {
			streak = 1
		}
		if streak >= 2 {
			// If we have at least 2 of the same char consecutively, we can type fewer
			count++
		}
	}
	return count
}
```
