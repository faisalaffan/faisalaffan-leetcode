# 3237 — Alt And Tab Simulation

## Deskripsi

**Soal:** [3237. Alt And Tab Simulation](https://leetcode.com/problems/alt-and-tab-simulation/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n + q)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** LIS (Longest Increasing Subsequence)

**Fungsi Solusi:** `func simulationResult(windows []int, queries []int) []int`

## Solusi Go

```go
package main

// LeetCode #3237: Alt and Tab Simulation
// https://leetcode.com/problems/alt-and-tab-simulation/
// Difficulty: Medium [Paid]
// Time: O(n + q) | Space: O(n)

import (
	"container/list"
	"fmt"
)

func simulationResult(windows []int, queries []int) []int {
	order := list.New()
  // Membuat map untuk pencarian O(1): key → value
	pos := make(map[int]*list.Element)
	for _, w := range windows {
		e := order.PushBack(w)
		pos[w] = e
	}

	for _, q := range queries {
		if e, ok := pos[q]; ok {
			order.MoveToFront(e)
		}
	}

  // Membuat slice untuk menyimpan hasil
	ans := make([]int, 0, order.Len())
	for e := order.Front(); e != nil; e = e.Next() {
		ans = append(ans, e.Value.(int))
	}
	return ans
}

func main() {
	fmt.Println(simulationResult([]int{1, 2, 3, 4}, []int{3, 1})) // Expected: [1 3 2 4]
	fmt.Println(simulationResult([]int{1, 2, 3}, []int{2}))        // Expected: [2 1 3]
}
```
