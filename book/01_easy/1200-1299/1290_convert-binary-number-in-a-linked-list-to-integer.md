# 1290 — Convert Binary Number In A Linked List To Integer

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan linked list. Tugasmu traversing atau memanipulasi list.

**Cara berpikir:** Traverse dari head. Fast/slow pointer untuk deteksi siklus/cari tengah. Dummy node mempermudah operasi di head.

**Fungsi Solusi:** `func getDecimalValue(head *ListNode) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1290: Convert Binary Number in a Linked List to Integer
// https://leetcode.com/problems/convert-binary-number-in-a-linked-list-to-integer/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// 1 -> 0 -> 1 = 5
	head := &ListNode{1, &ListNode{0, &ListNode{1, nil}}}
	fmt.Println(getDecimalValue(head)) // 5

	// 0 -> 0 = 0
	fmt.Println(getDecimalValue(&ListNode{0, &ListNode{0, nil}})) // 0
}

// LeetCode submission: getDecimalValue
func getDecimalValue(head *ListNode) int {
	ans := 0
	for head != nil {
		ans = ans*2 + head.Val
		head = head.Next
	}
	return ans
}
```
