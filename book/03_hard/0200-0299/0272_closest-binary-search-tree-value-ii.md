# 0272 — Closest Binary Search Tree Value Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func closestKValues(root *TreeNode, target float64, k int) []int
```

> **💡 Hint:** Inorder traversal + two-pointer sliding window.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Sliding Window, Binary Search

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #272: Closest Binary Search Tree Value II
// https://leetcode.com/problems/closest-binary-search-tree-value-ii/
// Difficulty: Hard [Paid]
//
// Approach: Inorder traversal + two-pointer sliding window.
//  1. Inorder traversal of BST gives sorted values.
//  2. Use two-pointer to maintain a window of k elements closest to target.
//  3. Expand from the window edges, always removing the farther element.

import (
	"fmt"
	"math"
)

// TreeNode definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Example: root=[4,2,5,1,3], target=3.714, k=2 -> [3,4]
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{Val: 5},
	}

	result := closestKValues(root, 3.714, 2)
	fmt.Println("Closest K Values:", result)

	// Test: k=3
	result2 := closestKValues(root, 3.714, 3)
	fmt.Println("Closest K Values (k=3):", result2)
}

// closestKValues returns k closest values to target in BST root.
func closestKValues(root *TreeNode, target float64, k int) []int {
	// Collect inorder traversal values.
	var values []int
	inorder(root, &values)

	// Two-pointer: find the window of size k.
	// First, find the starting point using binary search to locate first element >= target.
	idx := lowerBound(values, target)

	// Expand window: left goes backwards, right goes forwards.
	left := idx - 1
	right := idx

	for right-left-1 < k {
		if left < 0 {
			// Only right side available.
			right++
		} else if right >= len(values) {
			// Only left side available.
			left--
		} else {
			// Compare distances.
			distLeft := target - float64(values[left])
			distRight := float64(values[right]) - target
			if distLeft < distRight {
				left--
			} else {
				right++
			}
		}
	}

	// Extract window (left+1 ... right-1).
	return values[left+1 : right]
}

// inorder performs inorder traversal of BST.
func inorder(node *TreeNode, values *[]int) {
	if node == nil {
		return
	}
	inorder(node.Left, values)
	*values = append(*values, node.Val)
	inorder(node.Right, values)
}

// lowerBound finds the first index where values[i] >= target.
func lowerBound(values []int, target float64) int {
	lo, hi := 0, len(values)
	for lo < hi {
		mid := lo + (hi-lo)/2
		if float64(values[mid]) < target {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

// --- Alternative: O(n) without extra space using predecessor/successor ---

// closestKValuesO1Space uses inorder + reverse inorder to fill k closest.
// Not implemented here for brevity; the two-pointer approach above is standard.

// --- Older solution stub compatibility ---
func ClosestBinarySearchTreeValueIi() any {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{Val: 5},
	}
	return closestKValues(root, 3.714, 2)
}

// --- Testing utilities ---

// Helper to verify floating point comparison.
func equalFloat(a, b float64) bool {
	const eps = 1e-9
	return math.Abs(a-b) < eps
}
```
