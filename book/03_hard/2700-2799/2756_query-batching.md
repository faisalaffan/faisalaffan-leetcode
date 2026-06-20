# 2756 — Query Batching

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func queryBatching(queries []int, batchSize int, batchTime int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2756: Query Batching
// https://leetcode.com/problems/query-batching/
// Difficulty: Hard [Paid]
//
// Approach: Process queries in batches of up to batchSize.
// Each batch takes batchTime to execute. Queries arriving after the batch
// started wait for the next batch. Return total time to process all queries.

import (
	"fmt"
	"sort"
)

type query struct {
	arrival int
	idx     int
}

func queryBatching(queries []int, batchSize int, batchTime int) []int {
	n := len(queries)
	// Sort by arrival time, tracking original index
	sorted := make([]query, n)
	for i, t := range queries {
		sorted[i] = query{t, i}
	}
  // Custom sort
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].arrival < sorted[j].arrival
	})

  // Alokasi slice
	result := make([]int, n)
	time := 0
	ptr := 0

	for ptr < n {
		// Determine the start of this batch
		if time < sorted[ptr].arrival {
			time = sorted[ptr].arrival
		}

		// Collect up to batchSize queries that have arrived
		end := ptr + batchSize
		if end > n {
			end = n
		}

		// Oldest queries in this batch finish at time + batchTime
		time += batchTime
		for i := ptr; i < end; i++ {
			result[sorted[i].idx] = time
		}
		ptr = end
	}

	return result
}

func main() {
	// Example: queries arriving at [0, 1, 2, 5, 6], batchSize=2, batchTime=3
	// Batch 1: [0,1] finish at 3
	// Batch 2: [2,5] start at 5? No, start at max(3,2)=3, finish at 6
	// Batch 3: [6] start at max(6,6)=6, finish at 9
	fmt.Println(queryBatching([]int{0, 1, 2, 5, 6}, 2, 3))

	// Single query
	fmt.Println(queryBatching([]int{0}, 1, 5))

	// All arrive at same time
	fmt.Println(queryBatching([]int{0, 0, 0, 0}, 2, 3))

	// Staggered arrivals, full batch
	fmt.Println(queryBatching([]int{0, 0, 10, 10}, 2, 1))

	// Empty
	fmt.Println(queryBatching([]int{}, 2, 3))
}
```
