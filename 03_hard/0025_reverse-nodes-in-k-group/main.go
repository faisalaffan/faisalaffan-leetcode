package main

// LeetCode #25: Reverse Nodes in k-Group
// https://leetcode.com/problems/reverse-nodes-in-k-group/
// Difficulty: Hard

import "fmt"

// ListNode represents a node in a singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// reverseKGroup reverses the nodes of a linked list k at a time.
//
// Complexity: O(n) time, O(1) space
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k <= 1 {
		return head
	}

	// Count nodes to know where to stop
	count := 0
	curr := head
	for curr != nil {
		count++
		curr = curr.Next
	}

	dummy := &ListNode{Next: head}
	prev := dummy

	for count >= k {
		// Reverse k nodes starting from prev.Next
		start := prev.Next
		curr = start
		var prevNode *ListNode
		for i := 0; i < k; i++ {
			nextTemp := curr.Next
			curr.Next = prevNode
			prevNode = curr
			curr = nextTemp
		}
		// Connect the reversed segment
		start.Next = curr
		prev.Next = prevNode
		prev = start
		count -= k
	}

	return dummy.Next
}

// Helper to build a linked list from a slice.
func buildList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0]}
	curr := head
	for i := 1; i < len(vals); i++ {
		curr.Next = &ListNode{Val: vals[i]}
		curr = curr.Next
	}
	return head
}

// Helper to convert a linked list to a slice.
func listToSlice(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

func main() {
	// Test case from LeetCode: [1,2,3,4,5], k=2 -> [2,1,4,3,5]
	list1 := buildList([]int{1, 2, 3, 4, 5})
	result1 := reverseKGroup(list1, 2)
	fmt.Println("k=2:", listToSlice(result1))

	// Test case from LeetCode: [1,2,3,4,5], k=3 -> [3,2,1,4,5]
	list2 := buildList([]int{1, 2, 3, 4, 5})
	result2 := reverseKGroup(list2, 3)
	fmt.Println("k=3:", listToSlice(result2))

	// Edge case: k=1 (no change)
	list3 := buildList([]int{1, 2, 3})
	result3 := reverseKGroup(list3, 1)
	fmt.Println("k=1:", listToSlice(result3))
}
