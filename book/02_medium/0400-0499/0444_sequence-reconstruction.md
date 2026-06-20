# 0444 — Sequence Reconstruction

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan graf. Tugasmu menjelajahi atau menganalisis konektivitas graf.

**Cara berpikir:** Adjacency list `map[int][]int`. Gunakan BFS (queue) atau DFS (rekursif) dengan visited set untuk hindari siklus.

**Fungsi Solusi:** `func sequenceReconstruction(org []int, seqs [][]int) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** BFS

**Waktu:** O(n + m)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **BFS** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #444: Sequence Reconstruction
// https://leetcode.com/problems/sequence-reconstruction/
// Difficulty: Medium [Paid]
// Time: O(n + m) | Space: O(n)

import "fmt"

func sequenceReconstruction(org []int, seqs [][]int) bool {
	n := len(org)
	// Build indegree map and edges
  // Alokasi slice
	indegree := make([]int, n+1)
  // Matriks 2D
	graph := make([][]int, n+1)
	exists := make([]bool, n+1)

	for _, seq := range seqs {
		for _, num := range seq {
			if num < 1 || num > n {
				return false
			}
			exists[num] = true
		}
  // Linear scan O(n)
		for i := 0; i < len(seq)-1; i++ {
			u, v := seq[i], seq[i+1]
			graph[u] = append(graph[u], v)
			indegree[v]++
		}
	}

	// Check all numbers exist
	for i := 1; i <= n; i++ {
		if !exists[i] {
			return false
		}
	}

	// BFS: only one node with indegree 0 at each step
	queue := []int{}
	for i := 1; i <= n; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	idx := 0
	for len(queue) == 1 {
		u := queue[0]
		queue = queue[1:]
		if u != org[idx] {
			return false
		}
		idx++
		for _, v := range graph[u] {
			indegree[v]--
			if indegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	return idx == n
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", sequenceReconstruction([]int{1, 2, 3}, [][]int{{1, 2}, {1, 3}, {2, 3}}))
	// Expected: true

	// Test case 2
	fmt.Println("Test 2:", sequenceReconstruction([]int{1, 2, 3}, [][]int{{1, 2}, {1, 3}}))
	// Expected: false

	// Test case 3
	fmt.Println("Test 3:", sequenceReconstruction([]int{1, 2, 3}, [][]int{{1, 2}, {2, 3}, {3, 1}}))
	// Expected: false (cycle)
}
```
