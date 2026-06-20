# 0337 — House Robber Iii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func rob(root *TreeNode) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(h)

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #337: House Robber III
// https://leetcode.com/problems/house-robber-iii/
// Difficulty: Medium
// Time: O(n) | Space: O(h)

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	robRoot, skipRoot := dfs(root)
	if robRoot > skipRoot {
		return robRoot
	}
	return skipRoot
}

// Returns (robThis, skipThis)
func dfs(node *TreeNode) (int, int) {
	if node == nil {
		return 0, 0
	}

	leftRob, leftSkip := dfs(node.Left)
	rightRob, rightSkip := dfs(node.Right)

	// If we rob this node, we must skip children
	robThis := node.Val + leftSkip + rightSkip
	// If we skip this node, we can take best of each child
	skipThis := max(leftRob, leftSkip) + max(rightRob, rightSkip)

	return robThis, skipThis
}

func main() {
	// Test case 1: [3,2,3,null,3,null,1]
	root1 := &TreeNode{Val: 3}
	root1.Left = &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}
	root1.Right = &TreeNode{Val: 3, Right: &TreeNode{Val: 1}}
	fmt.Println("Test 1:", rob(root1))
	// Expected: 7 (3+3+1)

	// Test case 2: [3,4,5,1,3,null,1]
	root2 := &TreeNode{Val: 3}
	root2.Left = &TreeNode{Val: 4, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}
	root2.Right = &TreeNode{Val: 5, Right: &TreeNode{Val: 1}}
	fmt.Println("Test 2:", rob(root2))
	// Expected: 9

	// Test case 3: Single node
	root3 := &TreeNode{Val: 10}
	fmt.Println("Test 3:", rob(root3))
	// Expected: 10
}
```
