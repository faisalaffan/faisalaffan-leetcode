# 2142 — The Number Of Passengers In Each Bus I

## Deskripsi

**Soal:** [2142. The Number Of Passengers In Each Bus I](https://leetcode.com/problems/the-number-of-passengers-in-each-bus-i/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n + m log m)  
**Kompleksitas Ruang:** O(n + m)

**Algoritma:** —

**Fungsi Solusi:** `func busPassengers(buses [][]int, passengers []int) []int`

## Solusi Go

```go
package main

// LeetCode #2142: The Number of Passengers in Each Bus I
// https://leetcode.com/problems/the-number-of-passengers-in-each-bus-i/
// Difficulty: Medium [Paid]
// Time: O(n log n + m log m) | Space: O(n + m)

import (
	"fmt"
	"sort"
)

type Bus struct {
	ID      int
	Arrival int
}

func busPassengers(buses [][]int, passengers []int) []int {
	// Sort buses by arrival time
	sort.Slice(buses, func(i, j int) bool {
		return buses[i][1] < buses[j][1]
	})

	// Sort passengers
	sort.Ints(passengers)

  // Membuat slice untuk menyimpan hasil
	result := make([]int, len(buses))
	pIdx := 0

	for i, bus := range buses {
		capacity := bus[0]
		count := 0
		for pIdx < len(passengers) && count < capacity && passengers[pIdx] <= bus[1] {
			count++
			pIdx++
		}
		result[i] = count
	}

	return result
}

func main() {
	// Test case 1: buses[capacity, arrival_time]
	buses1 := [][]int{{2, 10}, {3, 20}}
	passengers1 := []int{3, 8, 15, 18, 25}
	fmt.Println("Test 1:", busPassengers(buses1, passengers1))
	// Expected: [1, 2]

	// Test case 2
	buses2 := [][]int{{1, 5}, {2, 10}}
	passengers2 := []int{3, 8, 12}
	fmt.Println("Test 2:", busPassengers(buses2, passengers2))
	// Expected: [1, 1]

	// Test case 3
	buses3 := [][]int{{5, 100}}
	passengers3 := []int{}
	fmt.Println("Test 3:", busPassengers(buses3, passengers3))
	// Expected: [0]
}
```
