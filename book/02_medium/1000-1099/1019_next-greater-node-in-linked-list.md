# 1019 — Next Greater Node In Linked List

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan linked list. Tugasmu traversing atau memanipulasi list.

**Cara berpikir:** Traverse dari head. Fast/slow pointer untuk deteksi siklus/cari tengah. Dummy node mempermudah operasi di head.

**Fungsi Solusi:** `func nextLargerNodes(head *ListNode) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Stack, Monotonic Stack

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Stack** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1019: Next Greater Node In Linked List
// https://leetcode.com/problems/next-greater-node-in-linked-list/
// Difficulty: Medium
//
// Approach: Convert linked list to array, then use monotonic stack
// Time: O(n)
// Space: O(n)

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// [2,1,5]
	head := &ListNode{2, &ListNode{1, &ListNode{5, nil}}}
	fmt.Println(nextLargerNodes(head)) // [5,5,0]

	// [2,7,4,3,5]
	head2 := &ListNode{2, &ListNode{7, &ListNode{4, &ListNode{3, &ListNode{5, nil}}}}}
	fmt.Println(nextLargerNodes(head2)) // [7,0,5,5,0]
}

func nextLargerNodes(head *ListNode) []int {
	// Convert to array
  // Alokasi slice
	vals := make([]int, 0)
	for cur := head; cur != nil; cur = cur.Next {
		vals = append(vals, cur.Val)
	}

  // Alokasi slice
	result := make([]int, len(vals))
  // Alokasi slice
	stack := make([]int, 0) // indices

	for i, v := range vals {
		for len(stack) > 0 && vals[stack[len(stack)-1]] < v {
			result[stack[len(stack)-1]] = v
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	// Remaining in stack have no next greater (already zero-initialized)
	return result
}
```
