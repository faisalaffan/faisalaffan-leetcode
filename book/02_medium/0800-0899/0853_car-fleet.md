# 0853 — Car Fleet

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func CarFleet(target int, position []int, speed []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #853: Car Fleet
// https://leetcode.com/problems/car-fleet/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(CarFleet(12, []int{10, 8, 0, 5, 3}, []int{2, 4, 1, 1, 3}))
	fmt.Println(CarFleet(10, []int{3}, []int{3}))
	fmt.Println(CarFleet(100, []int{0, 2, 4}, []int{4, 2, 1}))
}

// Time: O(n log n) | Space: O(n)
func CarFleet(target int, position []int, speed []int) int {
	n := len(position)
  // Edge case: input kosong — langsung return
	if n == 0 {
		return 0
	}

	type car struct {
		pos  int
		time float64
	}
	cars := make([]car, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range position {
		cars[i] = car{position[i], float64(target-position[i]) / float64(speed[i])}
	}

  // Custom sort dengan comparator
	sort.Slice(cars, func(i, j int) bool {
		return cars[i].pos > cars[j].pos
	})

	fleets := 1
	maxTime := cars[0].time
	for i := 1; i < n; i++ {
		if cars[i].time > maxTime {
			fleets++
			maxTime = cars[i].time
		}
	}

	return fleets
}
```
