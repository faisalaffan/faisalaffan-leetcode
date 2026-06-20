# 2103 — Rings And Rods

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func RingsAndRods(rings string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2103: Rings and Rods
// https://leetcode.com/problems/rings-and-rods/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(RingsAndRods("B0B6G0R6R0R6G9"))    // 1
	fmt.Println(RingsAndRods("B0R0G0R9R0B0G0"))    // 1
	fmt.Println(RingsAndRods("G4"))                 // 0
}

// Time: O(n), Space: O(1)
func RingsAndRods(rings string) int {
  // Alokasi slice integer
	rods := make([]int, 10)
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(rings); i += 2 {
		color := rings[i]
		rod := rings[i+1] - '0'
		switch color {
		case 'R':
			rods[rod] |= 1
		case 'G':
			rods[rod] |= 2
		case 'B':
			rods[rod] |= 4
		}
	}

	count := 0
	for _, v := range rods {
		if v == 7 { // R|G|B = 1|2|4 = 7
			count++
		}
	}
	return count
}
```
