# 0926 — Flip String To Monotone Increasing

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func minFlipsMonoIncr(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #926: Flip String to Monotone Increasing
// https://leetcode.com/problems/flip-string-to-monotone-increasing/
// Difficulty: Medium

import "fmt"

// Time: O(n) | Space: O(1)
func minFlipsMonoIncr(s string) int {
	ones, flips := 0, 0
	for _, ch := range s {
		if ch == '1' {
			ones++
		} else {
			flips = min(flips+1, ones)
		}
	}
	return flips
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minFlipsMonoIncr("00110"))
	fmt.Println(minFlipsMonoIncr("010110"))
	fmt.Println(minFlipsMonoIncr("00011000"))
}
```
