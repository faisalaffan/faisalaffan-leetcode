# 0382 — Linked List Random Node

## Deskripsi

**Soal:** [0382. Linked List Random Node](https://leetcode.com/problems/linked-list-random-node/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(1) for init, O(n) for getRandom  
**Kompleksitas Ruang:** O(1)

**Algoritma:** LIS (Longest Increasing Subsequence)

**Fungsi Solusi:** `func Constructor(head *ListNode) Solution`

## Solusi Go

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
