# 0234 — Palindrome Linked List

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu adalah memeriksa apakah string tersebut palindrome — dibaca sama dari depan dan belakang. Abaikan non-alfanumerik dan case.

**Cara berpikir:** Two Pointer — kiri dan kanan. Skip non-alfanumerik. Bandingkan.

**Fungsi Solusi:** `func IsPalindrome(head *ListNode) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #234: Palindrome Linked List
// https://leetcode.com/problems/palindrome-linked-list/
// Difficulty: Easy

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time: O(n) | Space: O(1)
func IsPalindrome(head *ListNode) bool {
	reverse := func(head *ListNode) *ListNode {
		var prev *ListNode
		for head != nil {
			next := head.Next
			head.Next = prev
			prev = head
			head = next
		}
		return prev
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	second := reverse(slow)
	first := head
	for second != nil {
		if first.Val != second.Val {
			return false
		}
		first = first.Next
		second = second.Next
	}
	return true
}

func main() {
	l1 := &ListNode{1, &ListNode{2, &ListNode{2, &ListNode{1, nil}}}}
	fmt.Println(IsPalindrome(l1))
	l2 := &ListNode{1, &ListNode{2, nil}}
	fmt.Println(IsPalindrome(l2))
}
```
