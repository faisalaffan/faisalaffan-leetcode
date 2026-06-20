# 3373 — Maximize The Number Of Target Nodes After Connecting Trees Ii

## Deskripsi

**Soal:** [3373. Maximize The Number Of Target Nodes After Connecting Trees Ii](https://leetcode.com/problems/maximize-the-number-of-target-nodes-after-connecting-trees-ii/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3373: Maximize the Number of Target Nodes After Connecting Trees II
// https://leetcode.com/problems/maximize-the-number-of-target-nodes-after-connecting-trees-ii/
// Difficulty: Hard
//
// Tree bipartition: color nodes by depth parity (even/odd).
// Even-length paths connect same-parity nodes. Connect tree1 node to tree2's
// larger parity group to maximize targets.

import "fmt"

func main() {
	// Example 1:
	// edges1 = [[0,1],[0,2],[2,3],[2,4]], edges2 = [[0,1],[0,2],[0,3],[2,7],[1,4],[4,5],[4,6]]
	// -> [8,7,7,8,8]
	fmt.Println(maxTargetNodes([][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}},
		[][]int{{0, 1}, {0, 2}, {0, 3}, {2, 7}, {1, 4}, {4, 5}, {4, 6}}))

	// Example 2:
	// edges1 = [[0,1],[0,2],[0,3],[0,4]], edges2 = [[0,1],[1,2],[2,3]]
	// -> [3,6,6,6,6]
	fmt.Println(maxTargetNodes([][]int{{0, 1}, {0, 2}, {0, 3}, {0, 4}},
		[][]int{{0, 1}, {1, 2}, {2, 3}}))

	// Single node in tree1, single node in tree2
	fmt.Println(maxTargetNodes([][]int{}, [][]int{}))

	// Small trees
	fmt.Println(maxTargetNodes([][]int{{0, 1}}, [][]int{{0, 1}}))
}

func maxTargetNodes(edges1 [][]int, edges2 [][]int) []int {
	// Process tree 1
	n1 := len(edges1) + 1
  // Membuat slice 2D untuk DP/tabel
	adj1 := make([][]int, n1)
	for _, e := range edges1 {
		u, v := e[0], e[1]
		adj1[u] = append(adj1[u], v)
		adj1[v] = append(adj1[v], u)
	}

  // Membuat slice untuk menyimpan hasil
	color1 := make([]int, n1)
	cnt1 := [2]int{}
	var dfs1 func(u, parent, col int)
	dfs1 = func(u, parent, col int) {
		color1[u] = col
		cnt1[col]++
		for _, v := range adj1[u] {
			if v != parent {
				dfs1(v, u, col^1)
			}
		}
	}
	dfs1(0, -1, 0)

	// Process tree 2
	n2 := len(edges2) + 1
  // Membuat slice 2D untuk DP/tabel
	adj2 := make([][]int, n2)
	for _, e := range edges2 {
		u, v := e[0], e[1]
		adj2[u] = append(adj2[u], v)
		adj2[v] = append(adj2[v], u)
	}

	cnt2 := [2]int{}
	var dfs2 func(u, parent, col int)
	dfs2 = func(u, parent, col int) {
		cnt2[col]++
		for _, v := range adj2[u] {
			if v != parent {
				dfs2(v, u, col^1)
			}
		}
	}
	dfs2(0, -1, 0)

	maxFromTree2 := max(cnt2[0], cnt2[1])

  // Membuat slice untuk menyimpan hasil
	ans := make([]int, n1)
	for i := 0; i < n1; i++ {
		ans[i] = cnt1[color1[i]] + maxFromTree2
	}
	return ans
}
```
