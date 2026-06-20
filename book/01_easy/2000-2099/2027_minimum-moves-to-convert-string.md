# 2027 — Minimum Moves To Convert String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumMovesToConvertString(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2027: Minimum Moves to Convert String
// https://leetcode.com/problems/minimum-moves-to-convert-string/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MinimumMovesToConvertString("XXX"))       // 1
	fmt.Println(MinimumMovesToConvertString("XXOX"))      // 2
	fmt.Println(MinimumMovesToConvertString("OOOO"))      // 0
}

// Time: O(n), Space: O(1)
func MinimumMovesToConvertString(s string) int {
	moves := 0
	i := 0
	for i < len(s) {
		if s[i] == 'X' {
			moves++
			i += 3
		} else {
			i++
		}
	}
	return moves
}
```
