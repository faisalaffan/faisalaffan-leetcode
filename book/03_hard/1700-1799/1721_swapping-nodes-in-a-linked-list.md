# 1721 — Swapping Nodes In A Linked List

## Deskripsi

**Soal:** [1721. Swapping Nodes In A Linked List](https://leetcode.com/problems/swapping-nodes-in-a-linked-list/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** LIS (Longest Increasing Subsequence)

**Fungsi Solusi:** `func createList(arr []int) *ListNode`

## Solusi Go

```go
package main

// LeetCode #1721: Swapping Nodes in a Linked List
// https://leetcode.com/problems/swapping-nodes-in-a-linked-list/
// Difficulty: Medium (categorized as Hard in this repo)

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func createList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	head := &ListNode{Val: arr[0]}
	curr := head
	for i := 1; i < len(arr); i++ {
		curr.Next = &ListNode{Val: arr[i]}
		curr = curr.Next
	}
	return head
}

func printList(head *ListNode) {
	fmt.Print("[")
	for curr := head; curr != nil; curr = curr.Next {
		if curr != head {
			fmt.Print(",")
		}
		fmt.Print(curr.Val)
	}
	fmt.Print("]")
}

func main() {
	// Test 1
	list1 := createList([]int{1, 2, 3, 4, 5})
	fmt.Printf("Test 1 - Input: [1,2,3,4,5], k=2\n")
	printList(list1)
	fmt.Println()
	result1 := swapNodes(list1, 2)
	printList(result1)
	fmt.Println(" (Expected: [1,4,3,2,5])\n")

	// Test 2
	list2 := createList([]int{7, 9, 6, 6, 7, 8, 3, 0, 9, 5})
	fmt.Printf("Test 2 - Input: [7,9,6,6,7,8,3,0,9,5], k=5\n")
	result2 := swapNodes(list2, 5)
	printList(result2)
	fmt.Println(" (Expected: [7,9,6,6,8,7,3,0,9,5])\n")

	// Test 3
	list3 := createList([]int{1})
	fmt.Printf("Test 3 - Input: [1], k=1\n")
	result3 := swapNodes(list3, 1)
	printList(result3)
	fmt.Println(" (Expected: [1])")
}

func swapNodes(head *ListNode, k int) *ListNode {
	// Find k-th node from the beginning
	first := head
	for i := 1; i < k; i++ {
		first = first.Next
	}

	// Find k-th node from the end using two-pointer technique
	second := head
	curr := first
	for curr.Next != nil {
		curr = curr.Next
		second = second.Next
	}

	// Swap values (not nodes)
	first.Val, second.Val = second.Val, first.Val
	return head
}
```
