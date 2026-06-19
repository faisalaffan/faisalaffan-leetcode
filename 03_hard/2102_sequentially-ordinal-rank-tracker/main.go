package main

// LeetCode #2102: Sequentially Ordinal Rank Tracker
// https://leetcode.com/problems/sequentially-ordinal-rank-tracker/
// Difficulty: Hard
//
// Approach: Two heaps (min-heap + max-heap).
// - low (min-heap by score desc, name asc): contains top k items, root = k-th best
// - high (max-heap by score asc, name desc): contains items beyond top k, root = (k+1)-th best
// Get() increments the query count and returns the k-th best item.

import (
	"container/heap"
	"fmt"
)

type Location struct {
	name  string
	score int
}

// LowHeap is a min-heap: root is the "worst" among the top-k items.
// Comparison: score ascending, then name ASC (earlier alphabetically = worse).
type LowHeap []Location

func (h LowHeap) Len() int      { return len(h) }
func (h LowHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h LowHeap) Less(i, j int) bool {
	if h[i].score != h[j].score {
		return h[i].score < h[j].score // lower score = worse = root
	}
	return h[i].name < h[j].name // earlier name = worse = root
}
func (h *LowHeap) Push(x any)   { *h = append(*h, x.(Location)) }
func (h *LowHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// HighHeap is a max-heap: root is the "best" among the leftovers.
// Comparison: score descending, then name DESC (later alphabetically = better).
type HighHeap []Location

func (h HighHeap) Len() int      { return len(h) }
func (h HighHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h HighHeap) Less(i, j int) bool {
	if h[i].score != h[j].score {
		return h[i].score > h[j].score // higher score = better = root
	}
	return h[i].name > h[j].name // later name = better = root
}
func (h *HighHeap) Push(x any)   { *h = append(*h, x.(Location)) }
func (h *HighHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type SORTracker struct {
	low     LowHeap
	high    HighHeap
	queries int
}

func Constructor() SORTracker {
	return SORTracker{
		low:  LowHeap{},
		high: HighHeap{},
	}
}

func (t *SORTracker) Add(name string, score int) {
	heap.Push(&t.low, Location{name, score})
	if len(t.low) > t.queries {
		heap.Push(&t.high, heap.Pop(&t.low))
	}
}

func (t *SORTracker) Get() string {
	t.queries++
	for len(t.low) < t.queries {
		heap.Push(&t.low, heap.Pop(&t.high))
	}
	return t.low[0].name
}

func main() {
	// Test with example from problem
	tracker := Constructor()
	tracker.Add("bradford", 2)
	tracker.Add("branford", 3)
	fmt.Printf("Get() = %s (expected branford)\n", tracker.Get())
	tracker.Add("alps", 2)
	fmt.Printf("Get() = %s (expected alps)\n", tracker.Get())
	tracker.Add("orland", 2)
	fmt.Printf("Get() = %s (expected bradford)\n", tracker.Get())
	tracker.Add("orlando", 3)
	fmt.Printf("Get() = %s (expected orlando)\n", tracker.Get())
	tracker.Add("alpine", 2)
	fmt.Printf("Get() = %s (expected orland)\n", tracker.Get())
	fmt.Printf("Get() = %s (expected alpine)\n", tracker.Get())
}
