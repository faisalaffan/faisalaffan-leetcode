# 2594 — Minimum Time To Repair Cars

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func repairCars(ranks []int, cars int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Binary Search

**Kompleksitas Waktu:** O(n log minTime)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Binary Search** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2594: Minimum Time to Repair Cars
// https://leetcode.com/problems/minimum-time-to-repair-cars/
// Difficulty: Medium
// Time: O(n log minTime) | Space: O(1)

import (
	"fmt"
	"math"
)

func repairCars(ranks []int, cars int) int64 {
	canRepair := func(t int64) bool {
		var total int64
		for _, r := range ranks {
			total += int64(math.Sqrt(float64(t / int64(r))))
			if total >= int64(cars) {
				return true
			}
		}
		return false
	}

	left := int64(1)
	right := int64(ranks[0]) * int64(cars) * int64(cars)
  // Range loop: iterasi dengan indeks + nilai
	for i := range ranks {
		candidate := int64(ranks[i]) * int64(cars) * int64(cars)
		if candidate < right {
			right = candidate
		}
	}

  // Two-pointer: gerakkan kiri atau kanan
	for left < right {
		mid := left + (right-left)/2
		if canRepair(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", repairCars([]int{4, 2, 3, 1}, 10))
	// Expected: 16

	// Test case 2
	fmt.Println("Test 2:", repairCars([]int{5, 1, 8}, 6))
	// Expected: 16

	// Test case 3
	fmt.Println("Test 3:", repairCars([]int{1, 1, 1}, 5))
	// Expected: 3
}
```
