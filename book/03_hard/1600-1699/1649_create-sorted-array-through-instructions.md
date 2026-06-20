# 1649 — Create Sorted Array Through Instructions

## Deskripsi

**Soal:** [1649. Create Sorted Array Through Instructions](https://leetcode.com/problems/create-sorted-array-through-instructions/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Fenwick Tree (Binary Indexed Tree)

**Fungsi Solusi:** `func NewBIT(n int) *BIT`

## Solusi Go

```go
package main

// LeetCode #1649: Create Sorted Array through Instructions
// https://leetcode.com/problems/create-sorted-array-through-instructions/
// Difficulty: Hard

import "fmt"

const MOD1649 = 1000000007

// Fenwick Tree (Binary Indexed Tree)
type BIT struct {
	tree []int
	n    int
}

func NewBIT(n int) *BIT {
	return &BIT{tree: make([]int, n+1), n: n}
}

func (b *BIT) Add(idx, val int) {
	for i := idx; i <= b.n; i += i & -i {
		b.tree[i] += val
	}
}

func (b *BIT) Sum(idx int) int {
	if idx <= 0 {
		return 0
	}
	if idx > b.n {
		idx = b.n
	}
	s := 0
	for i := idx; i > 0; i -= i & -i {
		s += b.tree[i]
	}
	return s
}

func createSortedArray(instructions []int) int {
	// Find maximum value to size the BIT
	maxVal := 0
	for _, v := range instructions {
		if v > maxVal {
			maxVal = v
		}
	}

	bit := NewBIT(maxVal)
	totalCost := 0
	totalInserted := 0

	for _, v := range instructions {
		less := bit.Sum(v - 1)               // elements < v
		total := totalInserted                // total elements so far
		greater := total - bit.Sum(v)         // elements > v

		cost := less
		if greater < cost {
			cost = greater
		}
		totalCost = (totalCost + cost) % MOD1649

		bit.Add(v, 1)
		totalInserted++
	}

	return totalCost
}

func main() {
	// Test case 1: [1,5,2,6,3] -> 3
	// 1: cost 0, 5: cost 0, 2: cost min(1,1)=1, 6: cost 0, 3: cost min(2,2)=2
	instructions := []int{1, 5, 2, 6, 3}
	result := createSortedArray(instructions)
	fmt.Printf("instructions=%v -> %d (expected 3)\n", instructions, result)

	// Test case 2: [1,2,3,4,5] -> 0 (inserted in order, no cost)
	instructions2 := []int{1, 2, 3, 4, 5}
	result2 := createSortedArray(instructions2)
	fmt.Printf("instructions=%v -> %d (expected 0)\n", instructions2, result2)

	// Test case 3: [1,2,1,2,1] -> 0
	// 1: cost 0, 2: cost 0, 1: cost min(0,1)=0, 2: cost min(1,0)=0, 1: cost min(0,2)=0
	instructions3 := []int{1, 2, 1, 2, 1}
	result3 := createSortedArray(instructions3)
	fmt.Printf("instructions=%v -> %d (expected 0)\n", instructions3, result3)

	// Test case 4: [5,4,3,2,1] -> 0 (all descending, always insert at rightmost)
	instructions4 := []int{5, 4, 3, 2, 1}
	result4 := createSortedArray(instructions4)
	fmt.Printf("instructions=%v -> %d\n", instructions4, result4)
}
```
