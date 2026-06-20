# 2137 — Pour Water Between Buckets To Make Water Levels Equal

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func equalizeWater(buckets []int, loss int) float64`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n log precision)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2137: Pour Water Between Buckets to Make Water Levels Equal
// https://leetcode.com/problems/pour-water-between-buckets-to-make-water-levels-equal/
// Difficulty: Medium [Paid]
// Time: O(n log precision) | Space: O(1)

import "fmt"

func equalizeWater(buckets []int, loss int) float64 {
	canReach := func(target float64) bool {
		need := 0.0
		give := 0.0
		for _, b := range buckets {
			amount := float64(b)
			if amount < target {
				need += target - amount
			} else {
				give += (amount - target) * (1.0 - float64(loss)/100.0)
			}
		}
		return give >= need
	}

	low, high := 0.0, 0.0
	for _, b := range buckets {
		if float64(b) > high {
			high = float64(b)
		}
	}

	for i := 0; i < 100; i++ {
		mid := (low + high) / 2
		if canReach(mid) {
			low = mid
		} else {
			high = mid
		}
	}

	return low
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", equalizeWater([]int{1, 2, 7}, 20))
	// Expected: ~2.000

	// Test case 2
	fmt.Println("Test 2:", equalizeWater([]int{2, 4, 6}, 50))
	// Expected: ~3.500

	// Test case 3
	fmt.Println("Test 3:", equalizeWater([]int{3, 3, 3}, 0))
	// Expected: 3.000
}
```
