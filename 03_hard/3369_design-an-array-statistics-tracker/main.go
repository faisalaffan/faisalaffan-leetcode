package main

// LeetCode #3369: Design an Array Statistics Tracker
// https://leetcode.com/problems/design-an-array-statistics-tracker/
// Difficulty: Hard [Paid]
//
// Design a data structure that supports:
// - add(element): add an element to the tracker
// - getMin(): return the minimum element
// - getMax(): return the maximum element
// - getMedian(): return the median element
// - getMean(): return the mean (average) of all elements
// - getMode(): return the mode (most frequent element)
// - remove(element): remove one occurrence of the element
//
// Approach: Use heaps for median (two heaps), maps for frequency tracking.

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	// Test StatisticsTracker
	tracker := Constructor()
	tracker.AddElement(5)
	tracker.AddElement(3)
	tracker.AddElement(7)
	fmt.Println(tracker.GetMin())    // 3
	fmt.Println(tracker.GetMax())    // 7
	fmt.Println(tracker.GetMean())    // 5
	fmt.Println(tracker.GetMedian())  // 5
	fmt.Println(tracker.GetMode())    // 5 (or any, all appear once)

	tracker.AddElement(3)
	fmt.Println(tracker.GetMode())    // 3
	fmt.Println(tracker.GetMedian())  // 3

	tracker.RemoveElement(5)
	fmt.Println(tracker.GetMedian())  // 3

	// Edge: single element
	tracker2 := Constructor()
	tracker2.AddElement(10)
	fmt.Println(tracker2.GetMin())    // 10
	fmt.Println(tracker2.GetMax())    // 10
	fmt.Println(tracker2.GetMean())   // 10
	fmt.Println(tracker2.GetMedian()) // 10
	fmt.Println(tracker2.GetMode())   // 10
}

// IntHeap for min-heap and max-heap
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type MaxIntHeap []int

func (h MaxIntHeap) Len() int           { return len(h) }
func (h MaxIntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxIntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxIntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxIntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type StatisticsTracker struct {
	count map[int]int   // element -> count
	min   int
	max   int
	sum   int64
	size  int

	// For median
	low  *MaxIntHeap // max-heap for left half
	high *IntHeap    // min-heap for right half

	// For mode
	modeVal   int
	modeCount int
}

func Constructor() StatisticsTracker {
	low := &MaxIntHeap{}
	high := &IntHeap{}
	heap.Init(low)
	heap.Init(high)
	return StatisticsTracker{
		count:     make(map[int]int),
		min:       math.MaxInt32,
		max:       math.MinInt32,
		low:       low,
		high:      high,
		modeVal:   math.MaxInt32,
		modeCount: 0,
	}
}

func (st *StatisticsTracker) AddElement(val int) {
	st.count[val]++
	st.size++
	st.sum += int64(val)
	if val < st.min {
		st.min = val
	}
	if val > st.max {
		st.max = val
	}

	// Update mode
	c := st.count[val]
	if c > st.modeCount || (c == st.modeCount && val < st.modeVal) {
		st.modeCount = c
		st.modeVal = val
	}

	// Add to heaps for median
	if st.low.Len() == 0 || val <= (*st.low)[0] {
		heap.Push(st.low, val)
	} else {
		heap.Push(st.high, val)
	}

	// Rebalance
	if st.low.Len() > st.high.Len()+1 {
		heap.Push(st.high, heap.Pop(st.low))
	} else if st.high.Len() > st.low.Len() {
		heap.Push(st.low, heap.Pop(st.high))
	}
}

func (st *StatisticsTracker) RemoveElement(val int) {
	if st.count[val] <= 0 {
		return
	}
	st.count[val]--
	st.size--
	st.sum -= int64(val)

	if st.count[val] == 0 {
		delete(st.count, val)
		// Reset mode if needed
		if val == st.modeVal {
			st.modeVal = math.MaxInt32
			st.modeCount = 0
			for v, c := range st.count {
				if c > st.modeCount || (c == st.modeCount && v < st.modeVal) {
					st.modeCount = c
					st.modeVal = v
				}
			}
		}
	}

	// Update min/max
	if val == st.min {
		st.min = math.MaxInt32
		for v := range st.count {
			if v < st.min {
				st.min = v
			}
		}
	}
	if val == st.max {
		st.max = math.MinInt32
		for v := range st.count {
			if v > st.max {
				st.max = v
			}
		}
	}
}

func (st *StatisticsTracker) GetMin() int {
	return st.min
}

func (st *StatisticsTracker) GetMax() int {
	return st.max
}

func (st *StatisticsTracker) GetMean() int {
	if st.size == 0 {
		return 0
	}
	return int(st.sum / int64(st.size))
}

func (st *StatisticsTracker) GetMedian() int {
	if st.size == 0 {
		return 0
	}
	return (*st.low)[0]
}

func (st *StatisticsTracker) GetMode() int {
	if st.size == 0 {
		return 0
	}
	return st.modeVal
}
