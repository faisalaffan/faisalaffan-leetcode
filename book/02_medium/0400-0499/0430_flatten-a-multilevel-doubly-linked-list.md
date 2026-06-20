# 0430 — Flatten A Multilevel Doubly Linked List

## Deskripsi

**Soal:** [0430. Flatten A Multilevel Doubly Linked List](https://leetcode.com/problems/flatten-a-multilevel-doubly-linked-list/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(d)

**Algoritma:** DFS (Depth-First Search / pencarian kedalaman), LIS (Longest Increasing Subsequence)

**Fungsi Solusi:** `func flatten(root *Node) *Node`

## Solusi Go

```go
package main

// LeetCode #430: Flatten a Multilevel Doubly Linked List
// https://leetcode.com/problems/flatten-a-multilevel-doubly-linked-list/
// Difficulty: Medium
// Time: O(n) | Space: O(d)

import "fmt"

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	flattenDFS(root)
	return root
}

// Returns the tail of the flattened list
func flattenDFS(node *Node) *Node {
	current := node
	var last *Node

	for current != nil {
		next := current.Next

		if current.Child != nil {
			childTail := flattenDFS(current.Child)

			// Connect current to child
			current.Next = current.Child
			current.Child.Prev = current

			// Connect child tail to next
			if next != nil {
				childTail.Next = next
				next.Prev = childTail
			}

			current.Child = nil
			last = childTail
		} else {
			last = current
		}
		current = next
	}
	return last
}

func printList(head *Node) {
	for head != nil {
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Print(" <-> ")
		}
		head = head.Next
	}
	fmt.Println()
}

func main() {
	// Build: 1 - 2 - 3 - 4 - 5 - 6
	//            |
	//            7 - 8 - 9 - 10
	//                |
	//               11 - 12
	n1 := &Node{Val: 1}
	n2 := &Node{Val: 2}
	n3 := &Node{Val: 3}
	n4 := &Node{Val: 4}
	n5 := &Node{Val: 5}
	n6 := &Node{Val: 6}
	n7 := &Node{Val: 7}
	n8 := &Node{Val: 8}
	n9 := &Node{Val: 9}
	n10 := &Node{Val: 10}
	n11 := &Node{Val: 11}
	n12 := &Node{Val: 12}

	n1.Next = n2
	n2.Prev = n1
	n2.Next = n3
	n3.Prev = n2
	n3.Next = n4
	n4.Prev = n3
	n4.Next = n5
	n5.Prev = n4
	n5.Next = n6
	n6.Prev = n5

	n3.Child = n7
	n7.Next = n8
	n8.Prev = n7
	n8.Next = n9
	n9.Prev = n8
	n9.Next = n10
	n10.Prev = n9

	n8.Child = n11
	n11.Next = n12
	n12.Prev = n11

	fmt.Print("Test 1: ")
	printList(flatten(n1))
	// Expected: 1 <-> 2 <-> 3 <-> 7 <-> 8 <-> 11 <-> 12 <-> 9 <-> 10 <-> 4 <-> 5 <-> 6

	// Test case 2: Nil
	fmt.Print("Test 2: ")
	printList(flatten(nil))

	// Test case 3: No children
	n21 := &Node{Val: 1}
	n22 := &Node{Val: 2}
	n21.Next = n22
	n22.Prev = n21
	fmt.Print("Test 3: ")
	printList(flatten(n21))
	// Expected: 1 <-> 2
}
```
