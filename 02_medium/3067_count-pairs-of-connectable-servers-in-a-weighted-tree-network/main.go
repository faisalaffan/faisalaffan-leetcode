package main

// LeetCode #3067: Count Pairs of Connectable Servers in a Weighted Tree Network
// https://leetcode.com/problems/count-pairs-of-connectable-servers-in-a-weighted-tree-network/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(countPairsOfConnectableServers([][]int{{0, 1, 1}, {1, 2, 5}}, 1))
	fmt.Println(countPairsOfConnectableServers([][]int{{0, 6, 3}, {6, 5, 3}, {0, 3, 1}, {3, 2, 7}, {3, 1, 6}, {3, 4, 2}}, 3))
}

func countPairsOfConnectableServers(edges [][]int, signalSpeed int) []int {
	n := len(edges) + 1
	g := make([][][2]int, n)
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		g[u] = append(g[u], [2]int{v, w})
		g[v] = append(g[v], [2]int{u, w})
	}

	ans := make([]int, n)
	for i := 0; i < n; i++ {
		cnt := []int{}
		for _, ne := range g[i] {
			c := dfs(i, ne[0], ne[1], signalSpeed, g)
			if c > 0 {
				cnt = append(cnt, c)
			}
		}
		total := 0
		for j := 0; j < len(cnt); j++ {
			total += cnt[j]
		}
		pairs := 0
		for j := 0; j < len(cnt); j++ {
			total -= cnt[j]
			pairs += cnt[j] * total
		}
		ans[i] = pairs
	}
	return ans
}

func dfs(prev, curr, dist, signalSpeed int, g [][][2]int) int {
	c := 0
	if dist%signalSpeed == 0 {
		c++
	}
	for _, ne := range g[curr] {
		if ne[0] != prev {
			c += dfs(curr, ne[0], dist+ne[1], signalSpeed, g)
		}
	}
	return c
}
