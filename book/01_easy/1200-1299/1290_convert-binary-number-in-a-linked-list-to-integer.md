# 1290 — Convert Binary Number In A Linked List To Integer

## Deskripsi

**Soal:** [1290. Convert Binary Number In A Linked List To Integer](https://leetcode.com/problems/convert-binary-number-in-a-linked-list-to-integer/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** LIS (Longest Increasing Subsequence)

## Solusi Go

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
