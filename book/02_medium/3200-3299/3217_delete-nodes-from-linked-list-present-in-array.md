# 3217 — Delete Nodes From Linked List Present In Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan linked list. Tugasmu traversing atau memanipulasi list.

**Cara berpikir:** Traverse dari head. Fast/slow pointer untuk deteksi siklus/cari tengah. Dummy node mempermudah operasi di head.

**Fungsi Solusi:** `func modifiedList(nums []int, head *ListNode) *ListNode`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(m)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3217: Delete Nodes From Linked List Present in Array
// https://leetcode.com/problems/delete-nodes-from-linked-list-present-in-array/
// Difficulty: Medium
// Time: O(n) | Space: O(m)

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func modifiedList(nums []int, head *ListNode) *ListNode {
  // HashMap: O(1) lookup
	del := make(map[int]bool)
	for _, v := range nums {
		del[v] = true
	}

	dummy := &ListNode{Next: head}
	cur := dummy

	for cur.Next != nil {
		if del[cur.Next.Val] {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}

func listToSlice(head *ListNode) []int {
	var ans []int
	for head != nil {
		ans = append(ans, head.Val)
		head = head.Next
	}
	return ans
}

func main() {
	head1 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	fmt.Println(listToSlice(modifiedList([]int{1, 2, 3}, head1))) // Expected: [4, 5]

	head2 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	fmt.Println(listToSlice(modifiedList([]int{5}, head2))) // Expected: [1, 2, 3, 4]
}
```
