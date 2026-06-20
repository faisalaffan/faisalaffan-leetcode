# 0404 — Sum Of Left Leaves

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func SumOfLeftLeaves(root *TreeNode) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(h)  
**Kompleksitas Ruang:** O(h)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #404: Sum of Left Leaves
// https://leetcode.com/problems/sum-of-left-leaves/
// Difficulty: Easy

import "fmt"

// TreeNode represents a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n), Space: O(h)
func SumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := 0
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		sum += root.Left.Val
	}
	sum += SumOfLeftLeaves(root.Left)
	sum += SumOfLeftLeaves(root.Right)
	return sum
}

func main() {
	// Test case 1: [3,9,20,null,null,15,7]
	root1 := &TreeNode{
		Val: 3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}
	fmt.Println(SumOfLeftLeaves(root1))

	// Test case 2: [1]
	root2 := &TreeNode{Val: 1}
	fmt.Println(SumOfLeftLeaves(root2))
}
```
