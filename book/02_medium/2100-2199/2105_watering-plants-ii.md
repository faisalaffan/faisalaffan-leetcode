# 2105 — Watering Plants Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func minimumRefill(plants []int, capacityA int, capacityB int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2105: Watering Plants II
// https://leetcode.com/problems/watering-plants-ii/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func minimumRefill(plants []int, capacityA int, capacityB int) int {
	n := len(plants)
	alice := 0
	bob := n - 1
	waterA := capacityA
	waterB := capacityB
	refills := 0

	for alice < bob {
		// Alice waters
		if waterA < plants[alice] {
			refills++
			waterA = capacityA
		}
		waterA -= plants[alice]
		alice++

		// Bob waters
		if waterB < plants[bob] {
			refills++
			waterB = capacityB
		}
		waterB -= plants[bob]
		bob--
	}

	// Same plant?
	if alice == bob {
		if waterA >= waterB {
			if waterA < plants[alice] {
				refills++
			}
		} else {
			if waterB < plants[bob] {
				refills++
			}
		}
	}

	return refills
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minimumRefill([]int{2, 2, 3, 3}, 5, 5))
	// Expected: 1

	// Test case 2
	fmt.Println("Test 2:", minimumRefill([]int{2, 2, 3, 3}, 3, 4))
	// Expected: 2

	// Test case 3
	fmt.Println("Test 3:", minimumRefill([]int{5}, 10, 8))
	// Expected: 0
}
```
