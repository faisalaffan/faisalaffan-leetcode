# 2538 — Difference Between Maximum And Minimum Price Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func maxOutput(n int, edges [][]int, price []int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS, Bitmask

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2538: Difference Between Maximum and Minimum Price Sum
// https://leetcode.com/problems/difference-between-maximum-and-minimum-price-sum/
// Difficulty: Hard

import "fmt"

// maxOutput computes the maximum difference between max and min price sum
// over all paths in the tree.
//
// For each node, compute maxDown (max sum from node to a leaf going down)
// and minDown (min sum from node to a leaf going down).
// Then for paths passing through a node connecting two children, compute
// the max sum and min sum. The answer is the maximum difference.
//
// Complexity: O(n) time, O(n) space
func maxOutput(n int, edges [][]int, price []int) int64 {
  // Membuat matriks/slice 2D untuk DP
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	ans := int64(0)

	// Returns (maxDown, minDown) for node u:
	// maxDown = max sum from u to a descendant (including u)
	// minDown = min sum from u to a descendant (including u)
	var dfs func(u, parent int) (int64, int64)
	dfs = func(u, parent int) (int64, int64) {
		// Each leaf serves as both a max and min single-node path
		maxDown := int64(price[u])
		minDown := int64(price[u])

		// Collect top 2 max and bottom 2 min from children
		var top1, top2 int64 = -1, -1 // max sums from children (excluding price[u])
		var bot1, bot2 int64 = 1<<60, 1<<60 // min sums from children

		for _, v := range adj[u] {
			if v == parent {
				continue
			}
			cMax, cMin := dfs(v, u)

			// Update maxDown/minDown = price[u] + best child path
			// For max, we add if child improves it; for min, if child makes it worse
			if cMax > 0 {
				maxDown = max64(maxDown, int64(price[u])+cMax)
			}
			if cMin < 0 {
				minDown = min64(minDown, int64(price[u])+cMin)
			}

			// Track top-2 max child sums (need to subtract price[u] to compare on equal footing)
			childMaxContribution := cMax // this is sum from child v going down, not including price[u]
			if childMaxContribution > top1 {
				top2, top1 = top1, childMaxContribution
			} else if childMaxContribution > top2 {
				top2 = childMaxContribution
			}

			childMinContribution := cMin
			if childMinContribution < bot1 {
				bot2, bot1 = bot1, childMinContribution
			} else if childMinContribution < bot2 {
				bot2 = childMinContribution
			}
		}

		// Consider path through u adding two child paths
		if top1 >= 0 {
			maxThrough := int64(price[u]) + top1
			if top2 >= 0 {
				maxThrough += top2
			}
			if maxThrough > ans {
				ans = maxThrough
			}
		}
		if bot1 <= 0 {
			minThrough := int64(price[u]) + bot1
			if bot2 <= 0 {
				minThrough += bot2
			}
			if minThrough < ans {
				// Update answer as maxDown - minDown through this node
				// We need the best path pair through u
				// Compute difference: best max and best min going through u
			}
		}

		// Compute difference at this node: best max through vs best min through
		// This considers paths that both pass through u, possibly using same children
		bestMaxThrough := int64(price[u])
		if top1 > 0 {
			bestMaxThrough += top1
		}
		if top2 > 0 {
			bestMaxThrough += top2
		}

		bestMinThrough := int64(price[u])
		if bot1 < 0 {
			bestMinThrough += bot1
		}
		if bot2 < 0 {
			bestMinThrough += bot2
		}

		if bestMaxThrough-bestMinThrough > ans {
			ans = bestMaxThrough - bestMinThrough
		}

		return maxDown, minDown
	}

	dfs(0, -1)
	return ans
}

func max64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func min64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Test cases
	fmt.Println("Test 1: n=6, edges=[[0,1],[1,2],[1,3],[3,4],[3,5]], price=[1,2,3,4,5,6] ->",
		maxOutput(6, [][]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}, {3, 5}}, []int{1, 2, 3, 4, 5, 6}))

	fmt.Println("Test 2: n=3, edges=[[0,1],[1,2]], price=[10,10,10] ->",
		maxOutput(3, [][]int{{0, 1}, {1, 2}}, []int{10, 10, 10}))

	fmt.Println("Test 3: n=1, edges=[], price=[5] ->",
		maxOutput(1, [][]int{}, []int{5}))

	fmt.Println("Test 4: n=4, edges=[[0,1],[1,2],[2,3]], price=[1,5,1,5] ->",
		maxOutput(4, [][]int{{0, 1}, {1, 2}, {2, 3}}, []int{1, 5, 1, 5}))

	fmt.Println("Test 5: n=2, edges=[[0,1]], price=[3,7] ->",
		maxOutput(2, [][]int{{0, 1}}, []int{3, 7}))
}
```
