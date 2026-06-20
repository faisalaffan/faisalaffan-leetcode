# 1857 — Largest Color Value In A Directed Graph

## Deskripsi

**Soal:** [1857. Largest Color Value In A Directed Graph](https://leetcode.com/problems/largest-color-value-in-a-directed-graph/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Dynamic Programming (DP), Topological Sort (pengurutan topologi), LIS (Longest Increasing Subsequence)

> **Ide Kunci:** Topological Sort + DP per Color.

## Solusi Go

```go
package main

// LeetCode #1857: Largest Color Value in a Directed Graph
// https://leetcode.com/problems/largest-color-value-in-a-directed-graph/
// Difficulty: Hard
//
// Approach: Topological Sort + DP per Color.
//   For each node, maintain a dp[c] = max count of color c on any path
//   ending at this node. Process nodes in topological (Kahn's) order.
//   For each edge u->v:
//     dp[v][c] = max(dp[v][c], dp[u][c] + (1 if colors[v]==c else 0))
//   The answer is the max dp value across all nodes and colors.
//   If there's a cycle (not all nodes processed), return -1.

import "fmt"

func main() {
	// Example 1
	fmt.Println("Example 1:",
		largestPathValue("abaca", [][]int{{0, 1}, {0, 2}, {2, 3}, {3, 4}}))
	// Expected: 3 (path 0->2->3->4 has 3 'a's)

	// Example 2: cycle
	fmt.Println("Example 2:",
		largestPathValue("a", [][]int{{0, 0}}))
	// Expected: -1 (self-loop is a cycle)

	// Example 3: no edges
	fmt.Println("Example 3:",
		largestPathValue("abc", [][]int{}))
	// Expected: 1 (each node alone has count 1 of its own color)

	// Edge case: single node
	fmt.Println("Edge (single):",
		largestPathValue("z", [][]int{}))
	// Expected: 1

	// Two nodes, two colors
	fmt.Println("Edge (two nodes):",
		largestPathValue("ab", [][]int{{0, 1}}))
	// Expected: 1 (path 0->1 has one a and one b, max frequency = 1)
}

func largestPathValue(colors string, edges [][]int) int {
	n := len(colors)

	// Build adjacency list and in-degree array
  // Membuat slice 2D untuk DP/tabel
	adj := make([][]int, n)
  // Membuat slice untuk menyimpan hasil
	inDeg := make([]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		inDeg[v]++
	}

	// dp[i][c] = max count of color c on any path ending at node i
  // Membuat slice untuk menyimpan hasil
	dp := make([][26]int, n)

	// Kahn's topological sort
  // Membuat slice untuk menyimpan hasil
	queue := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if inDeg[i] == 0 {
			queue = append(queue, i)
			dp[i][colors[i]-'a'] = 1
		}
	}

	processed := 0
	ans := 0

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		processed++

		// Update global answer with dp[u]
		for c := 0; c < 26; c++ {
			if dp[u][c] > ans {
				ans = dp[u][c]
			}
		}

		for _, v := range adj[u] {
			// Propagate dp from u to v
			for c := 0; c < 26; c++ {
				add := 0
				if int(colors[v]-'a') == c {
					add = 1
				}
				if dp[u][c]+add > dp[v][c] {
					dp[v][c] = dp[u][c] + add
				}
			}
			inDeg[v]--
			if inDeg[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	if processed != n {
		return -1 // cycle detected
	}

	return ans
}

// Stub kept for compatibility with the repo scaffold.
func LargestColorValueInADirectedGraph() any {
	return largestPathValue("abaca", [][]int{{0, 1}, {0, 2}, {2, 3}, {3, 4}})
}
```
