package main

// LeetCode #2213: Longest Substring of One Repeating Character
// https://leetcode.com/problems/longest-substring-of-one-repeating-character/
// Difficulty: Hard
//
// Segment tree: each node stores the longest run of a single character in its
// interval, plus prefix run length, suffix run length, and the characters at
// left and right boundaries. Updates at a single position; query returns the
// tree root's max.

import (
	"fmt"
)

func main() {
	// s = "babacc", queryCharacters = "bcb", queryIndices = [1,3,3]
	// Expected: [1, 3, 3]
	s := "babacc"
	queryCharacters := "bcb"
	queryIndices := []int{1, 3, 3}
	fmt.Println(longestRepeating(s, queryCharacters, queryIndices))

	// Single char
	fmt.Println(longestRepeating("a", "b", []int{0}))

	// Already repeating
	fmt.Println(longestRepeating("aaa", "a", []int{0}))
}

type SegNode struct {
	prefLen   int
	suffLen   int
	maxLen    int
	leftChar  byte
	rightChar byte
	size      int
}

type SegTree struct {
	tree []SegNode
	n    int
}

func merge(left, right SegNode) SegNode {
	var res SegNode
	res.leftChar = left.leftChar
	res.rightChar = right.rightChar
	res.size = left.size + right.size

	// Prefix length
	res.prefLen = left.prefLen
	if left.prefLen == left.size && left.rightChar == right.leftChar {
		res.prefLen = left.size + right.prefLen
	}

	// Suffix length
	res.suffLen = right.suffLen
	if right.suffLen == right.size && right.leftChar == left.rightChar {
		res.suffLen = right.size + left.suffLen
	}

	// Max length
	res.maxLen = max(left.maxLen, right.maxLen)
	if left.rightChar == right.leftChar {
		res.maxLen = max(res.maxLen, left.suffLen+right.prefLen)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func NewSegTree(s string) *SegTree {
	n := len(s)
	tree := make([]SegNode, 4*n)
	st := &SegTree{tree: tree, n: n}
	st.build(1, 1, n, s)
	return st
}

func (st *SegTree) build(idx, l, r int, s string) {
	if l == r {
		st.tree[idx] = SegNode{
			prefLen:   1,
			suffLen:   1,
			maxLen:    1,
			leftChar:  s[l-1],
			rightChar: s[l-1],
			size:      1,
		}
		return
	}
	mid := (l + r) / 2
	st.build(idx*2, l, mid, s)
	st.build(idx*2+1, mid+1, r, s)
	st.tree[idx] = merge(st.tree[idx*2], st.tree[idx*2+1])
}

func (st *SegTree) Update(pos int, c byte) {
	st.update(1, 1, st.n, pos, c)
}

func (st *SegTree) update(idx, l, r, pos int, c byte) {
	if l == r {
		st.tree[idx].leftChar = c
		st.tree[idx].rightChar = c
		return
	}
	mid := (l + r) / 2
	if pos <= mid {
		st.update(idx*2, l, mid, pos, c)
	} else {
		st.update(idx*2+1, mid+1, r, pos, c)
	}
	st.tree[idx] = merge(st.tree[idx*2], st.tree[idx*2+1])
}

func longestRepeating(s string, queryCharacters string, queryIndices []int) []int {
	st := NewSegTree(s)
	ans := make([]int, len(queryCharacters))
	for i, c := range queryCharacters {
		st.Update(queryIndices[i]+1, byte(c))
		ans[i] = st.tree[1].maxLen
	}
	return ans
}
