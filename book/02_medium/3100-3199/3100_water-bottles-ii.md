# 3100 — Water Bottles Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxBottlesDrunk(numBottles int, numExchange int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(log n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3100: Water Bottles II
// https://leetcode.com/problems/water-bottles-ii/
// Difficulty: Medium
// Time: O(log n) | Space: O(1)

import "fmt"

func maxBottlesDrunk(numBottles int, numExchange int) int {
	total := numBottles
	empty := numBottles

	for empty >= numExchange {
		empty -= numExchange
		numExchange++
		total++
		empty++
	}

	return total
}

func main() {
	fmt.Println(maxBottlesDrunk(13, 6)) // Expected: 15
	fmt.Println(maxBottlesDrunk(10, 3)) // Expected: 13
}
```
