# 2385 — Amount Of Time For Binary Tree To Be Infected

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func amountOfTime(root *TreeNode, start int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Two Pointer, BFS

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2385: Amount of Time for Binary Tree to Be Infected
// https://leetcode.com/problems/amount-of-time-for-binary-tree-to-be-infected/
// Difficulty: Medium
// Time: O(n) | Space: O(n)
// Convert tree to graph, BFS from start node to find farthest distance.

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Tree: [1,5,3,null,4,10,6,9,2]
	root := &TreeNode{1,
		&TreeNode{5, nil,
			&TreeNode{4,
				&TreeNode{9, nil, nil},
				&TreeNode{2, nil, nil},
			},
		},
		&TreeNode{3,
			&TreeNode{10, nil, nil},
			&TreeNode{6, nil, nil},
		},
	}
	fmt.Println(amountOfTime(root, 3)) // 4

	root2 := &TreeNode{1,
		&TreeNode{2,
			&TreeNode{3,
				&TreeNode{4,
					&TreeNode{5, nil, nil},
					nil,
				},
				nil,
			},
			nil,
		},
		nil,
	}
	fmt.Println(amountOfTime(root2, 1)) // 4
}

func amountOfTime(root *TreeNode, start int) int {
  // Membuat map (HashMap) — pencarian O(1)
	graph := make(map[int][]int)
	var buildGraph func(node *TreeNode)
	buildGraph = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil {
			graph[node.Val] = append(graph[node.Val], node.Left.Val)
			graph[node.Left.Val] = append(graph[node.Left.Val], node.Val)
			buildGraph(node.Left)
		}
		if node.Right != nil {
			graph[node.Val] = append(graph[node.Val], node.Right.Val)
			graph[node.Right.Val] = append(graph[node.Right.Val], node.Val)
			buildGraph(node.Right)
		}
	}
	buildGraph(root)

	visited := map[int]bool{start: true}
	queue := []int{start}
	time := -1
	for len(queue) > 0 {
		time++
		for sz := len(queue); sz > 0; sz-- {
			u := queue[0]
			queue = queue[1:]
			for _, v := range graph[u] {
				if !visited[v] {
					visited[v] = true
					queue = append(queue, v)
				}
			}
		}
	}
	return time
}
```
