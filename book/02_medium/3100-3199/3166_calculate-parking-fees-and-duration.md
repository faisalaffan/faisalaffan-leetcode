# 3166 — Calculate Parking Fees And Duration

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func calculateParkingFees(records [][]int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3166: Calculate Parking Fees and Duration
// https://leetcode.com/problems/calculate-parking-fees-and-duration/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func calculateParkingFees(records [][]int) []int {
  // Membuat map (HashMap) — pencarian O(1)
	fees := make(map[int]int)
  // Membuat map (HashMap) — pencarian O(1)
	durations := make(map[int]int)

	for _, r := range records {
		carID, entry, exit := r[0], r[1], r[2]
		dur := exit - entry
		durations[carID] += dur
		var fee int
		if dur <= 60 {
			fee = 10
		} else {
			fee = 10 + ((dur-60)+29)/30*5
		}
		fees[carID] += fee
	}

	var cars []int
	for id := range fees {
		cars = append(cars, id)
	}
  // Urutkan secara ascending — O(n log n)
	sort.Ints(cars)

  // Alokasi slice integer
	ans := make([]int, len(cars))
	for i, id := range cars {
		ans[i] = fees[id]
	}
	return ans
}

func main() {
	fmt.Println(calculateParkingFees([][]int{{1, 0, 30}, {1, 60, 120}, {2, 0, 90}})) // Expected: [20 15]
	fmt.Println(calculateParkingFees([][]int{{1, 0, 30}}))                             // Expected: [10]
}
```
