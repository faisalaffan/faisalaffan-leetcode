# 0612 — Shortest Distance In A Plane

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func ShortestDistance(points [][]float64) float64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #612: Shortest Distance in a Plane
// https://leetcode.com/problems/shortest-distance-in-a-plane/
// Difficulty: Medium [Paid]
// Time: O(n^2)
// Space: O(1)

import (
	"fmt"
	"math"
)

func main() {
	points := [][]float64{
		{-1, -1},
		{0, 0},
		{1, 1},
		{2, 2},
	}
	fmt.Printf("%.4f\n", ShortestDistance(points))
}

func ShortestDistance(points [][]float64) float64 {
	if len(points) < 2 {
		return 0
	}

	minDist := math.MaxFloat64
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			dx := points[i][0] - points[j][0]
			dy := points[i][1] - points[j][1]
			dist := math.Sqrt(dx*dx + dy*dy)
			if dist < minDist {
				minDist = dist
			}
		}
	}

	return minDist
}
```
