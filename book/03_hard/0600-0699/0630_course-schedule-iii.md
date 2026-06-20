# 0630 — Course Schedule Iii

## Deskripsi

**Soal:** [0630. Course Schedule Iii](https://leetcode.com/problems/course-schedule-iii/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Heap (priority queue)

## Solusi Go

```go
package main

// LeetCode #630: Course Schedule III
// https://leetcode.com/problems/course-schedule-iii/
// Difficulty: Hard

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	// Test cases
	testCases := []struct {
		courses [][]int
		want    int
	}{
		{[][]int{{100, 200}, {200, 1300}, {1000, 1250}, {2000, 3200}}, 3},
		{[][]int{{1, 2}}, 1},
		{[][]int{{3, 2}, {4, 3}}, 0},
		{[][]int{{5, 5}, {4, 6}, {2, 6}}, 2},
		{[][]int{{5, 6}, {3, 4}, {1, 2}}, 2},
		{[][]int{{7, 17}, {3, 12}, {10, 20}, {9, 10}, {5, 20}, {3, 19}}, 4},
		{[][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 10}}, 2},
		{[][]int{}, 0},
	}

	for _, tc := range testCases {
		got := scheduleCourse(tc.courses)
		status := "PASS"
		if got != tc.want {
			status = "FAIL"
		}
		fmt.Printf("%s: scheduleCourse(%v) = %d (want %d)\n", status, tc.courses, got, tc.want)
	}
}

// Max-heap for course durations
type maxHeap []int

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func scheduleCourse(courses [][]int) int {
	// Sort by lastDay (ascending)
	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1]
	})

	h := &maxHeap{}
	heap.Init(h)
	totalDuration := 0

	for _, course := range courses {
		duration, lastDay := course[0], course[1]
		heap.Push(h, duration)
		totalDuration += duration

		// If we exceed lastDay, drop the longest course taken
		if totalDuration > lastDay {
			maxDuration := heap.Pop(h).(int)
			totalDuration -= maxDuration
		}
	}

	return h.Len()
}
```
