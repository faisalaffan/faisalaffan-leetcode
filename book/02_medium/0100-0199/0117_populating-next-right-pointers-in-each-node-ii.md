# 0117 — Populating Next Right Pointers In Each Node Ii

## Deskripsi

**Soal:** [0117. Populating Next Right Pointers In Each Node Ii](https://leetcode.com/problems/populating-next-right-pointers-in-each-node-ii/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func connect(root *Node) *Node`

## Solusi Go

```go
package main

// LeetCode #117: Populating Next Right Pointers in Each Node II
// https://leetcode.com/problems/populating-next-right-pointers-in-each-node-ii/
// Difficulty: Medium

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	curr := root
	for curr != nil {
		dummy := &Node{}
		tail := dummy

		for curr != nil {
			if curr.Left != nil {
				tail.Next = curr.Left
				tail = tail.Next
			}
			if curr.Right != nil {
				tail.Next = curr.Right
				tail = tail.Next
			}
			curr = curr.Next
		}

		curr = dummy.Next
	}

	return root
}

func printLevels(root *Node) {
	curr := root
	for curr != nil {
		head := curr
		for head != nil {
			nextStr := "null"
			if head.Next != nil {
				nextStr = fmt.Sprintf("%d", head.Next.Val)
			}
			fmt.Printf("%d->%s ", head.Val, nextStr)
			head = head.Next
		}
		fmt.Println()
		curr = curr.Left
		if curr == nil {
			// Follow Next for non-perfect trees
			curr = root.Next
			for curr != nil && curr.Left == nil && curr.Right == nil {
				curr = curr.Next
			}
		}
	}
}

func main() {
	// Test case 1: [1,2,3,4,5,null,7]
	root := &Node{Val: 1,
		Left:  &Node{Val: 2, Left: &Node{Val: 4}, Right: &Node{Val: 5}},
		Right: &Node{Val: 3, Right: &Node{Val: 7}}}
	connect(root)
	printLevels(root)

	// Test case 2
	root = nil
	connect(root)
	fmt.Println("nil")
}

// Time: O(n) | Space: O(1)
```
