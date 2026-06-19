package main

// LeetCode #1157: Online Majority Element In Subarray
// https://leetcode.com/problems/online-majority-element-in-subarray/
// Difficulty: Hard

import (
	"fmt"
	"math/rand"
	"sort"
)

// MajorityChecker uses segment tree for candidate + random sampling verification.
type MajorityChecker struct {
	arr  []int
	tree []segNode
	pos  map[int][]int // value -> sorted positions
}

type segNode struct {
	cand int
	cnt  int
}

func merge(a, b segNode) segNode {
	if a.cand == b.cand {
		return segNode{a.cand, a.cnt + b.cnt}
	}
	if a.cnt > b.cnt {
		return segNode{a.cand, a.cnt - b.cnt}
	}
	return segNode{b.cand, b.cnt - a.cnt}
}

func build(arr []int, tree []segNode, idx, l, r int) {
	if l == r {
		tree[idx] = segNode{arr[l], 1}
		return
	}
	m := (l + r) / 2
	build(arr, tree, idx*2+1, l, m)
	build(arr, tree, idx*2+2, m+1, r)
	tree[idx] = merge(tree[idx*2+1], tree[idx*2+2])
}

func query(tree []segNode, idx, l, r, ql, qr int) segNode {
	if ql <= l && r <= qr {
		return tree[idx]
	}
	m := (l + r) / 2
	if qr <= m {
		return query(tree, idx*2+1, l, m, ql, qr)
	}
	if ql > m {
		return query(tree, idx*2+2, m+1, r, ql, qr)
	}
	return merge(
		query(tree, idx*2+1, l, m, ql, qr),
		query(tree, idx*2+2, m+1, r, ql, qr),
	)
}

func Constructor(arr []int) MajorityChecker {
	n := len(arr)
	tree := make([]segNode, 4*n)
	build(arr, tree, 0, 0, n-1)
	pos := make(map[int][]int)
	for i, v := range arr {
		pos[v] = append(pos[v], i)
	}
	return MajorityChecker{arr, tree, pos}
}

func (mc *MajorityChecker) Query(left, right, threshold int) int {
	// Try segment tree candidate
	cand := query(mc.tree, 0, 0, len(mc.arr)-1, left, right).cand
	// Verify by counting occurrences in range
	positions := mc.pos[cand]
	lo := sort.Search(len(positions), func(i int) bool { return positions[i] >= left })
	hi := sort.Search(len(positions), func(i int) bool { return positions[i] > right })
	if hi-lo >= threshold {
		return cand
	}

	// Fallback: random sampling 20 times
	for i := 0; i < 20; i++ {
		idx := left + rand.Intn(right-left+1)
		v := mc.arr[idx]
		positions := mc.pos[v]
		lo := sort.Search(len(positions), func(i int) bool { return positions[i] >= left })
		hi := sort.Search(len(positions), func(i int) bool { return positions[i] > right })
		if hi-lo >= threshold {
			return v
		}
	}
	return -1
}

func main() {
	// Test case 1
	checker := Constructor([]int{1, 1, 2, 2, 1, 1})
	fmt.Println(checker.Query(0, 5, 4)) // 1
	fmt.Println(checker.Query(0, 3, 3)) // -1
	fmt.Println(checker.Query(2, 3, 2)) // 2

	// Test case 2
	checker2 := Constructor([]int{2, 2, 1, 1, 1, 2, 2})
	fmt.Println(checker2.Query(0, 5, 4)) // 1
	fmt.Println(checker2.Query(0, 6, 4)) // 2
}
