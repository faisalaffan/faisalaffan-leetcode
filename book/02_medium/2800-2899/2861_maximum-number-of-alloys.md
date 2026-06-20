# 2861 — Maximum Number Of Alloys

## Deskripsi

**Soal:** [2861. Maximum Number Of Alloys](https://leetcode.com/problems/maximum-number-of-alloys/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n * log m)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func MaximumNumberOfAlloys(n int, k int, budget int, composition [][]int, stock []int, cost []int) int`

## Solusi Go

```go
package main

// LeetCode #2861: Maximum Number of Alloys
// https://leetcode.com/problems/maximum-number-of-alloys/
// Difficulty: Medium
// Time: O(n * log m) | Space: O(1)

import "fmt"

func MaximumNumberOfAlloys(n int, k int, budget int, composition [][]int, stock []int, cost []int) int {
	canMake := func(count int, comp []int) bool {
		var totalCost int64
		for i := 0; i < n; i++ {
			needed := int64(comp[i]) * int64(count)
			if needed > int64(stock[i]) {
				totalCost += (needed - int64(stock[i])) * int64(cost[i])
				if totalCost > int64(budget) {
					return false
				}
			}
		}
		return totalCost <= int64(budget)
	}

	best := 0
	for _, comp := range composition {
		lo, hi := 0, 1_000_000_000
		for lo < hi {
			mid := (lo + hi + 1) / 2
			if canMake(mid, comp) {
				lo = mid
			} else {
				hi = mid - 1
			}
		}
		if lo > best {
			best = lo
		}
	}

	return best
}

func main() {
	fmt.Println(MaximumNumberOfAlloys(3, 2, 15, [][]int{{1, 1, 1}, {2, 2, 2}}, []int{0, 0, 0}, []int{1, 2, 3}))
	fmt.Println(MaximumNumberOfAlloys(2, 2, 10, [][]int{{1, 2}, {2, 1}}, []int{5, 5}, []int{1, 1}))
}
```
