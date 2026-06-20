# 2058 — Find The Minimum And Maximum Number Of Nodes Between Critical Points

## Deskripsi

**Soal:** [2058. Find The Minimum And Maximum Number Of Nodes Between Critical Points](https://leetcode.com/problems/find-the-minimum-and-maximum-number-of-nodes-between-critical-points/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** LIS (Longest Increasing Subsequence)

**Fungsi Solusi:** `func nodesBetweenCriticalPoints(head *ListNode) []int`

## Solusi Go

```go
package main

// LeetCode #2058: Find the Minimum and Maximum Number of Nodes Between Critical Points
// https://leetcode.com/problems/find-the-minimum-and-maximum-number-of-nodes-between-critical-points/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func nodesBetweenCriticalPoints(head *ListNode) []int {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return []int{-1, -1}
	}

	prev := head
	curr := head.Next
	pos := 1
	firstCritical := -1
	lastCritical := -1
	minDist := int(1e9)

	for curr.Next != nil {
		if (curr.Val > prev.Val && curr.Val > curr.Next.Val) ||
			(curr.Val < prev.Val && curr.Val < curr.Next.Val) {
			if firstCritical == -1 {
				firstCritical = pos
			} else {
				dist := pos - lastCritical
				if dist < minDist {
					minDist = dist
				}
			}
			lastCritical = pos
		}
		prev = curr
		curr = curr.Next
		pos++
	}

	if firstCritical == lastCritical || firstCritical == -1 {
		return []int{-1, -1}
	}

	return []int{minDist, lastCritical - firstCritical}
}

func main() {
	// Test case 1
	head1 := &ListNode{3, &ListNode{1, nil}}
	fmt.Println("Test 1:", nodesBetweenCriticalPoints(head1))
	// Expected: [-1, -1] (not enough nodes)

	// Test case 2
	head2 := &ListNode{5, &ListNode{3, &ListNode{1, &ListNode{2, &ListNode{5, &ListNode{1, &ListNode{2, nil}}}}}}}
	fmt.Println("Test 2:", nodesBetweenCriticalPoints(head2))
	// Expected: [1, 3]

	// Test case 3
	head3 := &ListNode{1, &ListNode{3, &ListNode{2, &ListNode{2, &ListNode{3, &ListNode{2, &ListNode{2, &ListNode{2, &ListNode{7, nil}}}}}}}}}
	fmt.Println("Test 3:", nodesBetweenCriticalPoints(head3))
	// Expected: [3, 3]
}
```
