# 0608 — Tree Node

## Deskripsi

**Soal:** [0608. Tree Node](https://leetcode.com/problems/tree-node/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #608: Tree Node
// https://leetcode.com/problems/tree-node/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	// Tree nodes: {id, p_id} where p_id = -1 means root
	nodes := [][]int{
		{1, -1},
		{2, 1},
		{3, 1},
		{4, 2},
		{5, 2},
	}
	fmt.Println(ClassifyTreeNodes(nodes))
}

func ClassifyTreeNodes(nodes [][]int) map[int]string {
  // Membuat map untuk pencarian O(1): key → value
	parentMap := make(map[int]int)
  // Membuat map untuk pencarian O(1): key → value
	childrenMap := make(map[int][]int)

	for _, node := range nodes {
		id, pID := node[0], node[1]
		if pID == -1 {
			parentMap[id] = -1
		} else {
			parentMap[id] = pID
			childrenMap[pID] = append(childrenMap[pID], id)
		}
	}

  // Membuat map untuk pencarian O(1): key → value
	result := make(map[int]string)
	for _, node := range nodes {
		id := node[0]
		if parentMap[id] == -1 {
			result[id] = "Root"
		} else if len(childrenMap[id]) == 0 {
			result[id] = "Leaf"
		} else {
			result[id] = "Inner"
		}
	}

	return result
}
```
