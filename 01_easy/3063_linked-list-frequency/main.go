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
	freq := make(map[int]int)
	curr := head
	for curr != nil {
		freq[curr.Val]++
		curr = curr.Next
	}
	return freq
}
