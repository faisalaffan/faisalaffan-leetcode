# 0382 — Linked List Random Node

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func Constructor(head *ListNode) Solution
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(1) for init, O(n) for getRandom  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #382: Linked List Random Node
// https://leetcode.com/problems/linked-list-random-node/
// Difficulty: Medium
// Time: O(1) for init, O(n) for getRandom | Space: O(1)

import (
	"fmt"
	"math/rand"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	head *ListNode
}

func Constructor(head *ListNode) Solution {
	return Solution{head: head}
}

// Reservoir sampling: O(n), uniform probability
func (s *Solution) GetRandom() int {
	result := s.head.Val
	node := s.head.Next
	i := 1
	for node != nil {
		i++
		if rand.Intn(i) == 0 {
			result = node.Val
		}
		node = node.Next
	}
	return result
}

func main() {
	// Test case: 1->2->3
	head := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	sol := Constructor(head)

	// Run multiple times to show randomness
	counts := map[int]int{}
	for i := 0; i < 30000; i++ {
		counts[sol.GetRandom()]++
	}
	fmt.Println("Counts:", counts)
	// Expected: roughly 10000 each
}
```
