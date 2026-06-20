# 2095 — Delete The Middle Node Of A Linked List

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func deleteMiddle(head *ListNode) *ListNode
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2095: Delete the Middle Node of a Linked List
// https://leetcode.com/problems/delete-the-middle-node-of-a-linked-list/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow, fast := head, head
	var prev *ListNode

	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Delete slow (middle) node
	prev.Next = slow.Next

	return head
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}

func main() {
	// Test case 1
	head1 := &ListNode{1, &ListNode{3, &ListNode{4, &ListNode{7, &ListNode{1, &ListNode{2, &ListNode{6, nil}}}}}}}
	fmt.Print("Test 1: ")
	printList(deleteMiddle(head1))
	// Expected: [1,3,4,1,2,6]

	// Test case 2
	head2 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	fmt.Print("Test 2: ")
	printList(deleteMiddle(head2))
	// Expected: [1,2,4]

	// Test case 3
	head3 := &ListNode{1, nil}
	fmt.Print("Test 3: ")
	printList(deleteMiddle(head3))
	// Expected: []
}
```
