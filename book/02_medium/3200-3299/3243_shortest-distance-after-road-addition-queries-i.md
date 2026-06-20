# 3243 — Shortest Distance After Road Addition Queries I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func shortestDistanceAfterQueries(n int, queries [][]int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** BFS

**Kompleksitas Waktu:** O(q * (n + q))  
**Kompleksitas Ruang:** O(n + q)

> **Untuk fresh graduate:** Kuasai dulu teknik **BFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3243: Shortest Distance After Road Addition Queries I
// https://leetcode.com/problems/shortest-distance-after-road-addition-queries-i/
// Difficulty: Medium
// Time: O(q * (n + q)) | Space: O(n + q)

import (
	"fmt"
	"math"
)

func shortestDistanceAfterQueries(n int, queries [][]int) []int {
  // Membuat matriks/slice 2D untuk DP
	graph := make([][]int, n)
	for i := 0; i < n-1; i++ {
		graph[i] = append(graph[i], i+1)
	}

  // Alokasi slice integer
	ans := make([]int, len(queries))

	for qi, q := range queries {
		u, v := q[0], q[1]
		graph[u] = append(graph[u], v)

  // Alokasi slice integer
		dist := make([]int, n)
  // Range loop: iterasi dengan indeks + nilai
		for i := range dist {
			dist[i] = math.MaxInt32
		}
		dist[0] = 0
		queue := []int{0}

		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			for _, nb := range graph[cur] {
				if dist[cur]+1 < dist[nb] {
					dist[nb] = dist[cur] + 1
					queue = append(queue, nb)
				}
			}
		}
		ans[qi] = dist[n-1]
	}
	return ans
}

func main() {
	fmt.Println(shortestDistanceAfterQueries(5, [][]int{{2, 4}, {0, 2}, {0, 4}})) // Expected: [3, 2, 1]
	fmt.Println(shortestDistanceAfterQueries(4, [][]int{{0, 3}, {0, 2}}))          // Expected: [1, 1]
}
```
