package main

// LeetCode #3425: Longest Special Path
// https://leetcode.com/problems/longest-special-path/
// Difficulty: Hard
//
// Tree DFS with lastOccurrence map tracking value depths.
// Backtrack on return to maintain correct state.

import "fmt"

func main() {
	fmt.Println(LongestSpecialPath([][]int{{0, 1, 1}, {1, 2, 2}, {0, 3, 3}}, []int{1, 2, 1, 2}))
}

func LongestSpecialPath(edges [][]int, nums []int) []int {
	n := len(nums)
	adj := make([][][2]int, n)
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		adj[u] = append(adj[u], [2]int{v, w})
		adj[v] = append(adj[v], [2]int{u, w})
	}

	lastOccur := make(map[int]int)
	pathLen := 0
	maxLen := 0
	minNodes := 0

	var dfs func(u, parent, depth, startDepth int)
	dfs = func(u, parent, depth, startDepth int) {
		val := nums[u]

		oldDepth, existed := lastOccur[val]
		if existed && oldDepth >= startDepth {
			startDepth = oldDepth + 1
			// Recalculate pathLen from startDepth to current depth
			// pathLen represents total edge weight from root to current
			// We need the path length from startDepth to depth
		}

		lastOccur[val] = depth

		// Check if current path is longest
		for _, nb := range adj[u] {
			v, w := nb[0], nb[1]
			if v == parent {
				continue
			}
			newPathLen := pathLen + w
			oldPathLen := pathLen
			pathLen = newPathLen

			dfs(v, u, depth+1, startDepth)

			// Backtrack pathLen
			pathLen = oldPathLen
		}

		// Now compute valid path from this node upward
		// A valid path ends at this node and starts at startDepth
		nodeCount := depth - startDepth + 1
		if pathLen > maxLen || (pathLen == maxLen && nodeCount < minNodes) {
			maxLen = pathLen
			minNodes = nodeCount
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
