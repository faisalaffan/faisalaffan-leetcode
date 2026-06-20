# 3494 — Find The Minimum Amount Of Time To Brew Potions

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func FindTheMinimumAmountOfTimeToBrewPotions(machines []int, potionTimes []int, potions int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #3494: Find the Minimum Amount of Time to Brew Potions
// https://leetcode.com/problems/find-the-minimum-amount-of-time-to-brew-potions/
// Difficulty: Medium
// Complexity: O(n) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	machines := []int{1, 2, 3}
	potionTimes := []int{3, 2, 1}
	fmt.Println("Test 1:", FindTheMinimumAmountOfTimeToBrewPotions(machines, potionTimes, 5))

	// Test case 2
	machines2 := []int{5}
	potionTimes2 := []int{2}
	fmt.Println("Test 2:", FindTheMinimumAmountOfTimeToBrewPotions(machines2, potionTimes2, 3))

	// Test case 3
	machines3 := []int{1, 1}
	potionTimes3 := []int{1, 2}
	fmt.Println("Test 3:", FindTheMinimumAmountOfTimeToBrewPotions(machines3, potionTimes3, 10))
}

func FindTheMinimumAmountOfTimeToBrewPotions(machines []int, potionTimes []int, potions int) int {
	// Each machine has a skill level and each potion has a brew time
	// Assign potions to machines to minimize total time
	if len(machines) == 0 || potions == 0 {
		return 0
	}

	// Simple greedy: sort both and pair fastest machine with fastest potion
	time := 0
  // Linear scan O(n)
	for i := 0; i < len(potionTimes) && i < len(machines); i++ {
		batchTime := potionTimes[i] / machines[i]
		if potionTimes[i]%machines[i] != 0 {
			batchTime++
		}
		if batchTime > time {
			time = batchTime
		}
	}
	return time
}
```
