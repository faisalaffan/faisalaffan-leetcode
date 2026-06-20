# 1042 — Flower Planting With No Adjacent

## Deskripsi

**Soal:** [1042. Flower Planting With No Adjacent](https://leetcode.com/problems/flower-planting-with-no-adjacent/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n + e) where e = len(paths)  
**Kompleksitas Ruang:** O(n + e)

**Algoritma:** Greedy (pemilihan optimal lokal)

> **Ide Kunci:** Greedy coloring with a graph. Each garden gets a flower type

## Solusi Go

```go
package main

// LeetCode #1042: Flower Planting With No Adjacent
// https://leetcode.com/problems/flower-planting-with-no-adjacent/
// Difficulty: Medium
//
// Approach: Greedy coloring with a graph. Each garden gets a flower type
//           not used by its neighbors.
// Time: O(n + e) where e = len(paths)
// Space: O(n + e)

import "fmt"

func main() {
	fmt.Println(gardenNoAdj(3, [][]int{{1, 2}, {2, 3}, {3, 1}})) // [1,2,3]
	fmt.Println(gardenNoAdj(4, [][]int{{1, 2}, {3, 4}}))         // [1,2,1,2]
}

func gardenNoAdj(n int, paths [][]int) []int {
  // Membuat slice 2D untuk DP/tabel
	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, 0)
	}

	for _, p := range paths {
		u, v := p[0]-1, p[1]-1
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

  // Membuat slice untuk menyimpan hasil
	result := make([]int, n)
	for i := 0; i < n; i++ {
  // Membuat slice untuk menyimpan hasil
		used := make([]bool, 5)
		for _, neighbor := range graph[i] {
			if result[neighbor] != 0 {
				used[result[neighbor]] = true
			}
		}
		for flower := 1; flower <= 4; flower++ {
			if !used[flower] {
				result[i] = flower
				break
			}
		}
	}

	return result
}
```
