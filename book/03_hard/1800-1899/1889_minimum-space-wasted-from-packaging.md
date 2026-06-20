# 1889 — Minimum Space Wasted From Packaging

## Deskripsi

**Soal:** [1889. Minimum Space Wasted From Packaging](https://leetcode.com/problems/minimum-space-wasted-from-packaging/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

**Fungsi Solusi:** `func minWastedSpace(packages []int, boxes [][]int) int`

## Solusi Go

```go
package main

// LeetCode #1889: Minimum Space Wasted From Packaging
// https://leetcode.com/problems/minimum-space-wasted-from-packaging/
// Difficulty: Hard

import (
	"fmt"
	"math"
	"sort"
)

func minWastedSpace(packages []int, boxes [][]int) int {
	const mod = 1_000_000_007
	sort.Ints(packages)
	n := len(packages)
  // Membuat slice untuk menyimpan hasil
	prefix := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + packages[i]
	}

	ans := math.MaxInt64

	for _, supplier := range boxes {
		sort.Ints(supplier)
		if supplier[len(supplier)-1] < packages[n-1] {
			continue // cannot fit the largest package
		}
		waste := 0
		prevIdx := 0
		for _, box := range supplier {
			// find the last package that fits in this box
			idx := sort.Search(n-prevIdx, func(k int) bool {
				return packages[prevIdx+k] > box
			}) + prevIdx
			if idx > prevIdx {
				count := idx - prevIdx
				// waste = box * count - sum of packages in [prevIdx, idx)
				sum := prefix[idx] - prefix[prevIdx]
				waste += box*count - sum
				prevIdx = idx
				if waste > ans { // early break
					break
				}
			}
		}
		if prevIdx == n && waste < ans {
			ans = waste
		}
	}

	if ans == math.MaxInt64 {
		return -1
	}
	return ans % mod
}

func main() {
	// Example: packages=[2,3,5], boxes=[[4,8],[2,8]] -> 6
	fmt.Println(minWastedSpace([]int{2, 3, 5}, [][]int{{4, 8}, {2, 8}}))

	// Additional test
	fmt.Println(minWastedSpace([]int{3, 5, 8, 10, 11, 12}, [][]int{{12}, {11, 9}, {10, 5, 14}}))
}
```
