# 2921 — Maximum Profitable Triplets With Increasing Prices Ii

## Deskripsi

**Soal:** [2921. Maximum Profitable Triplets With Increasing Prices Ii](https://leetcode.com/problems/maximum-profitable-triplets-with-increasing-prices-ii/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Segment Tree (pohon segmen)

**Fungsi Solusi:** `func newSegTree2921(n int) *segTree2921`

> **Ide Kunci:** For each j as the middle element, find the best i (left of j with

## Solusi Go

```go
package main

// LeetCode #2921: Maximum Profitable Triplets With Increasing Prices II
// https://leetcode.com/problems/maximum-profitable-triplets-with-increasing-prices-ii/
// Difficulty: Hard [Paid]
//
// Find i < j < k such that prices[i] < prices[j] < prices[k], maximizing
// profits[i] + profits[j] + profits[k].
//
// Approach: For each j as the middle element, find the best i (left of j with
// lower price) and best k (right of j with higher price). Use segment trees
// keyed by compressed prices. One pass left-to-right to compute best left profit
// for each position, one pass right-to-left for best right profit.
// O(N log N) time, O(N) space.

import (
	"fmt"
	"sort"
)

const minInt = int(-1e15)

type segTreeNode struct {
	maxVal int
}

type segTree2921 struct {
	tree []segTreeNode
	n    int
}

func newSegTree2921(n int) *segTree2921 {
  // Membuat slice untuk menyimpan hasil
	tree := make([]segTreeNode, 4*n)
  // Iterasi seluruh elemen
	for i := range tree {
		tree[i].maxVal = minInt
	}
	return &segTree2921{tree: tree, n: n}
}

func (st *segTree2921) update(idx, l, r, pos, val int) {
	if l == r {
		if val > st.tree[idx].maxVal {
			st.tree[idx].maxVal = val
		}
		return
	}
	mid := (l + r) / 2
	if pos <= mid {
		st.update(idx*2, l, mid, pos, val)
	} else {
		st.update(idx*2+1, mid+1, r, pos, val)
	}
	if st.tree[idx*2].maxVal > st.tree[idx*2+1].maxVal {
		st.tree[idx].maxVal = st.tree[idx*2].maxVal
	} else {
		st.tree[idx].maxVal = st.tree[idx*2+1].maxVal
	}
}

func (st *segTree2921) query(idx, l, r, ql, qr int) int {
	if ql > r || qr < l || ql > qr {
		return minInt
	}
	if ql <= l && r <= qr {
		return st.tree[idx].maxVal
	}
	mid := (l + r) / 2
	leftMax := st.query(idx*2, l, mid, ql, qr)
	rightMax := st.query(idx*2+1, mid+1, r, ql, qr)
	if leftMax > rightMax {
		return leftMax
	}
	return rightMax
}

func maxProfitableTriplet(prices []int, profits []int) int {
	n := len(prices)
	// Coordinate compress prices
  // Membuat slice untuk menyimpan hasil
	sorted := make([]int, n)
	copy(sorted, prices)
	sort.Ints(sorted)
	m := 1
	for i := 1; i < n; i++ {
		if sorted[i] != sorted[m-1] {
			sorted[m] = sorted[i]
			m++
		}
	}
	sorted = sorted[:m]

	compress := func(p int) int {
		return sort.SearchInts(sorted, p) + 1 // 1-indexed
	}

	// Left pass: best profit for i < j with price[i] < price[j]
  // Membuat slice untuk menyimpan hasil
	leftBest := make([]int, n)
	segLeft := newSegTree2921(m)
	for j := 0; j < n; j++ {
		pos := compress(prices[j])
		best := segLeft.query(1, 1, m, 1, pos-1)
		if best == minInt {
			leftBest[j] = minInt
		} else {
			leftBest[j] = best
		}
		segLeft.update(1, 1, m, pos, profits[j])
	}

	// Right pass: best profit for k > j with price[k] > price[j]
  // Membuat slice untuk menyimpan hasil
	rightBest := make([]int, n)
	segRight := newSegTree2921(m)
	for j := n - 1; j >= 0; j-- {
		pos := compress(prices[j])
		best := segRight.query(1, 1, m, pos+1, m)
		if best == minInt {
			rightBest[j] = minInt
		} else {
			rightBest[j] = best
		}
		segRight.update(1, 1, m, pos, profits[j])
	}

	// Combine
	ans := minInt
	for j := 1; j < n-1; j++ {
		if leftBest[j] == minInt || rightBest[j] == minInt {
			continue
		}
		total := leftBest[j] + profits[j] + rightBest[j]
		if total > ans {
			ans = total
		}
	}

	if ans == minInt {
		return -1
	}
	return ans
}

func main() {
	// Example: prices=[10,20,30,40], profits=[1,2,3,4] => 1+2+4=7 or 1+3+4=8 or 2+3+4=9
	// i=0,j=1,k=3: 1+2+4=7, i=0,j=2,k=3: 1+3+4=8, i=1,j=2,k=3: 2+3+4=9
	fmt.Println(maxProfitableTriplet([]int{10, 20, 30, 40}, []int{1, 2, 3, 4}))

	// No valid triplet (prices not strictly increasing)
	fmt.Println(maxProfitableTriplet([]int{10, 10, 10}, []int{1, 2, 3}))

	// Reverse prices
	fmt.Println(maxProfitableTriplet([]int{30, 20, 10}, []int{3, 2, 1}))

	// Complex
	fmt.Println(maxProfitableTriplet([]int{5, 1, 4, 2, 3}, []int{10, 5, 8, 6, 7}))

	// Simple
	fmt.Println(maxProfitableTriplet([]int{1, 2, 3}, []int{1, 2, 3}))

	// Duplicate prices
	fmt.Println(maxProfitableTriplet([]int{1, 2, 2, 3}, []int{1, 5, 3, 4}))
}
```
