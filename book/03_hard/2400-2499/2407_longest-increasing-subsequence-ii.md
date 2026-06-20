# 2407 — Longest Increasing Subsequence Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func NewSegTree(size int) *SegTree`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Binary Search, DP, Segment Tree

**Waktu:** O(N log M) where M = max(nums).  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Binary Search** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2407: Longest Increasing Subsequence II
// https://leetcode.com/problems/longest-increasing-subsequence-ii/
// Difficulty: Hard
//
// Given an array nums and integer k, find the longest increasing subsequence
// where the difference between consecutive elements is at most k.
//
// Approach: Segment tree over values (1..max(nums)).
// For each value x = nums[i], the longest LIS ending at x is:
//   dp[x] = 1 + max(dp[x-k .. x-1])   (if such previous element exists)
// Use a segment tree to query the max in range [x-k, x-1] in O(log M).
// Time: O(N log M) where M = max(nums).

import "fmt"

type SegTree struct {
	tree []int
	n    int
}

func NewSegTree(size int) *SegTree {
	return &SegTree{
		tree: make([]int, 4*size),
		n:    size,
	}
}

func (st *SegTree) update(idx int, val int, node int, left int, right int) {
	if left == right {
		if val > st.tree[node] {
			st.tree[node] = val
		}
		return
	}
	mid := left + (right-left)/2
	if idx <= mid {
		st.update(idx, val, node*2, left, mid)
	} else {
		st.update(idx, val, node*2+1, mid+1, right)
	}
	st.tree[node] = max(st.tree[node*2], st.tree[node*2+1])
}

func (st *SegTree) query(ql int, qr int, node int, left int, right int) int {
	if ql > right || qr < left || ql > qr {
		return 0
	}
	if ql <= left && right <= qr {
		return st.tree[node]
	}
	mid := left + (right-left)/2
	return max(
		st.query(ql, qr, node*2, left, mid),
		st.query(ql, qr, node*2+1, mid+1, right),
	)
}

func lengthOfLIS(nums []int, k int) int {
  // Edge case: input kosong
	if len(nums) == 0 {
		return 0
	}

	maxVal := 0
	for _, v := range nums {
		if v > maxVal {
			maxVal = v
		}
	}

	st := NewSegTree(maxVal)
	ans := 1

	for _, x := range nums {
		// Query max dp in range [x-k, x-1]
		left := x - k
		if left < 1 {
			left = 1
		}
		best := st.query(left, x-1, 1, 1, maxVal)
		cur := best + 1
		st.update(x, cur, 1, 1, maxVal)
		if cur > ans {
			ans = cur
		}
	}

	return ans
}

func main() {
	// Example 1
	fmt.Println(lengthOfLIS([]int{4, 2, 1, 4, 3, 4, 5, 8, 15}, 3))
	// Example 2
	fmt.Println(lengthOfLIS([]int{7, 4, 5, 1, 8, 12, 4, 7}, 5))
	// Example 3
	fmt.Println(lengthOfLIS([]int{1, 5}, 1))
	// Edge: simple increasing
	fmt.Println(lengthOfLIS([]int{1, 2, 3, 4}, 1))
	// Edge: single element
	fmt.Println(lengthOfLIS([]int{10}, 3))
}
```
