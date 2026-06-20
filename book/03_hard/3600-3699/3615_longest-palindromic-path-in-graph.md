# 3615 — Longest Palindromic Path In Graph

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func longestPalindromicPath(n int, edges [][]int, label string) int
```

> **💡 Hint:** Since n <= 14, use DP over bitmask representing visited

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming, Bitmask

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3615: Longest Palindromic Path in Graph
// https://leetcode.com/problems/longest-palindromic-path-in-graph/
// Difficulty: Hard
//
// Find the longest path in a graph such that the sequence of node
// labels along the path forms a palindrome.
//
// Approach: Since n <= 14, use DP over bitmask representing visited
// nodes. dp[mask][i][j] = whether there's a palindromic path using
// nodes in mask starting at i ending at j.

import "fmt"

func main() {
	// Example 1
	fmt.Println(longestPalindromicPath(4, [][]int{{0, 1}, {1, 2}, {2, 3}, {0, 3}}, "abba"))
	// Example 2
	fmt.Println(longestPalindromicPath(3, [][]int{{0, 1}, {1, 2}}, "abc"))
	// Edge: single node
	fmt.Println(longestPalindromicPath(1, [][]int{}, "a"))
}

func longestPalindromicPath(n int, edges [][]int, label string) int {
  // Membuat matriks/slice 2D untuk DP
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// dp[mask][i][j] = mask includes i and j, path from i to j is palindrome
	totalMasks := 1 << n
  // Membuat matriks/slice 2D untuk DP
	dp := make([][][]bool, totalMasks)
	for mask := 0; mask < totalMasks; mask++ {
		dp[mask] = make([][]bool, n)
  // Range loop: iterasi dengan indeks + nilai
		for i := range dp[mask] {
			dp[mask][i] = make([]bool, n)
		}
	}

	ans := 0

	// Base: single node path
	for i := 0; i < n; i++ {
		mask := 1 << i
		dp[mask][i][i] = true
		if 1 > ans {
			ans = 1
		}
	}

	// Base: two-node edge
	for _, e := range edges {
		u, v := e[0], e[1]
		if label[u] == label[v] {
			mask := (1 << u) | (1 << v)
			dp[mask][u][v] = true
			dp[mask][v][u] = true
			if 2 > ans {
				ans = 2
			}
		}
	}

	// Extend paths
	for mask := 0; mask < totalMasks; mask++ {
		for i := 0; i < n; i++ {
			if mask&(1<<i) == 0 {
				continue
			}
			for j := 0; j < n; j++ {
				if mask&(1<<j) == 0 {
					continue
				}
				if !dp[mask][i][j] {
					continue
				}
				cnt := 0
				tmp := mask
				for tmp > 0 {
					cnt += tmp & 1
					tmp >>= 1
				}
				if cnt > ans {
					ans = cnt
				}

				// Try extending from i and j
				for _, ni := range adj[i] {
					if mask&(1<<ni) != 0 {
						continue
					}
					for _, nj := range adj[j] {
						if mask&(1<<nj) != 0 {
							continue
						}
						if ni != nj && label[ni] == label[nj] {
							nmask := mask | (1 << ni) | (1 << nj)
							dp[nmask][ni][nj] = true
							dp[nmask][nj][ni] = true
						} else if ni == nj && cnt > 0 {
							nmask := mask | (1 << ni)
							dp[nmask][ni][ni] = true
						}
					}
				}
			}
		}
	}

	// Single node extension (odd length palindrome with center expansion)
	for mask := 0; mask < totalMasks; mask++ {
		for i := 0; i < n; i++ {
			if mask&(1<<i) == 0 {
				continue
			}
			for j := 0; j < n; j++ {
				if !dp[mask][i][j] {
					continue
				}
				for _, ni := range adj[i] {
					if mask&(1<<ni) != 0 {
						continue
					}
					for _, nj := range adj[j] {
						if mask&(1<<nj) != 0 || ni != nj {
							continue
						}
						nmask := mask | (1 << ni)
						dp[nmask][ni][ni] = true
					}
				}
			}
		}
	}

	return ans
}
```
