package main

// LeetCode #2940: Find Building Where Alice and Bob Can Meet
// https://leetcode.com/problems/find-building-where-alice-and-bob-can-meet/
// Difficulty: Hard
//
// Approach: Segment tree for range maximum + binary search.
// Alice at a can reach building m (m > a) iff heights[m] > heights[a].
// Bob at b can reach m (m > b) iff heights[m] > heights[b].
// We need the smallest m >= max(a,b) with height > max(heights[a], heights[b]).
// Use a segment tree storing max heights for range queries and binary search
// for the leftmost valid index.

import (
	"fmt"
)

type segTree struct {
	n  int
	tr []int
}

func newSegTree(heights []int) *segTree {
	n := len(heights)
	tr := make([]int, 4*n)
	var build func(idx, l, r int)
	build = func(idx, l, r int) {
		if l == r {
			tr[idx] = heights[l]
			return
		}
		mid := (l + r) / 2
		build(idx*2, l, mid)
		build(idx*2+1, mid+1, r)
		if tr[idx*2] > tr[idx*2+1] {
			tr[idx] = tr[idx*2]
		} else {
			tr[idx] = tr[idx*2+1]
		}
	}
	build(1, 0, n-1)
	return &segTree{n: n, tr: tr}
}

func (st *segTree) queryMax(idx, l, r, ql, qr int) int {
	if ql <= l && r <= qr {
		return st.tr[idx]
	}
	mid := (l + r) / 2
	res := 0
	if ql <= mid {
		if v := st.queryMax(idx*2, l, mid, ql, qr); v > res {
			res = v
		}
	}
	if qr > mid {
		if v := st.queryMax(idx*2+1, mid+1, r, ql, qr); v > res {
			res = v
		}
	}
	return res
}

func (st *segTree) rangeMax(l, r int) int {
	if l > r {
		return 0
	}
	return st.queryMax(1, 0, st.n-1, l, r)
}

func leftmostBuilding(heights []int, queries [][]int) []int {
	n := len(heights)
	st := newSegTree(heights)
	ans := make([]int, len(queries))

	for qi, q := range queries {
		a, b := q[0], q[1]
		if a == b {
			ans[qi] = a
			continue
		}
		if a > b {
			a, b = b, a // ensure a <= b
		}

		// If Alice can jump directly to Bob's building
		if heights[a] < heights[b] {
			ans[qi] = b
			continue
		}

		// Need a building to the right of b with height > max(heights[a], heights[b])
		target := heights[a] // heights[a] >= heights[b] since we checked heights[a] < heights[b] case

		// Binary search for the leftmost index m in [b+1, n-1] with height > target
		lo, hi := b+1, n-1
		res := -1
		for lo <= hi {
			mid := (lo + hi) / 2
			if st.rangeMax(b+1, mid) > target {
				res = mid
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		}
		ans[qi] = res
	}
	return ans
}

func main() {
	// Example: heights=[6,4,8,5,2,7], queries=[[0,1],[0,3],[2,4],[3,4],[2,2]]
	// Expected: [2,5,-1,5,2]
	heights := []int{6, 4, 8, 5, 2, 7}
	queries := [][]int{{0, 1}, {0, 3}, {2, 4}, {3, 4}, {2, 2}}
	fmt.Println(leftmostBuilding(heights, queries))

	// Simple case
	fmt.Println(leftmostBuilding([]int{1, 2, 3, 4}, [][]int{{0, 3}}))
}
