# 2234 — Maximum Total Beauty Of The Gardens

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func maximumBeauty(flowers []int, newFlowers int64, target int, full int, partial int) int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Binary Search, Sorting, Prefix Sum

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Binary Search** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2234: Maximum Total Beauty of the Gardens
// https://leetcode.com/problems/maximum-total-beauty-of-the-gardens/
// Difficulty: Hard
//
// Alice has n gardens, each garden has some flowers initially.
// She can plant at most newFlowers additional flowers (total).
// The beauty score = fullGardens * full + incompleteMin * partial
//   - fullGardens: number of gardens with at least target flowers
//   - incompleteMin: minimum number of flowers among incomplete gardens (0 if all full)
// Goal: maximize total beauty.

import (
	"fmt"
	"sort"
)

// maximumBeauty returns maximum possible total beauty.
func maximumBeauty(flowers []int, newFlowers int64, target int, full int, partial int) int64 {
	n := len(flowers)
	// convert to int for easier math
	newF := int(newFlowers)

	// sort initially
  // Alokasi slice
	sorted := make([]int, n)
	copy(sorted, flowers)
  // Sort O(n log n)
	sort.Ints(sorted)

	// clip at target (excess flowers don't help completeness)
	for i := 0; i < n; i++ {
		if sorted[i] > target {
			sorted[i] = target
		}
	}
  // Sort O(n log n)
	sort.Ints(sorted)

	// prefix sums for efficient gap calculation
  // Alokasi slice
	prefix := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + sorted[i]
	}

	// try all possible counts of full gardens
	best := int64(0)

	// suffix: index from where gardens are full
	for fullCount := 0; fullCount <= n; fullCount++ {
		// if we make last 'fullCount' gardens full, what remaining flowers do we have?
		remaining := newF

		// cost to raise all fullCount gardens to target
		if fullCount > 0 {
			startIdx := n - fullCount
			cost := fullCount*target - (prefix[n] - prefix[startIdx])
			if cost > remaining {
				continue // not enough flowers for this many full gardens
			}
			remaining -= cost
		}

		// now maximize the minimum among remaining (incomplete) gardens
		incompleteCount := n - fullCount
		if incompleteCount == 0 {
			// all gardens full
			beauty := n * full
			if int64(beauty) > best {
				best = int64(beauty)
			}
			continue
		}

		// binary search for maximum possible minimum
		incomplete := sorted[:incompleteCount]

		lo, hi := 0, target-1
		for lo <= hi {
			mid := (lo + hi) / 2

			// how many flowers needed to raise all incomplete to at least mid?
			need := 0
			pos := sort.SearchInts(incomplete, mid)
			need = pos*mid - prefix[pos]

			if need <= remaining {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}

		minBeauty := hi // hi is the max min achievable
		if minBeauty < 0 {
			continue
		}

		beauty := fullCount*full + minBeauty*partial
		if int64(beauty) > best {
			best = int64(beauty)
		}
	}

	return best
}

func main() {
	// Example 1
	fmt.Println(maximumBeauty([]int{1, 3, 1, 1}, 7, 6, 12, 1))
	// Expected: 14

	// Example 2
	fmt.Println(maximumBeauty([]int{2, 4, 5, 3}, 10, 5, 2, 6))
	// Expected: 30
}
```
