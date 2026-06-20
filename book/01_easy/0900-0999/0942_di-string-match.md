# 0942 — Di String Match

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func diStringMatch(s string) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n). Space: O(n).  
**Kompleksitas Ruang:** O(n).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #942: DI String Match
// https://leetcode.com/problems/di-string-match/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(diStringMatch("IDID")) // [0,4,1,3,2]
	fmt.Println(diStringMatch("III"))  // [0,1,2,3]
	fmt.Println(diStringMatch("DDI"))  // [3,2,0,1]
}

// diStringMatch returns a permutation that matches the DI pattern.
// Time: O(n). Space: O(n).
func diStringMatch(s string) []int {
	n := len(s)
  // Alokasi slice integer
	result := make([]int, n+1)
	low, high := 0, n
	for i, c := range s {
		if c == 'I' {
			result[i] = low
			low++
		} else {
			result[i] = high
			high--
		}
	}
	result[n] = low // or high (they are equal)
	return result
}
```
