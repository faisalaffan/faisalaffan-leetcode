# 3544 — Subtree Inversion Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan graf. Tugasmu menjelajahi atau menganalisis konektivitas graf.

**Cara berpikir:** Adjacency list `map[int][]int`. Gunakan BFS (queue) atau DFS (rekursif) dengan visited set untuk hindari siklus.

**Fungsi Solusi:** `func subtreeInversionSum(n int, edges [][]int, nums []int) int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DFS, Sorting

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **DFS** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3544: Subtree Inversion Sum
// https://leetcode.com/problems/subtree-inversion-sum/
// Difficulty: Hard
//
// For each node u, count inversions within its subtree (pairs (a,b) where a is
// visited before b in DFS and nums[a] > nums[b]). Sum over all nodes.
//
// DFS with sorted-list merging (small-to-large):
//   For each node, merge children's sorted value lists, counting cross-child
//   inversions. Add count of descendants with value < nums[u].

import (
	"fmt"
	"sort"
)

func subtreeInversionSum(n int, edges [][]int, nums []int) int64 {
  // Matriks 2D
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	var ans int64

	var dfs func(u, p int) ([]int, int64)
	dfs = func(u, p int) ([]int, int64) {
  // Alokasi slice
		vals := make([]int, 0)
		var invTotal int64

		for _, v := range adj[u] {
			if v == p {
				continue
			}
			childVals, childInv := dfs(v, u)
			invTotal += childInv

			// Count cross inversions:
			// Pair (a in earlier children, b in this child) where a > b
			// Since vals is sorted and contains values from earlier children:
			cross := int64(0)
			i := 0
			for _, b := range childVals {
				for i < len(vals) && vals[i] <= b {
					i++
				}
				cross += int64(len(vals) - i)
			}
			invTotal += cross

			// Small-to-large merge
			if len(vals) < len(childVals) {
				vals, childVals = childVals, vals
				// After swapping, vals = larger (was childVals), childVals = smaller (was vals)
				// The cross count was correct (computed before swap).
			}
  // Alokasi slice
			merged := make([]int, 0, len(vals)+len(childVals))
			p1, p2 := 0, 0
			for p1 < len(vals) && p2 < len(childVals) {
				if vals[p1] <= childVals[p2] {
					merged = append(merged, vals[p1])
					p1++
				} else {
					merged = append(merged, childVals[p2])
					p2++
				}
			}
			merged = append(merged, vals[p1:]...)
			merged = append(merged, childVals[p2:]...)
			vals = merged
		}

		// Count descendants with value < nums[u]
		// vals is sorted (from children only, does not include nums[u])
		less := sort.Search(len(vals), func(i int) bool { return vals[i] >= nums[u] })
		invTotal += int64(less)

		// Insert nums[u] into sorted vals for parent
		pos := sort.Search(len(vals), func(i int) bool { return vals[i] >= nums[u] })
		vals = append(vals, 0)
		copy(vals[pos+1:], vals[pos:])
		vals[pos] = nums[u]

		ans += invTotal
		return vals, invTotal
	}

	dfs(0, -1)
	return ans
}

// Brute-force verification
func subtreeInversionSumBrute(n int, edges [][]int, nums []int) int64 {
  // Matriks 2D
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// Get DFS preorder
	var order []int
	var dfsOrder func(u, p int)
	dfsOrder = func(u, p int) {
		order = append(order, u)
		for _, v := range adj[u] {
			if v != p {
				dfsOrder(v, u)
			}
		}
	}
	dfsOrder(0, -1)

	// tin/tout to identify subtree ranges
  // Alokasi slice
	tin := make([]int, n)
  // Alokasi slice
	tout := make([]int, n)
	var dfsRange func(u, p int, idx *int)
	dfsRange = func(u, p int, idx *int) {
		tin[u] = *idx
		(*idx)++
		for _, v := range adj[u] {
			if v != p {
				dfsRange(v, u, idx)
			}
		}
		tout[u] = *idx - 1
	}
	idx := 0
	dfsRange(0, -1, &idx)

	var total int64
	for u := 0; u < n; u++ {
		// Count inversions within subtree u
		var subtreeNodes []int
		for _, v := range order {
			if tin[v] >= tin[u] && tin[v] <= tout[u] {
				subtreeNodes = append(subtreeNodes, v)
			}
		}
		subInv := int64(0)
  // Linear scan O(n)
		for i := 0; i < len(subtreeNodes); i++ {
			for j := i + 1; j < len(subtreeNodes); j++ {
				if nums[subtreeNodes[i]] > nums[subtreeNodes[j]] {
					subInv++
				}
			}
		}
		total += subInv
	}
	return total
}

func main() {
	// Test: tree from user spec
	// Tree: 0(3)-1(2), 0-2(5), 1-3(1), 1-4(4)
	// DFS: 0,1,3,4,2
	n1 := 5
	edges1 := [][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}}
	nums1 := []int{3, 2, 5, 1, 4}
	res1 := subtreeInversionSum(n1, edges1, nums1)
	brute1 := subtreeInversionSumBrute(n1, edges1, nums1)
	fmt.Printf("subtree inversion sum -> %d (brute: %d)\n", res1, brute1)

	// Simple tests
	// Single node: 0 inversions
	fmt.Printf("single -> %d (expected 0)\n",
		subtreeInversionSum(1, [][]int{}, []int{5}))

	// Two nodes [2,1]: subtree 1 has 0, subtree 0 has (0,1):2>1 = 1. Total: 1.
	fmt.Printf("[2,1] -> %d (expected 1)\n",
		subtreeInversionSum(2, [][]int{{0, 1}}, []int{2, 1}))

	// Two nodes [1,2]: no inversion
	fmt.Printf("[1,2] -> %d (expected 0)\n",
		subtreeInversionSum(2, [][]int{{0, 1}}, []int{1, 2}))

	// Line 0-1-2, values [3,1,2]
	// DFS: 0,1,2
	// subtree 2: 0
	// subtree 1: (1,2): 1<2 no -> 0
	// subtree 0: (0,1):3>1, (0,2):3>2 -> 2
	// Total: 2
	fmt.Printf("line [3,1,2] -> %d (expected 2)\n",
		subtreeInversionSum(3, [][]int{{0, 1}, {1, 2}}, []int{3, 1, 2}))
}
```
