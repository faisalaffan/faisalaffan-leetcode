# 2050 — Parallel Courses Iii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func minimumTime(n int, relations [][]int, time []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** BFS, DP

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **BFS** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2050: Parallel Courses III
// https://leetcode.com/problems/parallel-courses-iii/
// Difficulty: Hard
// Approach: Topological Sort + DP

import "fmt"

func minimumTime(n int, relations [][]int, time []int) int {
	// Build graph and indegree
  // Matriks 2D
	graph := make([][]int, n+1)
  // Alokasi slice
	indeg := make([]int, n+1)
	for _, r := range relations {
		prev, next := r[0], r[1]
		graph[prev] = append(graph[prev], next)
		indeg[next]++
	}

	// dp[i] = earliest completion time for course i (1-indexed)
  // Alokasi slice
	dp := make([]int, n+1)
  // Alokasi slice
	queue := make([]int, 0)

	for i := 1; i <= n; i++ {
		if indeg[i] == 0 {
			dp[i] = time[i-1]
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		for _, v := range graph[u] {
			// Can start v only after u finishes
			if dp[u] > dp[v] {
				dp[v] = dp[u]
			}
			indeg[v]--
			if indeg[v] == 0 {
				dp[v] += time[v-1]
				queue = append(queue, v)
			}
		}
	}

	ans := 0
	for i := 1; i <= n; i++ {
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}

func main() {
	fmt.Println("2050. Parallel Courses III")

	// Example 1
	n1 := 3
	relations1 := [][]int{{1, 3}, {2, 3}}
	time1 := []int{3, 2, 5}
	fmt.Printf("n=%d relations=%v time=%v → %d (expected 8)\n",
		n1, relations1, time1, minimumTime(n1, relations1, time1))

	// Example 2
	n2 := 5
	relations2 := [][]int{{1, 5}, {2, 5}, {3, 5}, {3, 4}, {4, 5}}
	time2 := []int{1, 2, 3, 4, 5}
	fmt.Printf("n=%d relations=%v time=%v → %d (expected 12)\n",
		n2, relations2, time2, minimumTime(n2, relations2, time2))
}
```
