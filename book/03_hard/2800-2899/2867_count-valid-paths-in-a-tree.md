# 2867 — Count Valid Paths In A Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func countValidPathsInATree(n int, edges [][]int, values []int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS, Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2867: Count Valid Paths in a Tree
// https://leetcode.com/problems/count-valid-paths-in-a-tree/
// Difficulty: Hard
//
// Count paths in an undirected tree where exactly one node has a prime value.
// Tree DP approach: root the tree at 0. For each node, compute dp0 = number
// of downward paths (node to descendant) with 0 primes, dp1 = number of
// downward paths with exactly 1 prime. Count through-paths that pass through
// a node combining two child subtrees. Total = sum(dp1[node]) + sum(through paths).

import "fmt"

const maxN = 100000

func countValidPathsInATree(n int, edges [][]int, values []int) int64 {
	// Sieve primes up to max value in values
	maxVal := 0
	for _, v := range values {
		if v > maxVal {
			maxVal = v
		}
	}
	isPrime := make([]bool, maxVal+1)
	for i := 2; i <= maxVal; i++ {
		isPrime[i] = true
	}
	for i := 2; i*i <= maxVal; i++ {
		if isPrime[i] {
			for j := i * i; j <= maxVal; j += i {
				isPrime[j] = false
			}
		}
	}

	// Build adjacency list
  // Membuat matriks/slice 2D untuk DP
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	var ans int64 = 0

	var dfs func(u, parent int) (int64, int64)
	dfs = func(u, parent int) (int64, int64) {
		// dp0 = paths starting at u going down with 0 primes
		// dp1 = paths starting at u going down with exactly 1 prime
		var dp0, dp1 int64 = 0, 0
		prime := isPrime[values[u]]

		if prime {
			dp1 = 1 // just the node itself
		} else {
			dp0 = 1 // just the node itself
		}

		// We need to handle through-paths as we iterate children
		// By maintaining running sums of child dp0 and dp1
		var sumChild0, sumChild1 int64 = 0, 0

		for _, v := range adj[u] {
			if v == parent {
				continue
			}
			child0, child1 := dfs(v, u)

			// Count through-paths passing through u from two different child subtrees
			if prime {
				// Node is prime: combine 0-prime paths from different children
				// f0[a] * f0[b] through prime node => exactly 1 prime total
				ans += sumChild0 * child0
			} else {
				// Node is not prime: combine 0-prime from one child with 1-prime from another
				ans += sumChild0*child1 + sumChild1*child0
			}

			sumChild0 += child0
			sumChild1 += child1

			// Extend child's paths upward to u
			if prime {
				// Extending 0-prime paths: now they have 1 prime (node u)
				dp1 += child0
				// Extending 1-prime paths: now they have 2 primes => invalid
				// child1 paths are not extended through prime node
			} else {
				dp0 += child0
				dp1 += child1
			}
		}

		// Count dp1 paths (downward with 1 prime starting at u) as valid paths
		ans += dp1

		return dp0, dp1
	}

	dfs(0, -1)
	return ans
}

func main() {
	// Example 1: n=5, edges=[[0,1],[0,2],[1,3],[1,4]], values=[2,3,1,4,5]
	fmt.Println(countValidPathsInATree(5,
		[][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}},
		[]int{2, 3, 1, 4, 5}))

	// Example 2: n=1, single node, value=2
	fmt.Println(countValidPathsInATree(1, [][]int{}, []int{2}))

	// Example 3: n=2, edge=[0,1], values=[2,4] -> only [0] valid (prime=2)
	fmt.Println(countValidPathsInATree(2,
		[][]int{{0, 1}},
		[]int{2, 4}))

	// Example 4: line of 3, prime at middle
	fmt.Println(countValidPathsInATree(3,
		[][]int{{0, 1}, {1, 2}},
		[]int{4, 2, 6}))
	// Paths: [1], [0,1], [1,2], [0,1,2] = 4

	// Example 5: star with center prime
	fmt.Println(countValidPathsInATree(4,
		[][]int{{0, 1}, {0, 2}, {0, 3}},
		[]int{2, 4, 6, 8}))
}
```
