# 2487 — Remove Nodes From Linked List

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func removeNodes(head *ListNode) *ListNode
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2487: Remove Nodes From Linked List
// https://leetcode.com/problems/remove-nodes-from-linked-list/
// Difficulty: Medium
// Time: O(n) | Space: O(n)
// Reverse list, track max so far, keep nodes >= max.

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// 5 -> 2 -> 13 -> 3 -> 8
	head := &ListNode{5, &ListNode{2, &ListNode{13, &ListNode{3, &ListNode{8, nil}}}}}
	res := removeNodes(head)
	for res != nil {
		fmt.Print(res.Val, " ") // 13 8
		res = res.Next
	}
	fmt.Println()
}

func removeNodes(head *ListNode) *ListNode {
	// Reverse
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	// Keep nodes >= max so far
	dummy := &ListNode{Next: prev}
	cur = dummy.Next
	maxSoFar := cur.Val
	for cur != nil && cur.Next != nil {
		if cur.Next.Val < maxSoFar {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
			maxSoFar = cur.Val
		}
	}

	// Reverse back
	var prev2 *ListNode
	cur = dummy.Next
	for cur != nil {
		next := cur.Next
		cur.Next = prev2
		prev2 = cur
		cur = next
	}
	return prev2
}
```
