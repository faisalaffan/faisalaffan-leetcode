# 0434 — Number Of Segments In A String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func NumberOfSegmentsInAString(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #434: Number of Segments in a String
// https://leetcode.com/problems/number-of-segments-in-a-string/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func NumberOfSegmentsInAString(s string) int {
	count := 0
	inSegment := false
	for _, c := range s {
		if c != ' ' && !inSegment {
			count++
			inSegment = true
		} else if c == ' ' {
			inSegment = false
		}
	}
	return count
}

func main() {
	fmt.Println(NumberOfSegmentsInAString("Hello, my name is John"))
	fmt.Println(NumberOfSegmentsInAString("Hello"))
	fmt.Println(NumberOfSegmentsInAString(""))
}
```
