# 1474 — Delete N Nodes After M Nodes Of A Linked List

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func deleteNodes(head *ListNode, m int, n int) *ListNode

import "fmt"

type ListNode struct
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1474: Delete N Nodes After M Nodes of a Linked List
// https://leetcode.com/problems/delete-n-nodes-after-m-nodes-of-a-linked-list/
// Difficulty: Easy [Paid]
//
// LeetCode submission: func deleteNodes(head *ListNode, m int, n int) *ListNode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// List: 1->2->3->4->5->6->7->8->9->10->11, m=2, n=3
	// Expected: 1->2->6->7->11
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 7, Next: &ListNode{Val: 8, Next: &ListNode{Val: 9, Next: &ListNode{Val: 10, Next: &ListNode{Val: 11}}}}}}}}}}}
	result := DeleteNNodesAfterMNodesOfALinkedList(head, 2, 3)
	for result != nil {
		fmt.Print(result.Val, " ")
		result = result.Next
	}
	fmt.Println() // 1 2 6 7 11
}

// Time: O(n), Space: O(1)
func DeleteNNodesAfterMNodesOfALinkedList(head *ListNode, m int, n int) *ListNode {
	cur := head
	for cur != nil {
		// Skip m nodes
		for i := 1; i < m && cur != nil; i++ {
			cur = cur.Next
		}
		if cur == nil {
			break
		}
		// Delete next n nodes
		temp := cur.Next
		for i := 0; i < n && temp != nil; i++ {
			temp = temp.Next
		}
		cur.Next = temp
		cur = temp
	}
	return head
}
```
