# 0449 — Serialize And Deserialize Bst

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func Constructor() Codec
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Bitmask

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #449: Serialize and Deserialize BST
// https://leetcode.com/problems/serialize-and-deserialize-bst/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct{}

func Constructor() Codec {
	return Codec{}
}

// Serialize using preorder traversal
func (c *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	result := []string{}
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		result = append(result, strconv.Itoa(node.Val))
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return strings.Join(result, ",")
}

// Deserialize using preorder + BST property
func (c *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}

	vals := strings.Split(data, ",")
  // Alokasi slice integer
	nums := make([]int, len(vals))
	for i, v := range vals {
		nums[i], _ = strconv.Atoi(v)
	}

	idx := 0
	var build func(lower, upper int) *TreeNode
	build = func(lower, upper int) *TreeNode {
		if idx >= len(nums) {
			return nil
		}
		val := nums[idx]
		if val < lower || val > upper {
			return nil
		}
		idx++
		return &TreeNode{
			Val:   val,
			Left:  build(lower, val),
			Right: build(val, upper),
		}
	}
	return build(-1<<31, 1<<31-1)
}

func inorderPrint(node *TreeNode) {
	if node == nil {
		return
	}
	inorderPrint(node.Left)
	fmt.Print(node.Val, " ")
	inorderPrint(node.Right)
}

func main() {
	c := Constructor()

	// Test case 1
	root1 := &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}
	data1 := c.serialize(root1)
	fmt.Println("Test 1 serialized:", data1)
	node1 := c.deserialize(data1)
	fmt.Print("Test 1 deserialized (inorder): ")
	inorderPrint(node1)
	fmt.Println()
	// Expected: "2,1,3" and inorder: 1 2 3

	// Test case 2: Nil
	data2 := c.serialize(nil)
	fmt.Println("Test 2 serialized:", data2)
	node2 := c.deserialize(data2)
	fmt.Print("Test 2 deserialized: ")
	inorderPrint(node2)
	fmt.Println()
	// Expected: "" and nothing

	// Test case 3
	root3 := &TreeNode{Val: 5, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 7, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 8}}}
	data3 := c.serialize(root3)
	fmt.Println("Test 3 serialized:", data3)
	node3 := c.deserialize(data3)
	fmt.Print("Test 3 deserialized (inorder): ")
	inorderPrint(node3)
	fmt.Println()
}
```
