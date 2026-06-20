# 2332 — The Latest Time To Catch A Bus

## Deskripsi

**Soal:** [2332. The Latest Time To Catch A Bus](https://leetcode.com/problems/the-latest-time-to-catch-a-bus/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O((n + m) log(n + m))  
**Kompleksitas Ruang:** O(n + m)

**Algoritma:** —

**Fungsi Solusi:** `func latestTimeCatchTheBus(buses []int, passengers []int, capacity int) int`

## Solusi Go

```go
package main

// LeetCode #2332: The Latest Time to Catch a Bus
// https://leetcode.com/problems/the-latest-time-to-catch-a-bus/
// Difficulty: Medium
// Time: O((n + m) log(n + m)) | Space: O(n + m)

import (
	"fmt"
	"sort"
)

func latestTimeCatchTheBus(buses []int, passengers []int, capacity int) int {
	sort.Ints(buses)
	sort.Ints(passengers)

	pi := 0
	for _, bus := range buses {
		count := 0
		for count < capacity && pi < len(passengers) && passengers[pi] <= bus {
			count++
			pi++
		}
		// Track last boarding time
		if pi > 0 && count == capacity {
			// Last passenger on this bus
			_ = 0
		}
	}

	// Find latest possible time
	time := buses[len(buses)-1]
  // Membuat map untuk pencarian O(1): key → value
	passSet := make(map[int]bool)
	for _, p := range passengers {
		passSet[p] = true
	}

	// If not all seats taken on last bus, try last bus departure
	// Check if we can arrive at bus time
	pi = 0
	lastBoarded := -1
	for _, bus := range buses {
		count := 0
		for count < capacity && pi < len(passengers) && passengers[pi] <= bus {
			lastBoarded = passengers[pi]
			count++
			pi++
		}
		if count < capacity {
			time = bus
		} else {
			time = lastBoarded - 1
		}
	}

	// Find latest time not taken by a passenger
	for passSet[time] {
		time--
	}
	return time
}

func main() {
	// Test case 1
	fmt.Println(latestTimeCatchTheBus([]int{10, 20}, []int{2, 17, 18, 19}, 2))
	// Expected: 16

	// Test case 2
	fmt.Println(latestTimeCatchTheBus([]int{20, 30, 10}, []int{19, 13, 26, 4, 25, 11, 21}, 2))
	// Expected: 20
}
```
