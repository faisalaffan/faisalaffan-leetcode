# 1217 — Minimum Cost To Move Chips To The Same Position

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minCostToMoveChips(position []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1217: Minimum Cost to Move Chips to The Same Position
// https://leetcode.com/problems/minimum-cost-to-move-chips-to-the-same-position/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(minCostToMoveChips([]int{1, 2, 3}))    // 1
	fmt.Println(minCostToMoveChips([]int{2, 2, 2, 3, 3})) // 2
}

// LeetCode submission: minCostToMoveChips
func minCostToMoveChips(position []int) int {
	even, odd := 0, 0
	for _, p := range position {
		if p%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	if even < odd {
		return even
	}
	return odd
}
```
