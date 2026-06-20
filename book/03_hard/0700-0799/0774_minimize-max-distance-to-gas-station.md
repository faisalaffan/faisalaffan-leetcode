# 0774 — Minimize Max Distance To Gas Station

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func minmaxGasDist(stations []int, k int) float64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Binary Search

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Binary Search** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #774: Minimize Max Distance to Gas Station
// https://leetcode.com/problems/minimize-max-distance-to-gas-station/
// Difficulty: Hard [Paid]
//
// Given sorted station positions and K additional stations to add,
// minimize the maximum distance between adjacent stations.
//
// Approach: Binary Search
// Search on the answer (minimum possible max gap). For a candidate D,
// count how many stations we need: for gap G, need ceil(G/D)-1 stations.
// If total needed <= K, D is feasible (we can achieve <= D).

import "fmt"
import "math"

func main() {
	fmt.Println(minmaxGasDist([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 9)) // 0.5
	fmt.Println(minmaxGasDist([]int{23, 24, 36, 39, 46, 56, 57, 65, 84, 98}, 1)) // 14.0
	fmt.Println(minmaxGasDist([]int{10, 19, 25, 27, 56, 63, 70, 87, 96, 97}, 3)) // 9.666667
}

func minmaxGasDist(stations []int, k int) float64 {
	// Find maximum gap
	var maxGap float64
	for i := 1; i < len(stations); i++ {
		gap := float64(stations[i] - stations[i-1])
		if gap > maxGap {
			maxGap = gap
		}
	}

	low, high := 0.0, maxGap

	// Binary search with 1e-6 precision
	for high-low > 1e-6 {
		mid := (low + high) / 2.0
		if feasible(stations, k, mid) {
			high = mid
		} else {
			low = mid
		}
	}

	return math.Round(high*1e6) / 1e6
}

func feasible(stations []int, k int, d float64) bool {
	count := 0
	for i := 1; i < len(stations); i++ {
		gap := float64(stations[i] - stations[i-1])
		// Number of additional stations needed to bring this gap <= d
		// ceil(gap / d) - 1
		needed := int(math.Ceil(gap/d)) - 1
		count += needed
		if count > k {
			return false
		}
	}
	return count <= k
}
```
