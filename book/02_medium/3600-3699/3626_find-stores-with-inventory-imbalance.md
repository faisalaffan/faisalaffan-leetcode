# 3626 — Find Stores With Inventory Imbalance

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func FindStoresWithInventoryImbalance(inventory []int, threshold int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #3626: Find Stores with Inventory Imbalance
// https://leetcode.com/problems/find-stores-with-inventory-imbalance/
// Difficulty: Medium
// Complexity: O(n) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	inventory := []int{10, 5, 15, 3, 20}
	threshold := 10
	fmt.Println("Test 1:", FindStoresWithInventoryImbalance(inventory, threshold))
	// Test case 2
	inventory2 := []int{100, 50, 75}
	threshold2 := 20
	fmt.Println("Test 2:", FindStoresWithInventoryImbalance(inventory2, threshold2))
	// Test case 3
	inventory3 := []int{1, 1, 1}
	threshold3 := 0
	fmt.Println("Test 3:", FindStoresWithInventoryImbalance(inventory3, threshold3))
}

func FindStoresWithInventoryImbalance(inventory []int, threshold int) int {
	// Count stores where difference from average exceeds threshold
	if len(inventory) == 0 {
		return 0
	}
	sum := 0
	for _, v := range inventory {
		sum += v
	}
	avg := float64(sum) / float64(len(inventory))
	count := 0
	for _, v := range inventory {
		diff := float64(v) - avg
		if diff < 0 {
			diff = -diff
		}
		if diff > float64(threshold) {
			count++
		}
	}
	return count
}
```
