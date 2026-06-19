package main

// LeetCode #23: Merge k Sorted Lists
// https://leetcode.com/problems/merge-k-sorted-lists/
// Difficulty: Hard

import (
	"container/heap"
	"fmt"
)

// ListNode represents a node in a singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// MinHeap implements heap.Interface for ListNode pointers.
type MinHeap []*ListNode

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].Val < h[j].Val }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(*ListNode)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// mergeKLists merges k sorted linked lists using a min-heap.
//
// Complexity: O(N log k) time, O(k) space where N = total nodes, k = number of lists
func mergeKLists(lists []*ListNode) *ListNode {
	h := &MinHeap{}
	heap.Init(h)

	// push the head of each non-empty list into the heap
	for _, list := range lists {
		if list != nil {
			heap.Push(h, list)
		}
	}

	dummy := &ListNode{}
	curr := dummy

	for h.Len() > 0 {
		node := heap.Pop(h).(*ListNode)
		curr.Next = node
		curr = curr.Next
		if node.Next != nil {
			heap.Push(h, node.Next)
		}
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
	// Test case from LeetCode: [[1,4,5],[1,3,4],[2,6]] -> [1,1,2,3,4,4,5,6]
	lists := []*ListNode{
		buildList([]int{1, 4, 5}),
		buildList([]int{1, 3, 4}),
		buildList([]int{2, 6}),
	}
	result := mergeKLists(lists)
	fmt.Println("Merged list:", listToSlice(result))

	// Edge case: empty lists
	var emptyLists []*ListNode
	result2 := mergeKLists(emptyLists)
	fmt.Println("Empty input:", listToSlice(result2))

	// Edge case: single list
	single := []*ListNode{buildList([]int{1})}
	result3 := mergeKLists(single)
	fmt.Println("Single list:", listToSlice(result3))
}
