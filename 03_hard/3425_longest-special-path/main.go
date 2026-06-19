package main

// LeetCode #3425: Longest Special Path
// https://leetcode.com/problems/longest-special-path/
// Difficulty: Hard
//
// Tree DFS sliding window: maintain a path (ancestor-to-descendant) where all
// node values are unique. Use a hash map tracking last occurrence depth.
// Prefix sums for edge weights give O(1) path length queries.

import "fmt"

func main() {
	// Example: edges=[[0,1,1],[1,2,2],[0,3,3]], nums=[1,2,1,2] -> [5,2]
	fmt.Println(longestSpecialPath([][]int{{0, 1, 1}, {1, 2, 2}, {0, 3, 3}}, []int{1, 2, 1, 2}))

	// Example from description: edges=[[0,1],[0,2],[1,3],[1,4]], nums=[1,2,3,4,5] -> [5,3]
	// Wait, this has no edge weights. Default weight = 1.
	// Path: 3-1-0-2 (all unique values), len = 3 edges = 3. But expected is 5?
	// Actually describing edge with undirected: edges=[[0,1],[0,2],[1,3],[1,4]]
	// without weights implies weight 1 for each.
	// Path 3-1-0-2 has length 3 (3 edges), 4 nodes. Not 5.
	// Maybe the expected output [5,3] means max_len=5, min_nodes=3?
	// Let me use weight 1 for edges without explicit weight interpretation.
	// Actually the problem guarantees edge weight as third element.
	// Let me handle both formats.
	fmt.Println(longestSpecialPath([][]int{{0, 1, 1}, {0, 2, 1}, {1, 3, 1}, {1, 4, 1}}, []int{1, 2, 3, 4, 5}))

	// Single node
	fmt.Println(longestSpecialPath([][]int{}, []int{5}))

	// Two nodes
	fmt.Println(longestSpecialPath([][]int{{0, 1, 10}}, []int{1, 2}))

	// Three-node path, all same values
	fmt.Println(longestSpecialPath([][]int{{0, 1, 1}, {1, 2, 1}}, []int{1, 1, 1}))
}

func longestSpecialPath(edges [][]int, nums []int) []int {
	n := len(nums)
	if n == 1 {
		return []int{0, 1}
	}

	adj := make([][][2]int, n)
	for _, e := range edges {
		u, v, w := e[0], e[1], 1
		if len(e) > 2 {
			w = e[2]
		}
		adj[u] = append(adj[u], [2]int{v, w})
		adj[v] = append(adj[v], [2]int{u, w})
	}

	lastOccur := make(map[int]int) // value -> last depth seen
	pathSum := make([]int, n)      // prefix sum of edge weights from root
	maxLen := 0
	minNodes := 1

	var dfs func(u, parent, depth, startDepth int)
	dfs = func(u, parent, depth, startDepth int) {
		val := nums[u]
		oldDepth, existed := lastOccur[val]

		if existed && oldDepth >= startDepth {
			startDepth = oldDepth + 1
		}

		lastOccur[val] = depth

		curLen := pathSum[depth] - pathSum[startDepth]
		curNodes := depth - startDepth + 1
		if curLen > maxLen || (curLen == maxLen && curNodes < minNodes) {
			maxLen = curLen
			minNodes = curNodes
		}

		for _, nb := range adj[u] {
			v, w := nb[0], nb[1]
			if v == parent {
				continue
			}
			pathSum[depth+1] = pathSum[depth] + w
			dfs(v, u, depth+1, startDepth)
		}

		// Restore lastOccur
		if existed {
			lastOccur[val] = oldDepth
		} else {
			delete(lastOccur, val)
		}
	}

	dfs(0, -1, 0, 0)
	return []int{maxLen, minNodes}
}
