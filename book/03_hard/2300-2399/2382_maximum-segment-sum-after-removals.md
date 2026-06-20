# 2382 — Maximum Segment Sum After Removals

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan pohon (tree). Tugasmu menjelajahi atau memanipulasi struktur pohon.

**Cara berpikir:** TreeNode punya Val, Left, Right. Gunakan DFS rekursif (pre/in/post-order).

**Fungsi Solusi:** `func maximumSegmentSum(nums []int, removeQueries []int) []int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, Union-Find

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2382: Maximum Segment Sum After Removals
// https://leetcode.com/problems/maximum-segment-sum-after-removals/
// Difficulty: Hard
//
// You are given two 0-indexed arrays nums and removeQueries of length n.
// For each i from 0 to n-1, remove the element at position removeQueries[i]
// (making it inactive). After each removal, find the maximum segment sum
// among the active elements (a segment is a contiguous block of active elements).
//
// Approach: Process removals in reverse. Start with all elements removed and
// add them back one by one. Use DSU (Union-Find) to merge adjacent active
// segments, tracking the sum of each segment. Keep a running max.

import "fmt"

func maximumSegmentSum(nums []int, removeQueries []int) []int64 {
	n := len(nums)

  // Alokasi slice
	parent := make([]int, n)
  // Alokasi slice
	segSum := make([]int64, n)
	active := make([]bool, n)

	for i := 0; i < n; i++ {
		parent[i] = -1
	}

	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

  // Alokasi slice
	ans := make([]int64, n)
	var maxSum int64

	for i := n - 1; i >= 0; i-- {
		ans[i] = maxSum
		idx := removeQueries[i]
		active[idx] = true
		parent[idx] = idx
		segSum[idx] = int64(nums[idx])

		// Merge with left neighbor
		if idx > 0 && active[idx-1] {
			leftRoot := find(idx - 1)
			if leftRoot != idx {
				segSum[idx] += segSum[leftRoot]
				parent[leftRoot] = idx
			}
		}

		// Merge with right neighbor
		if idx < n-1 && active[idx+1] {
			rightRoot := find(idx + 1)
			if rightRoot != idx {
				segSum[idx] += segSum[rightRoot]
				parent[rightRoot] = idx
			}
		}

		if segSum[idx] > maxSum {
			maxSum = segSum[idx]
		}
	}

	return ans
}

func main() {
	// Example 1
	nums1 := []int{1, 2, 5, 6, 1}
	queries1 := []int{0, 3, 2, 4, 1}
	fmt.Println(maximumSegmentSum(nums1, queries1))

	// Example 2
	nums2 := []int{5, -2, 4, 1}
	queries2 := []int{0, 2, 1, 3}
	fmt.Println(maximumSegmentSum(nums2, queries2))

	// Edge: single element
	nums3 := []int{10}
	queries3 := []int{0}
	fmt.Println(maximumSegmentSum(nums3, queries3))

	// Edge: all negative
	nums4 := []int{-1, -2, -3, -4}
	queries4 := []int{0, 1, 2, 3}
	fmt.Println(maximumSegmentSum(nums4, queries4))
}
```
