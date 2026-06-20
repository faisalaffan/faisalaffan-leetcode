# 0608 — Tree Node

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func ClassifyTreeNodes(nodes [][]int) map[int]string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

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
  // Membuat map (HashMap) — pencarian O(1)
	parentMap := make(map[int]int)
  // Membuat map (HashMap) — pencarian O(1)
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

  // Membuat map (HashMap) — pencarian O(1)
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
