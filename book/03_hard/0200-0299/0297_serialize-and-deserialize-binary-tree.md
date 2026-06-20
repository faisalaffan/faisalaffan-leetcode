# 0297 — Serialize And Deserialize Binary Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan pohon (tree). Tugasmu menjelajahi atau memanipulasi struktur pohon.

**Cara berpikir:** TreeNode punya Val, Left, Right. Gunakan DFS rekursif (pre/in/post-order).

**Fungsi Solusi:** `func Constructor297() Codec`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, BFS

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #297: Serialize and Deserialize Binary Tree
// https://leetcode.com/problems/serialize-and-deserialize-binary-tree/
// Difficulty: Hard
//
// Approach: BFS Level-Order (Iterative).
//   - Serialize: Level-order traversal using a queue. Use "null" for nil nodes.
//   - Trailing nulls are stripped to keep the serialized string concise.
//   - Deserialize: Use a queue to rebuild the tree from the BFS order.
//
// The serialized format is a comma-separated string like "4,2,5,1,3,null,null".

import (
	"fmt"
	"strconv"
	"strings"
)

// TreeNode definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Example: root=[4,2,5,1,3]
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{Val: 5},
	}

	codec := Constructor297()

	// Serialize.
	data := codec.serialize(root)
	fmt.Println("Serialized:", data)

	// Deserialize.
	decoded := codec.deserialize(data)
	fmt.Println("Deserialized root value:", decoded.Val)
	fmt.Println("Serialized again:", codec.serialize(decoded))

	// Test empty tree.
	emptyData := codec.serialize(nil)
	fmt.Println("Empty tree serialize:", emptyData)
	decodedEmpty := codec.deserialize(emptyData)
	fmt.Println("Empty tree deserialize:", decodedEmpty)

	// Test single node.
	single := &TreeNode{Val: 1}
	singleData := codec.serialize(single)
	fmt.Println("Single node serialize:", singleData)
	decodedSingle := codec.deserialize(singleData)
	fmt.Println("Single node deserialize root:", decodedSingle.Val)
}

// Codec handles serialization/deserialization.
type Codec struct{}

// Constructor297 creates a new Codec.
func Constructor297() Codec {
	return Codec{}
}

// serialize serializes a binary tree to a string.
func (c *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	var parts []string
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			parts = append(parts, "null")
		} else {
			parts = append(parts, strconv.Itoa(node.Val))
			queue = append(queue, node.Left, node.Right)
		}
	}

	// Strip trailing "null"s.
	end := len(parts) - 1
	for end >= 0 && parts[end] == "null" {
		end--
	}
	parts = parts[:end+1]

	return strings.Join(parts, ",")
}

// deserialize deserializes a string back to a binary tree.
func (c *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}

	parts := strings.Split(data, ",")
	if len(parts) == 0 || parts[0] == "" {
		return nil
	}

	val, _ := strconv.Atoi(parts[0])
	root := &TreeNode{Val: val}
	queue := []*TreeNode{root}
	index := 1

	for len(queue) > 0 && index < len(parts) {
		node := queue[0]
		queue = queue[1:]

		// Process left child.
		if index < len(parts) && parts[index] != "null" {
			val, _ = strconv.Atoi(parts[index])
			node.Left = &TreeNode{Val: val}
			queue = append(queue, node.Left)
		}
		index++

		// Process right child.
		if index < len(parts) && parts[index] != "null" {
			val, _ = strconv.Atoi(parts[index])
			node.Right = &TreeNode{Val: val}
			queue = append(queue, node.Right)
		}
		index++
	}

	return root
}

// Stub compatibility.
func SerializeAndDeserializeBinaryTree() any {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{Val: 5},
	}
	codec := Constructor297()
	data := codec.serialize(root)
	decoded := codec.deserialize(data)
	return decoded.Val // should be 4
}
```
