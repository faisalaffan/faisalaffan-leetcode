# 3796 — Find Maximum Value In A Constrained Sequence

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func FindMaximumValueInAConstrainedSequence(n int, restrictions [][]int, diff []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3796: Find Maximum Value in a Constrained Sequence
// https://leetcode.com/problems/find-maximum-value-in-a-constrained-sequence/
// Difficulty: Medium
// Time: O(n) | Space: O(n)
// Approach: Two-pass greedy constraint propagation. Forward pass applies
// constraints from left to right, backward pass propagates from right to left.

import "fmt"

func FindMaximumValueInAConstrainedSequence(n int, restrictions [][]int, diff []int) int {
	const INF = 1 << 60
  // Alokasi slice
	a := make([]int, n)
  // Range loop
	for i := range a {
		a[i] = INF
	}
	a[0] = 0

	// Apply restrictions
	for _, r := range restrictions {
		idx, maxVal := r[0], r[1]
		if a[idx] > maxVal {
			a[idx] = maxVal
		}
	}

	// Forward pass: propagate constraints left to right
	for i := 1; i < n; i++ {
		limit := a[i-1] + diff[i-1]
		if a[i] > limit {
			a[i] = limit
		}
	}

	// Backward pass: propagate constraints right to left
	for i := n - 2; i >= 0; i-- {
		limit := a[i+1] + diff[i]
		if a[i] > limit {
			a[i] = limit
		}
	}

	// Find maximum value
	ans := 0
	for _, v := range a {
		if v > ans {
			ans = v
		}
	}
	return ans
}

func main() {
	// Example 1
	n1 := 10
	restrictions1 := [][]int{{3, 1}, {8, 1}}
	diff1 := []int{2, 2, 3, 1, 4, 5, 1, 1, 2}
	fmt.Println(FindMaximumValueInAConstrainedSequence(n1, restrictions1, diff1)) // Expected: 6

	// Example 2
	n2 := 8
	restrictions2 := [][]int{{3, 2}}
	diff2 := []int{3, 5, 2, 4, 2, 3, 1}
	fmt.Println(FindMaximumValueInAConstrainedSequence(n2, restrictions2, diff2)) // Expected: 12
}
```
