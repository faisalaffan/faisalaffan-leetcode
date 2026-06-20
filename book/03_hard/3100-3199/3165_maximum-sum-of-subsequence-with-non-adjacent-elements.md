# 3165 — Maximum Sum Of Subsequence With Non Adjacent Elements

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func merge(a, b Node) Node
```

> **💡 Hint:** segment tree with 4-state nodes (s00, s01, s10, s11) for O(log n)

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Segment Tree

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Segment Tree** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3165: Maximum Sum of Subsequence With Non-adjacent Elements
// https://leetcode.com/problems/maximum-sum-of-subsequence-with-non-adjacent-elements/
// Difficulty: Hard
//
// Given nums and queries [pos, val], update nums[pos]=val then compute the
// maximum sum of a subsequence with no adjacent elements (House Robber style).
// Each query returns the max sum after the update.
//
// Approach: segment tree with 4-state nodes (s00, s01, s10, s11) for O(log n)
// per update/query.

import "fmt"

const MOD = 1000000007

type Node struct {
	s00, s01, s10, s11 int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func merge(a, b Node) Node {
	return Node{
		s00: max(a.s00+b.s10, a.s01+b.s00),
		s01: max(a.s00+b.s11, a.s01+b.s01),
		s10: max(a.s10+b.s10, a.s11+b.s00),
		s11: max(a.s10+b.s11, a.s11+b.s01),
	}
}

type SegTree struct {
	tree []Node
	n    int
}

func NewSegTree(arr []int) *SegTree {
	n := len(arr)
	tree := make([]Node, 4*n)
	st := &SegTree{tree: tree, n: n}
	st.build(arr, 1, 0, n-1)
	return st
}

func (st *SegTree) build(arr []int, idx, l, r int) {
	if l == r {
		st.tree[idx] = Node{s11: max(arr[l], 0)}
		return
	}
	mid := (l + r) / 2
	st.build(arr, idx*2, l, mid)
	st.build(arr, idx*2+1, mid+1, r)
	st.tree[idx] = merge(st.tree[idx*2], st.tree[idx*2+1])
}

func (st *SegTree) update(idx, l, r, pos, val int) {
	if l == r {
		st.tree[idx] = Node{s11: max(val, 0)}
		return
	}
	mid := (l + r) / 2
	if pos <= mid {
		st.update(idx*2, l, mid, pos, val)
	} else {
		st.update(idx*2+1, mid+1, r, pos, val)
	}
	st.tree[idx] = merge(st.tree[idx*2], st.tree[idx*2+1])
}

func (st *SegTree) Query() int {
	return st.tree[1].s11 % MOD
}

func maximumSumSubsequence(nums []int, queries [][]int) []int {
  // Edge case: input kosong — langsung return
	if len(nums) == 0 {
		return make([]int, len(queries))
	}
	st := NewSegTree(nums)
  // Alokasi slice integer
	ans := make([]int, len(queries))
	for i, q := range queries {
		pos, val := q[0], q[1]
		st.update(1, 0, st.n-1, pos, val)
		ans[i] = st.Query()
	}
	return ans
}

func main() {
	nums := []int{3, 5, 9}
	queries := [][]int{{1, -2}, {0, -1}}
	fmt.Println(maximumSumSubsequence(nums, queries))
	// Expect: [9, 5]  (after each update, the max non-adjacent sum)
}
```
