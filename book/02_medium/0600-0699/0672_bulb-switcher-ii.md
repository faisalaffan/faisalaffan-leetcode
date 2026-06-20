# 0672 — Bulb Switcher Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func flipLights(n int, m int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #672: Bulb Switcher II
// https://leetcode.com/problems/bulb-switcher-ii/
// Difficulty: Medium
// Time: O(1)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(flipLights(1, 1))
	fmt.Println(flipLights(2, 1))
	fmt.Println(flipLights(3, 1))
}

func flipLights(n int, m int) int {
	// Only first 3 bulbs matter (key insight)
	if m == 0 {
		return 1
	}
	if n == 1 {
		return 2
	}
	if n == 2 {
		if m == 1 {
			return 3
		}
		return 4
	}
	// n >= 3
	if m == 1 {
		return 4
	}
	if m == 2 {
		return 7
	}
	return 8
}
```
