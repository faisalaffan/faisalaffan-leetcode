# 0853 — Car Fleet

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func CarFleet(target int, position []int, speed []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

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
  // Edge case: input kosong
	if n == 0 {
		return 0
	}

	type car struct {
		pos  int
		time float64
	}
	cars := make([]car, n)
  // Range loop
	for i := range position {
		cars[i] = car{position[i], float64(target-position[i]) / float64(speed[i])}
	}

  // Custom sort
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
