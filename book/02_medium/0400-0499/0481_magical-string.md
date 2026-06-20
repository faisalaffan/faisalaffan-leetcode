# 0481 — Magical String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func MagicalString(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #481: Magical String
// https://leetcode.com/problems/magical-string/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(MagicalString(6))
	fmt.Println(MagicalString(1))
}

func MagicalString(n int) int {
	if n <= 0 {
		return 0
	}
	if n <= 3 {
		return 1
	}

  // Alokasi slice integer
	s := make([]int, n)
	s[0], s[1], s[2] = 1, 2, 2
	count := 1
	writeIdx := 3
	readIdx := 2

	for writeIdx < n {
		val := 3 - s[writeIdx-1] // toggle between 1 and 2
		countTimes := s[readIdx]
		for i := 0; i < countTimes && writeIdx < n; i++ {
			s[writeIdx] = val
			if val == 1 {
				count++
			}
			writeIdx++
		}
		readIdx++
	}

	return count
}
```
