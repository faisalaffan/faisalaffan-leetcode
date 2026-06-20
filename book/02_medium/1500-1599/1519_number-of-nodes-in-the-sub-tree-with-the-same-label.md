# 1519 — Number Of Nodes In The Sub Tree With The Same Label

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan graf. Tugasmu menjelajahi atau menganalisis konektivitas graf.

**Cara berpikir:** Adjacency list `map[int][]int`. Gunakan BFS (queue) atau DFS (rekursif) dengan visited set untuk hindari siklus.

**Fungsi Solusi:** `func CountSubTrees(n int, edges [][]int, labels string) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DFS

**Waktu:** O(N), Space: O(N)  |  **Ruang:** O(N)

> 🎓 **Fresh Grad Tips:** Kuasai **DFS** — sering muncul di interview!

## 💻 Solusi Go

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
  // Matriks 2D
	graph := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

  // Alokasi slice
	result := make([]int, n)
	visited := make([]bool, n)

	var dfs func(node int) []int
	dfs = func(node int) []int {
		visited[node] = true
		// Count array for 26 lowercase letters
  // Alokasi slice
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
