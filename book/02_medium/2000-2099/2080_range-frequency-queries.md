# 2080 — Range Frequency Queries

## Deskripsi

**Soal:** [2080. Range Frequency Queries](https://leetcode.com/problems/range-frequency-queries/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n) init, O(log n) per query  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func Constructor(arr []int) RangeFreqQuery`

## Solusi Go

```go
package main

// LeetCode #2080: Range Frequency Queries
// https://leetcode.com/problems/range-frequency-queries/
// Difficulty: Medium
// Time: O(n) init, O(log n) per query | Space: O(n)

import (
	"fmt"
	"sort"
)

type RangeFreqQuery struct {
	pos map[int][]int
}

func Constructor(arr []int) RangeFreqQuery {
  // Membuat map untuk pencarian O(1): key → value
	pos := make(map[int][]int)
	for i, v := range arr {
		pos[v] = append(pos[v], i)
	}
	return RangeFreqQuery{pos: pos}
}

func (rfq *RangeFreqQuery) Query(left int, right int, value int) int {
	positions, ok := rfq.pos[value]
	if !ok {
		return 0
	}
	// First index >= left
	l := sort.Search(len(positions), func(i int) bool {
		return positions[i] >= left
	})
	// First index > right
	r := sort.Search(len(positions), func(i int) bool {
		return positions[i] > right
	})
	return r - l
}

func main() {
	rfq := Constructor([]int{12, 33, 4, 56, 22, 2, 34, 33, 22, 12, 34, 56})
	fmt.Println("Test 1 Query(1, 2, 4):", rfq.Query(1, 2, 4))   // 1
	fmt.Println("Test 2 Query(0, 11, 33):", rfq.Query(0, 11, 33)) // 2
	fmt.Println("Test 3 Query(0, 5, 99):", rfq.Query(0, 5, 99))   // 0

	rfq2 := Constructor([]int{1, 1, 1, 2, 2})
	fmt.Println("Test 4 Query(0, 2, 1):", rfq2.Query(0, 2, 1)) // 3
	fmt.Println("Test 5 Query(3, 4, 2):", rfq2.Query(3, 4, 2)) // 2
}
```
