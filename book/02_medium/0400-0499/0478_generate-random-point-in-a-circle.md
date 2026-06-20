# 0478 — Generate Random Point In A Circle

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func Constructor(radiusAndCenter []float64) Solution
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1) per call  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #478: Generate Random Point in a Circle
// https://leetcode.com/problems/generate-random-point-in-a-circle/
// Difficulty: Medium
// Time: O(1) per call
// Space: O(1)

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	sol := Constructor([]float64{1, 0, 0})
	// Generate a few random points
	for i := 0; i < 3; i++ {
		p := sol.RandPoint()
		fmt.Printf("%.4f %.4f\n", p[0], p[1])
	}
}

type Solution struct {
	radius, xCenter, yCenter float64
}

func Constructor(radiusAndCenter []float64) Solution {
	return Solution{
		radius:  radiusAndCenter[0],
		xCenter: radiusAndCenter[1],
		yCenter: radiusAndCenter[2],
	}
}

func (s *Solution) RandPoint() []float64 {
	// Use random angle and random radius (sqrt for uniform distribution)
	angle := rand.Float64() * 2 * math.Pi
	r := s.radius * math.Sqrt(rand.Float64())
	x := s.xCenter + r*math.Cos(angle)
	y := s.yCenter + r*math.Sin(angle)
	return []float64{x, y}
}
```
