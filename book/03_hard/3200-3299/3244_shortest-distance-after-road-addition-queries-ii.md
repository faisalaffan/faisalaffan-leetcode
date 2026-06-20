# 3244 — Shortest Distance After Road Addition Queries Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func shortestDistanceAfterQueries(n int, queries [][]int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3244: Shortest Distance After Road Addition Queries II
// https://leetcode.com/problems/shortest-distance-after-road-addition-queries-ii/
// Difficulty: Hard
//
// Maintain a "next" pointer array. Initially next[i] = i+1.
// When adding road (a,b), if next[a] < b, skip all intermediate nodes
// between a and b that are still active, decrementing the distance for each.
// Each node is skipped at most once, so total O(n) across all queries.

import "fmt"

func main() {
	// Example 1: n=5, queries=[[2,4],[0,2],[0,4]] => [3,2,1]
	fmt.Println(shortestDistanceAfterQueries(5, [][]int{{2, 4}, {0, 2}, {0, 4}}))
	// Example 2: single query
	fmt.Println(shortestDistanceAfterQueries(4, [][]int{{0, 3}}))
	// Example 3: no-op query
	fmt.Println(shortestDistanceAfterQueries(4, [][]int{{1, 2}}))
	// Example 4: multiple queries same range
	fmt.Println(shortestDistanceAfterQueries(6, [][]int{{1, 4}, {2, 5}, {0, 5}}))
	// Example 5: n=2
	fmt.Println(shortestDistanceAfterQueries(2, [][]int{{0, 1}}))
}

func shortestDistanceAfterQueries(n int, queries [][]int) []int {
	// next[i] = the next active node reachable from i (i+1 initially)
  // Alokasi slice integer
	next := make([]int, n)
	for i := 0; i < n-1; i++ {
		next[i] = i + 1
	}
	next[n-1] = n - 1 // sentinel

	dist := n - 1 // initial path length (0->1->2->...->n-1)
  // Alokasi slice integer
	ans := make([]int, len(queries))

	for qi, q := range queries {
		a, b := q[0], q[1]
		if next[a] >= b {
			// This query doesn't improve the path
			ans[qi] = dist
			continue
		}

		// Skip all active nodes between a and b
		// start from the node after a
		i := next[a]
		for i < b {
			// Skip this node: connect its predecessor directly to b
			next[i], i = b, next[i]
			dist--
		}
		next[a] = b
		ans[qi] = dist
	}

	return ans
}
```
