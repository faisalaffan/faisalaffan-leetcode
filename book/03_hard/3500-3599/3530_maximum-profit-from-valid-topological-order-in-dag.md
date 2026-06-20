# 3530 — Maximum Profit From Valid Topological Order In Dag

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxProfit(n int, edges [][]int, score []int) int
```

> **💡 Hint:** DP over subsets (bitmask DP). For each mask, try adding any

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming, Topological Sort, Bitmask

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3530: Maximum Profit from Valid Topological Order in DAG
// https://leetcode.com/problems/maximum-profit-from-valid-topological-order-in-dag/
// Difficulty: Hard
//
// Given a DAG with n nodes (0..n-1), edges, and node scores, find a valid
// topological order that maximizes the profit. Profit is defined as the sum
// of scores of nodes that appear at certain positions.
//
// Approach: DP over subsets (bitmask DP). For each mask, try adding any
// node whose prerequisites are satisfied.

import "fmt"

func main() {
	// Example 1
	fmt.Println(maxProfit(3, [][]int{{0, 1}, {1, 2}}, []int{1, 2, 3}))
	// Example 2
	fmt.Println(maxProfit(2, [][]int{{0, 1}}, []int{5, 3}))
	// Example 3: no edges
	fmt.Println(maxProfit(2, [][]int{}, []int{10, 20}))
	// Edge: single node
	fmt.Println(maxProfit(1, [][]int{}, []int{7}))
}

func maxProfit(n int, edges [][]int, score []int) int {
	// Build adjacency and indegree
  // Membuat matriks/slice 2D untuk DP
	adj := make([][]int, n)
  // Alokasi slice integer
	inDegree := make([]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		inDegree[v]++
	}

	// DP over masks
	totalMasks := 1 << n
  // Alokasi slice integer
	dp := make([]int, totalMasks)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dp {
		dp[i] = -1 << 30
	}
	dp[0] = 0

	// Precompute indegree contributions for each mask
	for mask := 0; mask < totalMasks; mask++ {
		if dp[mask] < 0 {
			continue
		}
		// Count how many nodes are already placed
		pos := 0
		for i := 0; i < n; i++ {
			if mask&(1<<i) != 0 {
				pos++
			}
		}

		// Track current indegree state
  // Alokasi slice integer
		curDegree := make([]int, n)
		copy(curDegree, inDegree)

		// Subtract edges from already placed nodes
		for u := 0; u < n; u++ {
			if mask&(1<<u) != 0 {
				for _, v := range adj[u] {
					curDegree[v]--
				}
			}
		}

		// Try adding any node with indegree 0 not yet placed
		for v := 0; v < n; v++ {
			if mask&(1<<v) == 0 && curDegree[v] == 0 {
				newMask := mask | (1 << v)
				profit := score[v] // profit for being at this position
				val := dp[mask] + profit
				if val > dp[newMask] {
					dp[newMask] = val
				}
			}
		}
	}

	return dp[totalMasks-1]
}
```
