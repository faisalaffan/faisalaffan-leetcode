package main

// LeetCode #3949: Subtree Inversion Sum II
// https://leetcode.com/problems/subtree-inversion-sum-ii/
// Difficulty: Hard [Paid]
//
// Given a rooted tree with values, count inversion pairs within
// each subtree where parent-child values are inverted. Sum the
// inversion counts across all valid selections.
//
// Approach: DFS with Fenwick tree for inversion counting. For
// each node, count inversions between current node and its
// descendants using value frequency BIT.

import "fmt"

func main() {
	// Example 1
	fmt.Println(maxSum([][]int{{0, 1}, {1, 2}}, []int{3, 1, 2}, 2))
	// Example 2
	fmt.Println(maxSum([][]int{{0, 1}, {0, 2}}, []int{5, 3, 4}, 1))
	// Edge: single node
	fmt.Println(maxSum([][]int{}, []int{7}, 1))
}

const MAX_VAL = 100000

type BIT struct {
	tree []int
}

func NewBIT(size int) *BIT {
	return &BIT{tree: make([]int, size+2)}
}

func (b *BIT) Update(idx, delta int) {
	idx++
	for idx < len(b.tree) {
		b.tree[idx] += delta
		idx += idx & -idx
	}
}

func (b *BIT) Query(idx int) int {
	idx++
	sum := 0
	for idx > 0 {
		sum += b.tree[idx]
		idx -= idx & -idx
	}
	return sum
}

func (b *BIT) Range(l, r int) int {
	if l > r {
		return 0
	}
	return b.Query(r) - b.Query(l-1)
}

func maxSum(edges [][]int, nums []int, k int) int {
	n := len(nums)
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	totalSum := 0
	bit := NewBIT(MAX_VAL)

	var dfs func(u, p int)
	dfs = func(u, p int) {
		// Count existing values greater than nums[u]
		greater := bit.Range(nums[u]+1, MAX_VAL)
		totalSum += greater

		bit.Update(nums[u], 1)

		for _, v := range adj[u] {
			if v == p {
				continue
			}
			dfs(v, u)
		}

		bit.Update(nums[u], -1)
	}

	dfs(0, -1)
	return totalSum
}
