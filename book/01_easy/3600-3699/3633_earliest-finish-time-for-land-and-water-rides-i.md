# 3633 — Earliest Finish Time For Land And Water Rides I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func EarliestFinishTimeForLandAndWaterRidesI(landStartTime, landDuration, waterStartTime, waterDuration []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n*m) where n = len(landStartTime), m = len(waterStartTime)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3633: Earliest Finish Time for Land and Water Rides I
// https://leetcode.com/problems/earliest-finish-time-for-land-and-water-rides-i/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(EarliestFinishTimeForLandAndWaterRidesI([]int{2, 8}, []int{4, 1}, []int{6}, []int{3}))
	fmt.Println(EarliestFinishTimeForLandAndWaterRidesI([]int{5}, []int{3}, []int{1}, []int{10}))
}

// Time: O(n*m) where n = len(landStartTime), m = len(waterStartTime)
// Space: O(1)
func EarliestFinishTimeForLandAndWaterRidesI(landStartTime, landDuration, waterStartTime, waterDuration []int) int {
	ans := int(^uint(0) >> 1) // max int

	// Land -> Water
  // Range loop
	for i := range landStartTime {
		finishLand := landStartTime[i] + landDuration[i]
		for j := range waterStartTime {
			startWater := max(finishLand, waterStartTime[j])
			finish := startWater + waterDuration[j]
			if finish < ans {
				ans = finish
			}
		}
	}

	// Water -> Land
  // Range loop
	for i := range waterStartTime {
		finishWater := waterStartTime[i] + waterDuration[i]
		for j := range landStartTime {
			startLand := max(finishWater, landStartTime[j])
			finish := startLand + landDuration[j]
			if finish < ans {
				ans = finish
			}
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```
