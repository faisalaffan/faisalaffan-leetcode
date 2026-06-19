package main

// LeetCode #817: Linked List Components
// https://leetcode.com/problems/linked-list-components/
// Difficulty: Medium

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// Test case 1: [0,1,2,3], nums=[0,1,3] -> 2
	head1 := &ListNode{0, &ListNode{1, &ListNode{2, &ListNode{3, nil}}}}
	fmt.Println(LinkedListComponents(head1, []int{0, 1, 3}))

	// Test case 2: [0,1,2,3,4], nums=[0,3,1,4] -> 2
	head2 := &ListNode{0, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}}
	fmt.Println(LinkedListComponents(head2, []int{0, 3, 1, 4}))

	// Test case 3: nil list
	fmt.Println(LinkedListComponents(nil, []int{1}))
}

// Time: O(n) | Space: O(m) where m = len(nums)
func LinkedListComponents(head *ListNode, nums []int) int {
	set := make(map[int]bool)
	for _, v := range nums {
		set[v] = true
	}

	count := 0
	inComponent := false
	for cur := head; cur != nil; cur = cur.Next {
		if set[cur.Val] {
			if !inComponent {
				count++
				inComponent = true
			}
		} else {
			inComponent = false
		}
	}

	return count
}
