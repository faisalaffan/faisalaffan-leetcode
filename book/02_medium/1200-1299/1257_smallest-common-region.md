# 1257 — Smallest Common Region

## Deskripsi

**Soal:** [1257. Smallest Common Region](https://leetcode.com/problems/smallest-common-region/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n) where n = total regions  
**Kompleksitas Ruang:** O(n)

**Algoritma:** LIS (Longest Increasing Subsequence)

**Fungsi Solusi:** `func findSmallestRegion(regions [][]string, region1 string, region2 string) string`

## Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #1257: Smallest Common Region
// https://leetcode.com/problems/smallest-common-region/
// Difficulty: Medium [Paid]

// Given region hierarchy, find smallest common region.
// Build parent map, then find common ancestor.

// Time: O(n) where n = total regions
// Space: O(n)

func findSmallestRegion(regions [][]string, region1 string, region2 string) string {
  // Membuat map untuk pencarian O(1): key → value
	parent := make(map[string]string)

	for _, list := range regions {
		for i := 1; i < len(list); i++ {
			parent[list[i]] = list[0]
		}
	}

	// Find path from region1 to root
  // Membuat map untuk pencarian O(1): key → value
	path := make(map[string]bool)
	r := region1
	path[r] = true
	for {
		p, exists := parent[r]
		if !exists {
			break
		}
		path[p] = true
		r = p
	}

	// Find common ancestor
	r = region2
	for {
		if path[r] {
			return r
		}
		r = parent[r]
	}
}

func main() {
	regions := [][]string{
		{"Earth", "North America", "South America"},
		{"North America", "USA", "Canada"},
		{"USA", "California", "Texas"},
		{"Canada", "Ontario", "Quebec"},
	}
	fmt.Printf("%q (expected: %q)\n",
		findSmallestRegion(regions, "California", "Ontario"), "North America")

	fmt.Printf("%q (expected: %q)\n",
		findSmallestRegion(regions, "California", "Texas"), "USA")
}
```
