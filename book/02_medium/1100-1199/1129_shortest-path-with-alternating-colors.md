# 1129 — Shortest Path With Alternating Colors

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int
```

> **💡 Hint:** BFS with 2 states per node (reached by red edge / blue edge)

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** BFS

**Kompleksitas Waktu:** O(n + e)  
**Kompleksitas Ruang:** O(n + e)

> **Untuk fresh graduate:** Kuasai dulu teknik **BFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1129: Shortest Path with Alternating Colors
// https://leetcode.com/problems/shortest-path-with-alternating-colors/
// Difficulty: Medium
//
// Approach: BFS with 2 states per node (reached by red edge / blue edge)
// Time: O(n + e)
// Space: O(n + e)

import "fmt"

func main() {
	fmt.Println(shortestAlternatingPaths(3, [][]int{{0, 1}}, [][]int{{2, 1}}))          // [0,1,-1]
	fmt.Println(shortestAlternatingPaths(3, [][]int{{0, 1}}, [][]int{{1, 2}}))          // [0,1,2]
}

func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
  // Membuat matriks/slice 2D untuk DP
	redGraph := make([][]int, n)
  // Membuat matriks/slice 2D untuk DP
	blueGraph := make([][]int, n)

	for _, e := range redEdges {
		redGraph[e[0]] = append(redGraph[e[0]], e[1])
	}
	for _, e := range blueEdges {
		blueGraph[e[0]] = append(blueGraph[e[0]], e[1])
	}

	// dist[node][0] = distance reaching node via red edge
	// dist[node][1] = distance reaching node via blue edge
  // Alokasi slice integer
	dist := make([][2]int, n)
	for i := 1; i < n; i++ {
		dist[i] = [2]int{-1, -1}
	}

  // Alokasi slice integer
	queue := make([][2]int, 0)
	queue = append(queue, [2]int{0, 0}) // reached 0 via red (start counts as either)
	queue = append(queue, [2]int{0, 1}) // reached 0 via blue

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		node, color := cur[0], cur[1]
		nextColor := 1 - color

		if nextColor == 0 { // next edge should be red
			for _, next := range redGraph[node] {
				if dist[next][nextColor] == -1 {
					dist[next][nextColor] = dist[node][color] + 1
					queue = append(queue, [2]int{next, nextColor})
				}
			}
		} else { // next edge should be blue
			for _, next := range blueGraph[node] {
				if dist[next][nextColor] == -1 {
					dist[next][nextColor] = dist[node][color] + 1
					queue = append(queue, [2]int{next, nextColor})
				}
			}
		}
	}

  // Alokasi slice integer
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = minDist(dist[i])
	}
	return result
}

func minDist(d [2]int) int {
	if d[0] == -1 {
		return d[1]
	}
	if d[1] == -1 {
		return d[0]
	}
	if d[0] < d[1] {
		return d[0]
	}
	return d[1]
}
```
