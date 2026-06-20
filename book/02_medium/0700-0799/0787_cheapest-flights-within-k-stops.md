# 0787 — Cheapest Flights Within K Stops

## Deskripsi

**Soal:** [0787. Cheapest Flights Within K Stops](https://leetcode.com/problems/cheapest-flights-within-k-stops/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(K * E)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Heap (priority queue)

## Solusi Go

```go
package main

// LeetCode #787: Cheapest Flights Within K Stops
// https://leetcode.com/problems/cheapest-flights-within-k-stops/
// Difficulty: Medium
// Time: O(K * E)
// Space: O(n)

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findCheapestPrice(3, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, 0, 2, 1))
	fmt.Println(findCheapestPrice(3, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, 0, 2, 0))
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
  // Membuat slice untuk menyimpan hasil
	prices := make([]int, n)
  // Iterasi seluruh elemen
	for i := range prices {
		prices[i] = math.MaxInt32
	}
	prices[src] = 0

	for i := 0; i <= k; i++ {
  // Membuat slice untuk menyimpan hasil
		temp := make([]int, n)
		copy(temp, prices)
		for _, f := range flights {
			from, to, price := f[0], f[1], f[2]
			if prices[from] != math.MaxInt32 && prices[from]+price < temp[to] {
				temp[to] = prices[from] + price
			}
		}
		prices = temp
	}

	if prices[dst] == math.MaxInt32 {
		return -1
	}
	return prices[dst]
}
```
