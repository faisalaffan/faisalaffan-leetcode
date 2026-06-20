# 2187 — Minimum Time To Complete Trips

## Deskripsi

**Soal:** [2187. Minimum Time To Complete Trips](https://leetcode.com/problems/minimum-time-to-complete-trips/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log m)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func minimumTime(time []int, totalTrips int) int64`

## Solusi Go

```go
package main

// LeetCode #2187: Minimum Time to Complete Trips
// https://leetcode.com/problems/minimum-time-to-complete-trips/
// Difficulty: Medium
// Time: O(n log m) | Space: O(1)

import "fmt"

func minimumTime(time []int, totalTrips int) int64 {
	lo, hi := int64(1), int64(1)
	for _, t := range time {
		if int64(t) > hi {
			hi = int64(t)
		}
	}
	hi *= int64(totalTrips)

	for lo < hi {
		mid := lo + (hi-lo)/2
		trips := int64(0)
		for _, t := range time {
			trips += mid / int64(t)
			if trips >= int64(totalTrips) {
				break
			}
		}
		if trips >= int64(totalTrips) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func main() {
	// Test case 1
	fmt.Println(minimumTime([]int{1, 2, 3}, 5))
	// Expected: 3

	// Test case 2
	fmt.Println(minimumTime([]int{2}, 1))
	// Expected: 2

	// Test case 3
	fmt.Println(minimumTime([]int{5, 10, 10}, 9))
	// Expected: 25
}
```
