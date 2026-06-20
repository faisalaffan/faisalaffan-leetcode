# 1519 — Number Of Nodes In The Sub Tree With The Same Label

## Deskripsi

**Soal:** [1519. Number Of Nodes In The Sub Tree With The Same Label](https://leetcode.com/problems/number-of-nodes-in-the-sub-tree-with-the-same-label/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(N), Space: O(N)  
**Kompleksitas Ruang:** O(N)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1519: Number of Nodes in the Sub-Tree With the Same Label
// https://leetcode.com/problems/number-of-nodes-in-the-sub-tree-with-the-same-label/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(CountSubTrees(7, [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}}, "abaedcd"))
	fmt.Println(CountSubTrees(4, [][]int{{0, 1}, {1, 2}, {0, 3}}, "bbbb"))
	fmt.Println(CountSubTrees(5, [][]int{{0, 1}, {0, 2}, {1, 3}, {0, 4}}, "aabab"))
}

func CountSubTrees(n int, edges [][]int, labels string) []int {
	// Time: O(N), Space: O(N)
  // Membuat slice 2D untuk DP/tabel
	graph := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

  // Membuat slice untuk menyimpan hasil
	result := make([]int, n)
  // Membuat slice untuk menyimpan hasil
	visited := make([]bool, n)

	var dfs func(node int) []int
	dfs = func(node int) []int {
		visited[node] = true
		// Count array for 26 lowercase letters
  // Membuat slice untuk menyimpan hasil
		count := make([]int, 26)
		count[labels[node]-'a'] = 1

		for _, nei := range graph[node] {
			if visited[nei] {
				continue
			}
			childCount := dfs(nei)
			for i := 0; i < 26; i++ {
				count[i] += childCount[i]
			}
		}

		result[node] = count[labels[node]-'a']
		return count
	}

	dfs(0)
	return result
}
```
