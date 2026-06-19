package main

// LeetCode #3544: Subtree Inversion Sum
// https://leetcode.com/problems/subtree-inversion-sum/
// Difficulty: Hard
//
// For each node, count inversions (pairs (a,b) where a appears before b in DFS
// preorder and nums[a] > nums[b]) within its subtree. Sum over all nodes.
//
// Use DFS with small-to-large merging of sorted value lists.
// For each node u:
//   subtreeInv[u] = sum(childInv) + crossSubtreeInversions + countLess(nums[u])
// where crossSubtreeInversions counts pairs (a in earlier child, b in later child)
// with nums[a] > nums[b], and countLess counts descendants with value < nums[u].

import (
	"fmt"
	"sort"
)

func subtreeInversionSum(n int, edges [][]int, nums []int) int64 {
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	var ans int64

	var dfs func(u, p int) ([]int, int64)
	dfs = func(u, p int) ([]int, int64) {
		vals := []int{nums[u]}
		var invTotal int64

		for _, v := range adj[u] {
			if v == p {
				continue
			}
			childVals, childInv := dfs(v, u)
			invTotal += childInv

			// Count cross inversions between vals (earlier children) and childVals (this child)
			// Pair (a in vals, b in childVals) where a > b and a visited before b
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
			}
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
		// vals is sorted; binary search for first element >= nums[u]
		less := sort.Search(len(vals), func(i int) bool { return vals[i] >= nums[u] })
		invTotal += int64(less)

		ans += invTotal
		return vals, invTotal
	}

	dfs(0, -1)
	return ans
}

func main() {
	// Test: n=5, edges=[[0,1],[0,2],[1,3],[1,4]], nums=[3,2,5,1,4]
	// Tree:
	//     0(3)
	//    / \
	//  1(2) 2(5)
	//  / \
	// 3(1) 4(4)
	// DFS preorder: 0,1,3,4,2
	res := subtreeInversionSum(5,
		[][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}},
		[]int{3, 2, 5, 1, 4})
	fmt.Printf("subtree inversion sum -> %d (expected 10)\n", res)

	// Test: single node
	res2 := subtreeInversionSum(1, [][]int{}, []int{5})
	fmt.Printf("single node -> %d (expected 0)\n", res2)

	// Test: two nodes, values [2,1]
	// Tree: 0(2)-1(1)
	// Subtrees:
	//   subtree_1: {1} -> 0 inversions
	//   subtree_0: {0,1} -> (0,1) where 2>1 -> 1 inversion
	// Total: 0+1 = 1
	res3 := subtreeInversionSum(2, [][]int{{0, 1}}, []int{2, 1})
	fmt.Printf("two nodes [2,1] -> %d (expected 1)\n", res3)

	// Test: two nodes, values [1,2] (no inversion)
	res4 := subtreeInversionSum(2, [][]int{{0, 1}}, []int{1, 2})
	fmt.Printf("two nodes [1,2] -> %d (expected 0)\n", res4)

	// Test: line of 3, values [3,1,2]
	// Tree: 0(3)-1(1)-2(2)
	// DFS: 0,1,2
	// subtrees:
	//   subtree_2: {2} -> 0
	//   subtree_1: {1,2} -> (1,2): 1<2 no -> 0
	//   subtree_0: {0,1,2} -> (0,1):3>1, (0,2):3>2 -> 2
	// Total: 0+0+2 = 2
	res5 := subtreeInversionSum(3, [][]int{{0, 1}, {1, 2}}, []int{3, 1, 2})
	fmt.Printf("line [3,1,2] -> %d (expected 2)\n", res5)
}
