# 1171 — Remove Zero Sum Consecutive Nodes From Linked List

## Deskripsi

**Soal:** [1171. Remove Zero Sum Consecutive Nodes From Linked List](https://leetcode.com/problems/remove-zero-sum-consecutive-nodes-from-linked-list/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** HashMap (tabel pencarian O(1)), Prefix Sum (jumlah kumulatif), LIS (Longest Increasing Subsequence)

**Fungsi Solusi:** `func removeZeroSumSublists(head *ListNode) *ListNode`

## Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #1171: Remove Zero Sum Consecutive Nodes from Linked List
// https://leetcode.com/problems/remove-zero-sum-consecutive-nodes-from-linked-list/
// Difficulty: Medium

// Use prefix sum with a hashmap. If same prefix sum repeats,
// the segment between has sum 0, so remove it.

// Time: O(n)
// Space: O(n)

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeZeroSumSublists(head *ListNode) *ListNode {
	dummy := &ListNode{0, head}
	prefix := 0
  // Membuat map untuk pencarian O(1): key → value
	seen := make(map[int]*ListNode)
	seen[0] = dummy

	for cur := dummy; cur != nil; cur = cur.Next {
		prefix += cur.Val
		seen[prefix] = cur
	}

	prefix = 0
	for cur := dummy; cur != nil; cur = cur.Next {
		prefix += cur.Val
		cur.Next = seen[prefix].Next
	}

	return dummy.Next
}

func listToSlice(head *ListNode) []int {
  // Membuat slice untuk menyimpan hasil
	result := make([]int, 0)
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

func sliceToList(nums []int) *ListNode {
  // Edge case: input kosong
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{nums[0], nil}
	cur := head
	for i := 1; i < len(nums); i++ {
		cur.Next = &ListNode{nums[i], nil}
		cur = cur.Next
	}
	return head
}

func main() {
	l1 := sliceToList([]int{1, 2, -3, 3, 1})
	fmt.Printf("%v (expected: [3 1] or [1 2 1])\n", listToSlice(removeZeroSumSublists(l1)))

	l2 := sliceToList([]int{1, 2, 3, -3, 4})
	fmt.Printf("%v (expected: [1 2 4])\n", listToSlice(removeZeroSumSublists(l2)))

	l3 := sliceToList([]int{1, -1})
	fmt.Printf("%v (expected: [])\n", listToSlice(removeZeroSumSublists(l3)))
}
```
