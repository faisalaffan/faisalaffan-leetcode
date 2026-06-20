# 0129 — Sum Root To Leaf Numbers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func sumNumbers(root *TreeNode) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #129: Sum Root to Leaf Numbers
// https://leetcode.com/problems/sum-root-to-leaf-numbers/
// Difficulty: Medium

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	var dfs func(node *TreeNode, curr int) int
	dfs = func(node *TreeNode, curr int) int {
		if node == nil {
			return 0
		}
		curr = curr*10 + node.Val
		if node.Left == nil && node.Right == nil {
			return curr
		}
		return dfs(node.Left, curr) + dfs(node.Right, curr)
	}
	return dfs(root, 0)
}

func main() {
	// Test case 1: [1,2,3] -> 12+13=25
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	fmt.Println(sumNumbers(root)) // 25

	// Test case 2: [4,9,0,5,1] -> 495+491+40=1026
	root = &TreeNode{Val: 4,
		Left:  &TreeNode{Val: 9, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 1}},
		Right: &TreeNode{Val: 0}}
	fmt.Println(sumNumbers(root)) // 1026

	// Test case 3
	fmt.Println(sumNumbers(nil)) // 0
}

// Time: O(n) | Space: O(n)
```
