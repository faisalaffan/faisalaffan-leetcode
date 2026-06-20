# 3910 — Count Connected Subgraphs With Even Node Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan graf. Tugasmu menjelajahi atau menganalisis konektivitas graf.

**Cara berpikir:** Adjacency list `map[int][]int`. Gunakan BFS (queue) atau DFS (rekursif) dengan visited set untuk hindari siklus.

**Fungsi Solusi:** `func evenSumSubgraphs(nums []int, edges [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3910: Count Connected Subgraphs with Even Node Sum
// https://leetcode.com/problems/count-connected-subgraphs-with-even-node-sum/
// Difficulty: Hard
//
// Count connected subgraphs where the sum of node values is even.
// Graph is a tree with n nodes and n-1 edges.
//
// Approach: Tree DP. For each node, compute count of connected
// subgraphs in its subtree with even and odd sums. Track whether
// the subgraph includes the node (for connectivity).

import "fmt"

func main() {
	// Example 1
	fmt.Println(evenSumSubgraphs([]int{1, 2, 3, 4}, [][]int{{0, 1}, {1, 2}, {2, 3}}))
	// Example 2
	fmt.Println(evenSumSubgraphs([]int{1, 1, 1}, [][]int{{0, 1}, {1, 2}}))
	// Edge: single node
	fmt.Println(evenSumSubgraphs([]int{2}, [][]int{}))
}

const MOD = 1000000007

func evenSumSubgraphs(nums []int, edges [][]int) int {
	n := len(nums)
  // Edge case: input kosong
	if n == 0 {
		return 0
	}

  // Matriks 2D
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// DFS returns (even, odd, inclEven, inclOdd)
	// where inclEven/Odd counts connected subgraphs that include the node
	var dfs func(u, p int) (int, int, int, int)
	dfs = func(u, p int) (int, int, int, int) {
		even, odd := 0, 0
		inclEven, inclOdd := 0, 0

		if nums[u]%2 == 0 {
			inclEven = 1
		} else {
			inclOdd = 1
		}

		for _, v := range adj[u] {
			if v == p {
				continue
			}
			ce, co, ie, io := dfs(v, u)

			// Combine
			newEven := (even + ce) % MOD
			newOdd := (odd + co) % MOD

			// Include node -> child can attach or not
			// New including-this-node subgraphs:
			nei := inclEven
			noi := inclOdd

			// Add child attached: parity combines
			// (inclEven * ie) -> even + even = even
			// (inclEven * io) -> even + odd = odd
			// (inclOdd * ie) -> odd + even = odd
			// (inclOdd * io) -> odd + odd = even
			nei = (nei + (inclEven * ie % MOD) + (inclOdd * io % MOD)) % MOD
			noi = (noi + (inclEven * io % MOD) + (inclOdd * ie % MOD)) % MOD

			even = newEven
			odd = newOdd
			inclEven = nei
			inclOdd = noi
		}

		even = (even + inclEven) % MOD
		odd = (odd + inclOdd) % MOD

		return even, odd, inclEven, inclOdd
	}

	even, _, _, _ := dfs(0, -1)
	return even
}
```
