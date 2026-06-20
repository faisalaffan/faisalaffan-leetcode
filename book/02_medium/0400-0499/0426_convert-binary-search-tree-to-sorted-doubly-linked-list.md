# 0426 — Convert Binary Search Tree To Sorted Doubly Linked List

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan pohon (tree). Tugasmu menjelajahi atau memanipulasi struktur pohon.

**Cara berpikir:** TreeNode punya Val, Left, Right. Gunakan DFS rekursif (pre/in/post-order).

**Fungsi Solusi:** `func treeToDoublyList(root *Node) *Node`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Binary Search

**Waktu:** O(n)  |  **Ruang:** O(h)

> 🎓 **Fresh Grad Tips:** Kuasai **Binary Search** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #426: Convert Binary Search Tree to Sorted Doubly Linked List
// https://leetcode.com/problems/convert-binary-search-tree-to-sorted-doubly-linked-list/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(h)

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func treeToDoublyList(root *Node) *Node {
	if root == nil {
		return nil
	}

	var first, last *Node

	var inorder func(node *Node)
	inorder = func(node *Node) {
		if node == nil {
			return
		}
		inorder(node.Left)

		if last != nil {
			last.Right = node
			node.Left = last
		} else {
			first = node
		}
		last = node

		inorder(node.Right)
	}

	inorder(root)

	// Close the circular doubly linked list
	last.Right = first
	first.Left = last
	return first
}

func printList(head *Node) {
	if head == nil {
		fmt.Println("nil")
		return
	}
	cur := head
	for {
		fmt.Print(cur.Val)
		cur = cur.Right
		if cur == head {
			break
		}
		fmt.Print(" <-> ")
	}
	fmt.Println(" (circular)")
}

func main() {
	// Test case 1: [4,2,5,1,3]
	root1 := &Node{Val: 4}
	root1.Left = &Node{Val: 2, Left: &Node{Val: 1}, Right: &Node{Val: 3}}
	root1.Right = &Node{Val: 5}
	fmt.Print("Test 1: ")
	printList(treeToDoublyList(root1))
	// Expected: 1 <-> 2 <-> 3 <-> 4 <-> 5 (circular)

	// Test case 2: Single node
	root2 := &Node{Val: 1}
	fmt.Print("Test 2: ")
	printList(treeToDoublyList(root2))
	// Expected: 1 (circular)

	// Test case 3: Nil
	fmt.Print("Test 3: ")
	printList(treeToDoublyList(nil))
	// Expected: nil
}
```
