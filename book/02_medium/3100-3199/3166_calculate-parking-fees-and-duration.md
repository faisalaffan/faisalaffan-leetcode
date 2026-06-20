# 3166 — Calculate Parking Fees And Duration

## Deskripsi

**Soal:** [3166. Calculate Parking Fees And Duration](https://leetcode.com/problems/calculate-parking-fees-and-duration/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func calculateParkingFees(records [][]int) []int`

## Solusi Go

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
  // Membuat map untuk pencarian O(1): key → value
	fees := make(map[int]int)
  // Membuat map untuk pencarian O(1): key → value
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
	sort.Ints(cars)

  // Membuat slice untuk menyimpan hasil
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
