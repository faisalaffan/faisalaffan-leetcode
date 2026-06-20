# 3601 — Find Drivers With Improved Fuel Efficiency

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func FindDriversWithImprovedFuelEfficiency(before, after []float64) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #3601: Find Drivers with Improved Fuel Efficiency
// https://leetcode.com/problems/find-drivers-with-improved-fuel-efficiency/
// Difficulty: Medium
// Complexity: O(n) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	before := []float64{10, 15, 20}
	after := []float64{12, 14, 22}
	fmt.Println("Test 1:", FindDriversWithImprovedFuelEfficiency(before, after))
	// Test case 2
	before2 := []float64{10, 10}
	after2 := []float64{9, 11}
	fmt.Println("Test 2:", FindDriversWithImprovedFuelEfficiency(before2, after2))
	// Test case 3
	before3 := []float64{5}
	after3 := []float64{6}
	fmt.Println("Test 3:", FindDriversWithImprovedFuelEfficiency(before3, after3))
}

func FindDriversWithImprovedFuelEfficiency(before, after []float64) int {
	count := 0
  // Linear scan O(n)
	for i := 0; i < len(before) && i < len(after); i++ {
		if after[i] > before[i] {
			count++
		}
	}
	return count
}
```
