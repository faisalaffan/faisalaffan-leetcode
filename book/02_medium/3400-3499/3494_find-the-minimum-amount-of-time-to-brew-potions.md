# 3494 — Find The Minimum Amount Of Time To Brew Potions

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindTheMinimumAmountOfTimeToBrewPotions(machines []int, potionTimes []int, potions int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

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
  // Loop linear O(n): iterasi setiap elemen
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
