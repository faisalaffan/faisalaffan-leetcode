# 3907 — Count Smaller Elements With Opposite Parity

## Deskripsi

**Soal:** [3907. Count Smaller Elements With Opposite Parity](https://leetcode.com/problems/count-smaller-elements-with-opposite-parity/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(N log M)  
**Kompleksitas Ruang:** O(M) where M = max value

**Algoritma:** —

**Fungsi Solusi:** `func NewBIT(size int) *BIT`

> **Ide Kunci:** Process right to left. Use two BITs (even, odd) to count smaller

## Solusi Go

```go
package main

// LeetCode #3907: Count Smaller Elements With Opposite Parity
// https://leetcode.com/problems/count-smaller-elements-with-opposite-parity/
// Difficulty: Medium [Paid]
// Time: O(N log M) | Space: O(M) where M = max value
// Approach: Process right to left. Use two BITs (even, odd) to count smaller
// elements with opposite parity.

import (
	"fmt"
	"sort"
)

type BIT struct {
	tree []int
}

func NewBIT(size int) *BIT {
	return &BIT{tree: make([]int, size+2)}
}

func (b *BIT) Update(idx int, val int) {
	idx++
	for idx < len(b.tree) {
		b.tree[idx] += val
		idx += idx & -idx
	}
}

func (b *BIT) Query(idx int) int {
	idx++
	sum := 0
	for idx > 0 {
		sum += b.tree[idx]
		idx -= idx & -idx
	}
	return sum
}

func CountSmallerElementsWithOppositeParity(nums []int) []int {
	n := len(nums)
  // Membuat slice untuk menyimpan hasil
	ans := make([]int, n)

	// Coordinate compress
  // Membuat slice untuk menyimpan hasil
	sorted := make([]int, n)
	copy(sorted, nums)
	sort.Ints(sorted)
  // Membuat map untuk pencarian O(1): key → value
	rank := make(map[int]int)
	for i, v := range sorted {
		rank[v] = i
	}

	evenBit := NewBIT(n)
	oddBit := NewBIT(n)

	for i := n - 1; i >= 0; i-- {
		r := rank[nums[i]]
		if nums[i]%2 == 0 {
			// Count odd elements smaller than nums[i]
			ans[i] = oddBit.Query(r - 1)
			evenBit.Update(r, 1)
		} else {
			// Count even elements smaller than nums[i]
			ans[i] = evenBit.Query(r - 1)
			oddBit.Update(r, 1)
		}
	}

	return ans
}

func main() {
	// Example 1
	fmt.Println(CountSmallerElementsWithOppositeParity([]int{5, 2, 4, 1, 3})) // Expected: [2 1 2 0 0]

	// Example 2
	fmt.Println(CountSmallerElementsWithOppositeParity([]int{4, 4, 1})) // Expected: [1 1 0]

	// Example 3
	fmt.Println(CountSmallerElementsWithOppositeParity([]int{7})) // Expected: [0]
}
```
