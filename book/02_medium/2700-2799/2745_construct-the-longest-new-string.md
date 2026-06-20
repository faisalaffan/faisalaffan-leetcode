# 2745 — Construct The Longest New String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func ConstructTheLongestNewString(x int, y int, z int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2745: Construct the Longest New String
// https://leetcode.com/problems/construct-the-longest-new-string/
// Difficulty: Medium
// Time: O(1) | Space: O(1)

import "fmt"

func ConstructTheLongestNewString(x int, y int, z int) int {
	// AA can pair with BB, BB can pair with AA, AB can go anywhere
	// Max pairs of AA and BB: min(x, y) + extra if any left
	used := 0

	// Use pairs of AA and BB
	pairs := x
	if y < pairs {
		pairs = y
	}
	used += pairs * 2

	// If both AA and BB have remaining, can add one more
	if x > pairs {
		used++
	}
	if y > pairs {
		used++
	}

	// AB can be inserted anywhere, but consumes z
	used += z

	return used * 2
}

func main() {
	fmt.Println(ConstructTheLongestNewString(1, 1, 1))
	fmt.Println(ConstructTheLongestNewString(2, 0, 2))
	fmt.Println(ConstructTheLongestNewString(0, 0, 5))
}
```
