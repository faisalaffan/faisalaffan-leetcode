package main

// LeetCode #3165: Maximum Sum of Subsequence With Non-adjacent Elements
// https://leetcode.com/problems/maximum-sum-of-subsequence-with-non-adjacent-elements/
// Difficulty: Hard
//
// Given nums and queries [pos, val], update nums[pos]=val then compute the
// maximum sum of a subsequence with no adjacent elements (House Robber style).
// Each query returns the max sum after the update. Use a segment tree with
// 4-state nodes (s00, s01, s10, s11) for O(log n) per query.

import (
	"fmt"
)

const MOD = 1000000007

func main() {
	nums := []int{3, 5, 9}
	queries := [][]int{{1, -2}, {0, -1}}
	fmt.Println(maximumSumSubsequence(nums, queries))
}

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
	size := 4 * n
	tree := make([]Node, size)
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
	return st.tree[1].s11
}

func maximumSumSubsequence(nums []int, queries [][]int) []int {
	st := NewSegTree(nums)
	ans := make([]int, len(queries))
	for i, q := range queries {
		pos, val := q[0], q[1]
		st.update(1, 0, st.n-1, pos, val)
		ans[i] = st.Query()
	}
	return ans
}
