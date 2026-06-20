# 2471 — Minimum Number Of Operations To Sort A Binary Tree By Level

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumOperations(root *TreeNode) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Two Pointer

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2471: Minimum Number of Operations to Sort a Binary Tree by Level
// https://leetcode.com/problems/minimum-number-of-operations-to-sort-a-binary-tree-by-level/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)
// BFS level-order. For each level, count min swaps to sort (cycle decomposition).

import (
	"fmt"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{1,
		&TreeNode{4,
			&TreeNode{7, nil, nil},
			&TreeNode{6, nil, nil},
		},
		&TreeNode{3,
			&TreeNode{8, nil, nil},
			&TreeNode{5, nil, nil},
		},
	}
	// Level 1: [1] sorted. Level 2: [4,3] -> swap, 1 op. Level 3: [7,6,8,5] -> 2 ops
	fmt.Println(minimumOperations(root)) // 3

	root2 := &TreeNode{1,
		&TreeNode{2, nil, nil},
		&TreeNode{3, nil, nil},
	}
	fmt.Println(minimumOperations(root2)) // 0
}

func minimumOperations(root *TreeNode) int {
	q := []*TreeNode{root}
	ans := 0
	for len(q) > 0 {
		n := len(q)
  // Alokasi slice integer
		vals := make([]int, n)
		for i := 0; i < n; i++ {
			vals[i] = q[i].Val
		}

		// Count min swaps to sort vals
		ans += minSwaps(vals)

		next := make([]*TreeNode, 0)
		for _, node := range q {
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		q = next
	}
	return ans
}

func minSwaps(arr []int) int {
	n := len(arr)
  // Alokasi slice integer
	sorted := make([]int, n)
	copy(sorted, arr)
  // Urutkan secara ascending — O(n log n)
	sort.Ints(sorted)

  // Membuat map (HashMap) — pencarian O(1)
	pos := make(map[int]int)
	for i, v := range arr {
		pos[v] = i
	}

	visited := make([]bool, n)
	swaps := 0
	for i := 0; i < n; i++ {
		if visited[i] || arr[i] == sorted[i] {
			continue
		}
		cycle := 0
		j := i
		for !visited[j] {
			visited[j] = true
			j = pos[sorted[j]]
			cycle++
		}
		swaps += cycle - 1
	}
	return swaps
}
```
