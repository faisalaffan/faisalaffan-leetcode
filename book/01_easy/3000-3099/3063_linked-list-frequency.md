# 3063 — Linked List Frequency

## Deskripsi

**Soal:** [3063. Linked List Frequency](https://leetcode.com/problems/linked-list-frequency/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** LIS (Longest Increasing Subsequence)

## Solusi Go

```go
package main

// LeetCode #3063: Linked List Frequency
// https://leetcode.com/problems/linked-list-frequency/
// Difficulty: Easy [Paid]

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// LeetCode name: frequencies
	head := &ListNode{1, &ListNode{2, &ListNode{2, &ListNode{3, &ListNode{3, &ListNode{3, nil}}}}}}
	fmt.Println(LinkedListFrequency(head)) // map[1:1 2:2 3:3]
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: frequencies
func LinkedListFrequency(head *ListNode) map[int]int {
  // Membuat map untuk pencarian O(1): key → value
	freq := make(map[int]int)
	curr := head
	for curr != nil {
		freq[curr.Val]++
		curr = curr.Next
	}
	return freq
}
```
