# 3593 — Minimum Increments To Equalize Leaf Paths

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan pohon (tree). Tugasmu menjelajahi atau memanipulasi struktur pohon.

**Cara berpikir:** TreeNode punya Val, Left, Right. Gunakan DFS rekursif (pre/in/post-order).

**Fungsi Solusi:** `func MinimumIncrementsToEqualizeLeafPaths(tree []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3593: Minimum Increments to Equalize Leaf Paths
// https://leetcode.com/problems/minimum-increments-to-equalize-leaf-paths/
// Difficulty: Medium
// Complexity: O(n) time, O(h) space

import "fmt"

func main() {
	// Test case 1: binary tree represented as array
	tree := []int{1, 2, 3}
	fmt.Println("Test 1:", MinimumIncrementsToEqualizeLeafPaths(tree))
	// Test case 2
	tree2 := []int{1, 2, 3, 4, 5}
	fmt.Println("Test 2:", MinimumIncrementsToEqualizeLeafPaths(tree2))
	// Test case 3
	tree3 := []int{1}
	fmt.Println("Test 3:", MinimumIncrementsToEqualizeLeafPaths(tree3))
}

func MinimumIncrementsToEqualizeLeafPaths(tree []int) int {
	n := len(tree)
	if n <= 1 {
		return 0
	}
	// Find max path sum from root to leaf
	// For each leaf, compute path sum and find max
	maxSum := 0
  // Alokasi slice
	pathSums := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		left := 2*i + 1
		right := 2*i + 2
		if left >= n && right >= n {
			pathSums[i] = tree[i]
		} else {
			childMax := 0
			if left < n && pathSums[left] > childMax {
				childMax = pathSums[left]
			}
			if right < n && pathSums[right] > childMax {
				childMax = pathSums[right]
			}
			pathSums[i] = tree[i] + childMax
		}
		if pathSums[i] > maxSum {
			maxSum = pathSums[i]
		}
	}

	// Count increments needed
	increments := 0
	for i := n - 1; i >= 0; i-- {
		left := 2*i + 1
		right := 2*i + 2
		needed := maxSum
		if left >= n && right >= n {
			needed = maxSum
		} else {
			if left < n {
				diff := maxSum - pathSums[left]
				increments += diff
				pathSums[left] += diff
			}
			if right < n {
				diff := maxSum - pathSums[right]
				increments += diff
				pathSums[right] += diff
			}
			needed = tree[i]
			if left < n && pathSums[left] < needed {
				needed = pathSums[left]
			}
			if right < n && pathSums[right] < needed {
				needed = pathSums[right]
			}
		}
		_ = needed // placeholder
	}
	return increments
}
```
