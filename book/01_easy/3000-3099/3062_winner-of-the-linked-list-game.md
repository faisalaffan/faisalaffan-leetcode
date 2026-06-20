# 3062 — Winner Of The Linked List Game

## Deskripsi

**Soal:** [3062. Winner Of The Linked List Game](https://leetcode.com/problems/winner-of-the-linked-list-game/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** LIS (Longest Increasing Subsequence)

## Solusi Go

```go
package main

// LeetCode #3062: Winner of the Linked List Game
// https://leetcode.com/problems/winner-of-the-linked-list-game/
// Difficulty: Easy [Paid]
//
// Note: This is a premium problem. We implement the equivalent logic.
// The game: pairs of consecutive nodes. Even score increases if first > second,
// odd score increases if second > first. Return "Even" or "Odd".

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// LeetCode name: gameResult
	// Create: 2 -> 1 -> 4 -> 5 -> 2 -> 9
	head := &ListNode{2, &ListNode{1, &ListNode{4, &ListNode{5, &ListNode{2, &ListNode{9, nil}}}}}}
	fmt.Println(WinnerOfTheLinkedListGame(head)) // Odd

	head2 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	fmt.Println(WinnerOfTheLinkedListGame(head2)) // Odd
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: gameResult
func WinnerOfTheLinkedListGame(head *ListNode) string {
	evenScore := 0
	oddScore := 0
	curr := head
	for curr != nil && curr.Next != nil {
		first, second := curr.Val, curr.Next.Val
		if first > second {
			evenScore++
		} else if second > first {
			oddScore++
		}
		curr = curr.Next.Next
	}
	if evenScore > oddScore {
		return "Even"
	}
	if oddScore > evenScore {
		return "Odd"
	}
	return "Tie"
}
```
