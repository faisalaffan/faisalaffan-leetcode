package main

// LeetCode #3885: Design Event Manager
// https://leetcode.com/problems/design-event-manager/
// Difficulty: Medium
// Time: O(N log N) init, O(log N) per operation | Space: O(N)
// Approach: Use max-heap (priority queue) on (-priority, eventId) + map for event priorities.

import (
	"container/heap"
	"fmt"
)

type Event struct {
	priority int
	eventId  int
}

type MaxHeap []Event

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool {
	if h[i].priority != h[j].priority {
		return h[i].priority > h[j].priority
	}
	return h[i].eventId < h[j].eventId
}
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(Event)) }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type EventManager struct {
	pq       *MaxHeap
	priorities map[int]int // eventId -> priority
}

func Constructor(events [][]int) EventManager {
	pq := &MaxHeap{}
	heap.Init(pq)
	pm := make(map[int]int)
	for _, e := range events {
		id, pri := e[0], e[1]
		pm[id] = pri
		heap.Push(pq, Event{priority: pri, eventId: id})
	}
	return EventManager{pq: pq, priorities: pm}
}

func (em *EventManager) UpdatePriority(eventId int, newPriority int) {
	em.priorities[eventId] = newPriority
	heap.Push(em.pq, Event{priority: newPriority, eventId: eventId})
}

func (em *EventManager) PollHighest() int {
	for em.pq.Len() > 0 {
		top := (*em.pq)[0]
		if pri, ok := em.priorities[top.eventId]; ok && pri == top.priority {
			heap.Pop(em.pq)
			delete(em.priorities, top.eventId)
			return top.eventId
		}
		heap.Pop(em.pq) // stale entry
	}
	return -1
}

func main() {
	// Example 1
	em := Constructor([][]int{{5, 7}, {2, 7}, {9, 4}})
	fmt.Println(em.PollHighest()) // Expected: 2
	em.UpdatePriority(9, 7)
	fmt.Println(em.PollHighest()) // Expected: 5
	fmt.Println(em.PollHighest()) // Expected: 9

	// Example 2
	em2 := Constructor([][]int{{4, 1}, {7, 2}})
	fmt.Println(em2.PollHighest()) // Expected: 7
	fmt.Println(em2.PollHighest()) // Expected: 4
	fmt.Println(em2.PollHighest()) // Expected: -1
}
