# 2916 — Subarrays Distinct Element Sum Of Squares Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func newSegTree(n int) *segTree`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Binary Search, Segment Tree

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2916: Subarrays Distinct Element Sum of Squares II
// https://leetcode.com/problems/subarrays-distinct-element-sum-of-squares-ii/
// Difficulty: Hard
//
// Compute sum over all subarrays of (number of distinct elements)^2.
// Use a segment tree with lazy propagation over positions. For each right endpoint j,
// we maintain for each left endpoint i the distinct count g(i,j) in subarray [i..j].
// When extending by element nums[j+1] with value v, increment g(i,j+1) by 1 for all
// i where v was not present before (i.e., i in (last[v], j+1]).
// Segment tree supports range increment and range sum query for both g and g^2.
// O(N log N) time, O(N) space.

import "fmt"

const mod2916 = 1000000007

type node struct {
	sum, sq, lazy int64
}

type segTree struct {
	tree []node
	n    int
}

func newSegTree(n int) *segTree {
	return &segTree{
		tree: make([]node, 4*n),
		n:    n,
	}
}

func (st *segTree) apply(idx, l, r int, val int64) {
	length := int64(r - l + 1)
	nd := &st.tree[idx]
	// (x+v)^2 = x^2 + 2*x*v + v^2
	// sum_sq += 2*val*sum + val*val*len
	nd.sq = (nd.sq + 2*val*nd.sum%mod2916 + val*val%mod2916*length%mod2916) % mod2916
	nd.sum = (nd.sum + val*length) % mod2916
	nd.lazy = (nd.lazy + val) % mod2916
}

func (st *segTree) push(idx, l, r int) {
	if st.tree[idx].lazy != 0 && l != r {
		mid := (l + r) / 2
		st.apply(idx*2, l, mid, st.tree[idx].lazy)
		st.apply(idx*2+1, mid+1, r, st.tree[idx].lazy)
		st.tree[idx].lazy = 0
	}
}

func (st *segTree) rangeAdd(idx, l, r, ql, qr int, val int64) {
	if ql > r || qr < l {
		return
	}
	if ql <= l && r <= qr {
		st.apply(idx, l, r, val)
		return
	}
	st.push(idx, l, r)
	mid := (l + r) / 2
	st.rangeAdd(idx*2, l, mid, ql, qr, val)
	st.rangeAdd(idx*2+1, mid+1, r, ql, qr, val)
	st.tree[idx].sum = (st.tree[idx*2].sum + st.tree[idx*2+1].sum) % mod2916
	st.tree[idx].sq = (st.tree[idx*2].sq + st.tree[idx*2+1].sq) % mod2916
}

func (st *segTree) rangeQuery(idx, l, r, ql, qr int) (int64, int64) {
	if ql > r || qr < l {
		return 0, 0
	}
	if ql <= l && r <= qr {
		return st.tree[idx].sum, st.tree[idx].sq
	}
	st.push(idx, l, r)
	mid := (l + r) / 2
	s1, sq1 := st.rangeQuery(idx*2, l, mid, ql, qr)
	s2, sq2 := st.rangeQuery(idx*2+1, mid+1, r, ql, qr)
	return (s1 + s2) % mod2916, (sq1 + sq2) % mod2916
}

func sumCounts(nums []int) int {
	n := len(nums)
  // HashMap: O(1) lookup
	last := make(map[int]int)
	seg := newSegTree(n)

	var ans int64 = 0

	for j, val := range nums {
		left := 0
		if prevIdx, ok := last[val]; ok {
			left = prevIdx + 1
		}
		// Increment g(i,j) by 1 for all i in [left, j]
		seg.rangeAdd(1, 0, n-1, left, j, 1)

		// Query sum of squares over all i <= j
		_, sumSq := seg.rangeQuery(1, 0, n-1, 0, j)
		ans = (ans + sumSq) % mod2916

		last[val] = j
	}

	return int(ans)
}

func main() {
	// Example: [1,2,1]
	// Subarrays: [1]->1, [2]->1, [1]->1, [1,2]->2, [2,1]->2, [1,2,1]->2
	// Sum of squares: 1+1+1+4+4+4 = 15
	fmt.Println(sumCounts([]int{1, 2, 1}))

	// Example: [1,1]
	// Subarrays: [1]->1, [1]->1, [1,1]->1
	// Sum: 1+1+1 = 3
	fmt.Println(sumCounts([]int{1, 1}))

	// Example: [1,2,3]
	// All distinct: each subarray has 1+2+...+n distinct = n(n+1)/2
	// Actually: [1]=1, [2]=1, [3]=1, [1,2]=2, [2,3]=2, [1,2,3]=3
	// Sum: 1+1+1+4+4+9 = 20
	fmt.Println(sumCounts([]int{1, 2, 3}))

	// Single element
	fmt.Println(sumCounts([]int{5}))

	// Repeating pattern
	fmt.Println(sumCounts([]int{1, 2, 3, 2}))

	// All same
	fmt.Println(sumCounts([]int{2, 2, 2}))

	// Longer
	fmt.Println(sumCounts([]int{1, 2, 3, 4, 5}))
}
```
