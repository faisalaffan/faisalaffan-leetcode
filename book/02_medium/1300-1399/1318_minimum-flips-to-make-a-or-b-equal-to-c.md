# 1318 — Minimum Flips To Make A Or B Equal To C

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minFlips(a int, b int, c int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(bit length) = O(1) since 32 bits  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1318: Minimum Flips to Make a OR b Equal to c
// https://leetcode.com/problems/minimum-flips-to-make-a-or-b-equal-to-c/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(minFlips(2, 6, 5)) // 3

	// Test case 2
	fmt.Println(minFlips(4, 2, 7)) // 1

	// Test case 3
	fmt.Println(minFlips(1, 2, 3)) // 0
}

// Time: O(bit length) = O(1) since 32 bits
// Space: O(1)
func minFlips(a int, b int, c int) int {
	flips := 0
	for i := 0; i < 32; i++ {
		bitC := (c >> i) & 1
		bitA := (a >> i) & 1
		bitB := (b >> i) & 1

		if bitC == 1 {
			if bitA == 0 && bitB == 0 {
				flips++ // need to flip one of them to 1
			}
		} else {
			if bitA == 1 {
				flips++
			}
			if bitB == 1 {
				flips++
			}
		}
	}
	return flips
}
```
