# 1628 — Design An Expression Tree With Evaluate Function

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func BuildExpressionTree(postfix []string) *Node
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Stack

**Kompleksitas Waktu:** O(N), Space: O(N)  
**Kompleksitas Ruang:** O(N)

> **Untuk fresh graduate:** Kuasai dulu teknik **Stack** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1628: Design an Expression Tree With Evaluate Function
// https://leetcode.com/problems/design-an-expression-tree-with-evaluate-function/
// Difficulty: Medium [Paid]

import (
	"fmt"
	"strconv"
)

func main() {
	// Build expression: (3 + 4) * (5 - 2)
	// Postfix: ["3", "4", "+", "5", "2", "-", "*"]
	postfix := []string{"3", "4", "+", "5", "2", "-", "*"}
	tree := BuildExpressionTree(postfix)
	result := tree.Evaluate()
	fmt.Println("Result:", result) // should be 21

	// Simple: 2 + 3
	postfix2 := []string{"2", "3", "+"}
	tree2 := BuildExpressionTree(postfix2)
	fmt.Println("Result:", tree2.Evaluate()) // should be 5
}

// Node is an expression tree node.
type Node struct {
	val   string
	left  *Node
	right *Node
}

func (n *Node) Evaluate() int {
	if n.left == nil && n.right == nil {
		val, _ := strconv.Atoi(n.val)
		return val
	}

	leftVal := n.left.Evaluate()
	rightVal := n.right.Evaluate()

	switch n.val {
	case "+":
		return leftVal + rightVal
	case "-":
		return leftVal - rightVal
	case "*":
		return leftVal * rightVal
	case "/":
		return leftVal / rightVal
	}
	return 0
}

func BuildExpressionTree(postfix []string) *Node {
	// Time: O(N), Space: O(N)
	stack := make([]*Node, 0)

	for _, token := range postfix {
		node := &Node{val: token}
		if token == "+" || token == "-" || token == "*" || token == "/" {
			node.right = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			node.left = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, node)
	}

	return stack[0]
}
```
