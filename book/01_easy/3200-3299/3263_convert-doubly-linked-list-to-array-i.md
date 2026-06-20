# 3263 — Convert Doubly Linked List To Array I

## Deskripsi

**Soal:** [3263. Convert Doubly Linked List To Array I](https://leetcode.com/problems/convert-doubly-linked-list-to-array-i/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n). Space: O(n).  
**Kompleksitas Ruang:** O(n).

**Algoritma:** LIS (Longest Increasing Subsequence)

## Solusi Go

```go
package main

// LeetCode #3263: Convert Doubly Linked List to Array I
// https://leetcode.com/problems/convert-doubly-linked-list-to-array-i/
// Difficulty: Easy [Paid]

import "fmt"

func main() {
	// 1 <-> 2 <-> 3
	head := &Node{Val: 1}
	head.Next = &Node{Val: 2, Prev: head}
	head.Next.Next = &Node{Val: 3, Prev: head.Next}
	fmt.Println(ConvertDoublyLinkedListToArrayI(head))
}

// Node represents a doubly-linked list node.
type Node struct {
	Val  int
	Next *Node
	Prev *Node
}

// ConvertDoublyLinkedListToArrayI converts a doubly linked list to an integer array.
// Time: O(n). Space: O(n).
func ConvertDoublyLinkedListToArrayI(head *Node) []int {
	result := []int{}
	for cur := head; cur != nil; cur = cur.Next {
		result = append(result, cur.Val)
	}
	return result
}
```
