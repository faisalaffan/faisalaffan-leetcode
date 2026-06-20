# 0475 — Heaters

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func Heaters(houses []int, heaters []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Binary Search

**Kompleksitas Waktu:** O(n log n + m log n) where n = len(heaters), m = len(houses)  
**Kompleksitas Ruang:** O(log n) for sorting

> **Untuk fresh graduate:** Kuasai dulu teknik **Binary Search** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #475: Heaters
// https://leetcode.com/problems/heaters/
// Difficulty: Medium
// Time: O(n log n + m log n) where n = len(heaters), m = len(houses)
// Space: O(log n) for sorting

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(Heaters([]int{1, 2, 3}, []int{2}))
	fmt.Println(Heaters([]int{1, 2, 3, 4}, []int{1, 4}))
	fmt.Println(Heaters([]int{1, 5}, []int{2}))
}

func Heaters(houses []int, heaters []int) int {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(heaters)
	maxRadius := 0

	for _, house := range houses {
		// Binary search to find nearest heater
		idx := sort.SearchInts(heaters, house)
		minDist := int(^uint(0) >> 1) // MaxInt

		if idx < len(heaters) {
			dist := heaters[idx] - house
			if dist < 0 {
				dist = -dist
			}
			if dist < minDist {
				minDist = dist
			}
		}
		if idx > 0 {
			dist := house - heaters[idx-1]
			if dist < 0 {
				dist = -dist
			}
			if dist < minDist {
				minDist = dist
			}
		}

		if minDist > maxRadius {
			maxRadius = minDist
		}
	}

	return maxRadius
}
```
