package main

// LeetCode #2402: Meeting Rooms III
// https://leetcode.com/problems/meeting-rooms-iii/
// Difficulty: Hard
//
// You have n meeting rooms numbered 0 to n-1. Given meetings[start_i, end_i),
// assign each meeting to the smallest-numbered available room. If no room is
// available, delay the meeting until a room is free (keeping duration same).
// Return the room that hosts the most meetings.
//
// Approach: Use two min-heaps:
//   - available: room indices (min-heap by room number)
//   - busy: (endTime, roomIndex) (min-heap by endTime)
// Process meetings in sorted order by start time. For each meeting, release
// all busy rooms that are now free. Assign the meeting to the smallest
// available room. If no room is free, use the earliest-freed room.

import (
	"container/heap"
	"fmt"
	"sort"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type Busy struct {
	endTime int
	room    int
}

type BusyHeap []Busy

func (h BusyHeap) Len() int           { return len(h) }
func (h BusyHeap) Less(i, j int) bool { return h[i].endTime < h[j].endTime }
func (h BusyHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *BusyHeap) Push(x interface{}) {
	*h = append(*h, x.(Busy))
}

func (h *BusyHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func mostBooked(n int, meetings [][]int) int {
	// Sort meetings by start time
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	avail := &IntHeap{}
	for i := 0; i < n; i++ {
		heap.Push(avail, i)
	}

	busy := &BusyHeap{}
	count := make([]int, n)

	for _, m := range meetings {
		start, end := m[0], m[1]
		duration := end - start

		// Release rooms that have finished by now
		for busy.Len() > 0 && (*busy)[0].endTime <= start {
			b := heap.Pop(busy).(Busy)
			heap.Push(avail, b.room)
		}

		var room int
		if avail.Len() > 0 {
			// Room available, use smallest-numbered
			room = heap.Pop(avail).(int)
		} else {
			// No room available, pick the soonest free room and delay
			b := heap.Pop(busy).(Busy)
			room = b.room
			// Meeting is delayed; duration stays same
			start = b.endTime
		}

		count[room]++
		heap.Push(busy, Busy{endTime: start + duration, room: room})
	}

	// Find room with max count, smallest number in case of tie
	maxRoom, maxCount := 0, count[0]
	for i := 1; i < n; i++ {
		if count[i] > maxCount {
			maxCount = count[i]
			maxRoom = i
		}
	}

	return maxRoom
}

func main() {
	// Example 1
	n1 := 2
	meetings1 := [][]int{{0, 10}, {1, 5}, {2, 7}, {3, 4}}
	fmt.Println(mostBooked(n1, meetings1))

	// Example 2
	n2 := 3
	meetings2 := [][]int{{1, 20}, {2, 10}, {3, 5}, {4, 9}, {6, 8}}
	fmt.Println(mostBooked(n2, meetings2))

	// Single room
	n3 := 1
	meetings3 := [][]int{{0, 5}, {2, 3}, {4, 7}}
	fmt.Println(mostBooked(n3, meetings3))

	// All meetings overlap
	n4 := 2
	meetings4 := [][]int{{0, 10}, {0, 5}, {0, 7}}
	fmt.Println(mostBooked(n4, meetings4))
}
