# 1325 — Delete Leaves With A Given Value

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func removeLeafNodes(root *TreeNode, target int) *TreeNode
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Stack

**Kompleksitas Waktu:** O(n) where n is number of nodes  
**Kompleksitas Ruang:** O(h) for recursion stack, h = tree height

> **Untuk fresh graduate:** Kuasai dulu teknik **Stack** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1325: Delete Leaves With a Given Value
// https://leetcode.com/problems/delete-leaves-with-a-given-value/
// Difficulty: Medium

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Test case 1
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 2}},
		Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 4}},
	}
	result := removeLeafNodes(root, 2)
	fmt.Println(result) // [1,null,3,null,4]

	// Test case 2
	root2 := &TreeNode{Val: 1, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 1}}
	result2 := removeLeafNodes(root2, 1)
	fmt.Println(result2) // nil

	// Test case 3
	root3 := &TreeNode{Val: 1, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 2}}}
	result3 := removeLeafNodes(root3, 3)
	fmt.Println(result3) // [1,3,null,null,2]
}

// Time: O(n) where n is number of nodes
// Space: O(h) for recursion stack, h = tree height
func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left = removeLeafNodes(root.Left, target)
	root.Right = removeLeafNodes(root.Right, target)

	// Post-order: if current node is a leaf and its value equals target, delete it
	if root.Left == nil && root.Right == nil && root.Val == target {
		return nil
	}

	return root
}

// For printing in main
func (n *TreeNode) String() string {
	if n == nil {
		return "nil"
	}
	return fmt.Sprintf("%d %s %s", n.Val, n.Left, n.Right)
}
```
