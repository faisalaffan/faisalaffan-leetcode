package main

import (
	"fmt"
)

// LeetCode #1273: Delete Tree Nodes
// https://leetcode.com/problems/delete-tree-nodes/
// Difficulty: Medium [Paid]

// Delete subtrees where sum of node values = 0.
// Return number of remaining nodes.

// Time: O(n)
// Space: O(n)

func deleteTreeNodes(nodes int, parent []int, value []int) int {
	children := make([][]int, nodes)
	for i := 1; i < nodes; i++ {
		p := parent[i]
		children[p] = append(children[p], i)
	}

	var dfs func(node int) (int, int) // returns (sum, count)
	dfs = func(node int) (int, int) {
		sum := value[node]
		count := 1
		for _, child := range children[node] {
			childSum, childCount := dfs(child)
			sum += childSum
			count += childCount
		}
		if sum == 0 {
			return 0, 0
		}
		return sum, count
	}

	_, cnt := dfs(0)
	return cnt
}

func main() {
	fmt.Printf("%d (expected: 2)\n",
		deleteTreeNodes(7, []int{-1, 0, 0, 1, 2, 2, 2}, []int{1, -2, 4, 0, -2, -1, 1}))

	fmt.Printf("%d (expected: 4)\n",
		deleteTreeNodes(4, []int{-1, 0, 0, 1}, []int{1, 2, -1, -2}))
}
