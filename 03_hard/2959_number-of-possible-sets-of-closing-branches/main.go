package main

// LeetCode #2959: Number of Possible Sets of Closing Branches
// https://leetcode.com/problems/number-of-possible-sets-of-closing-branches/
// Difficulty: Hard

import "fmt"

func numberOfSets(n int, maxDistance int, roads [][]int) int {
	const inf = 1 << 29
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, n)
		for j := range g[i] {
			g[i][j] = inf
		}
		g[i][i] = 0
	}
	for _, r := range roads {
		u, v, w := r[0], r[1], r[2]
		if w < g[u][v] {
			g[u][v] = w
			g[v][u] = w
		}
	}
	ans := 0
	for mask := 0; mask < (1 << n); mask++ {
		dist := make([][]int, n)
		for i := range dist {
			dist[i] = make([]int, n)
			copy(dist[i], g[i])
		}
		for k := 0; k < n; k++ {
			if mask>>k&1 == 0 {
				continue
			}
			for i := 0; i < n; i++ {
				if mask>>i&1 == 0 || dist[i][k] == inf {
					continue
				}
				for j := 0; j < n; j++ {
					if mask>>j&1 == 0 {
						continue
					}
					if nd := dist[i][k] + dist[k][j]; nd < dist[i][j] {
						dist[i][j] = nd
					}
				}
			}
		}
		ok := 1
		for i := 0; i < n && ok == 1; i++ {
			if mask>>i&1 == 0 {
				continue
			}
			for j := i + 1; j < n; j++ {
				if mask>>j&1 == 0 {
					continue
				}
				if dist[i][j] > maxDistance {
					ok = 0
					break
				}
			}
		}
		ans += ok
	}
	return ans
}

func main() {
	fmt.Println(numberOfSets(3, 5, [][]int{{0, 1, 2}, {1, 2, 10}, {0, 2, 10}}))
	fmt.Println(numberOfSets(3, 5, [][]int{{0, 1, 20}, {0, 2, 5}, {1, 2, 2}}))
}
