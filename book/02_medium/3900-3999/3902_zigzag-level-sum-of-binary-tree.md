# 3902 — Zigzag Level Sum Of Binary Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func ZigzagLevelSumOfBinaryTree(root *TreeNode) []int
```

> **💡 Hint:** BFS level-order. At odd levels (left-to-right), stop at first

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, BFS

**Kompleksitas Waktu:** O(N)  
**Kompleksitas Ruang:** O(N)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3902: Zigzag Level Sum of Binary Tree
// https://leetcode.com/problems/zigzag-level-sum-of-binary-tree/
// Difficulty: Medium [Paid]
// Time: O(N) | Space: O(N)
// Approach: BFS level-order. At odd levels (left-to-right), stop at first
// node without left child. At even levels (right-to-left), stop at first
// node without right child. Collect children from all nodes.

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func ZigzagLevelSumOfBinaryTree(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ans := []int{}
	queue := []*TreeNode{root}
	level := 1

	for len(queue) > 0 {
		nextQ := []*TreeNode{}
		sum := 0

		if level%2 == 1 {
			// Odd: left to right, stop at first node without LEFT child
			for _, node := range queue {
				if node.Left == nil {
					break
				}
				sum += node.Val
			}
		} else {
			// Even: right to left, stop at first node without RIGHT child
			for i := len(queue) - 1; i >= 0; i-- {
				if queue[i].Right == nil {
					break
				}
				sum += queue[i].Val
			}
		}

		// Collect children from ALL nodes at this level
		for _, node := range queue {
			if node.Left != nil {
				nextQ = append(nextQ, node.Left)
			}
			if node.Right != nil {
				nextQ = append(nextQ, node.Right)
			}
		}

		ans = append(ans, sum)
		queue = nextQ
		level++
	}

	return ans
}

func makeTree(vals []interface{}) *TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}
	root := &TreeNode{Val: vals[0].(int)}
	queue := []*TreeNode{root}
	idx := 1
	for len(queue) > 0 && idx < len(vals) {
		node := queue[0]
		queue = queue[1:]
		if idx < len(vals) && vals[idx] != nil {
			node.Left = &TreeNode{Val: vals[idx].(int)}
			queue = append(queue, node.Left)
		}
		idx++
		if idx < len(vals) && vals[idx] != nil {
			node.Right = &TreeNode{Val: vals[idx].(int)}
			queue = append(queue, node.Right)
		}
		idx++
	}
	return root
}

func main() {
	// Example 1: root = [5,2,8,1,null,9,6]
	root1 := makeTree([]interface{}{5, 2, 8, 1, nil, 9, 6})
	fmt.Println(ZigzagLevelSumOfBinaryTree(root1)) // Expected: [5 8 0]

	// Example 2: root = [1,2,3,4,5,null,7]
	root2 := makeTree([]interface{}{1, 2, 3, 4, 5, nil, 7})
	fmt.Println(ZigzagLevelSumOfBinaryTree(root2)) // Expected: [1 5 0]
}
```
