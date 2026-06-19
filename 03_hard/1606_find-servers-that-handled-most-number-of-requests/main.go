package main

// LeetCode #1606: Find Servers That Handled Most Number of Requests
// https://leetcode.com/problems/find-servers-that-handled-most-number-of-requests/
// Difficulty: Hard

import (
	"container/heap"
	"fmt"
	"sort"
)

// --- Fenwick Tree (BIT) for ordered set of available servers ---

type BIT struct {
	tree []int
	n    int
}

func newBIT(n int) *BIT {
	b := &BIT{tree: make([]int, n+1), n: n}
	for i := 1; i <= n; i++ {
		b.add(i, 1)
	}
	return b
}

func (b *BIT) add(idx, val int) {
	for i := idx; i <= b.n; i += i & -i {
		b.tree[i] += val
	}
}

func (b *BIT) sum(idx int) int {
	s := 0
	for i := idx; i > 0; i -= i & -i {
		s += b.tree[i]
	}
	return s
}

// find smallest idx with prefix sum >= target (1-indexed)
func (b *BIT) find(target int) int {
	idx := 0
	bitMask := 1
	for bitMask <= b.n {
		bitMask <<= 1
	}
	bitMask >>= 1

	for bitMask > 0 {
		t := idx + bitMask
		if t <= b.n && b.tree[t] < target {
			target -= b.tree[t]
			idx = t
		}
		bitMask >>= 1
	}
	return idx + 1
}

// ceil returns smallest available server >= x, or wraps to first available
func (b *BIT) ceil(x int) int {
	total := b.sum(b.n)
	if total == 0 {
		return -1
	}
	before := b.sum(x)
	if before < total {
		return b.find(before + 1) - 1 // convert to 0-indexed
	}
	return b.find(1) - 1 // wrap around
}

func (b *BIT) remove(x int) {
	b.add(x+1, -1)
}

func (b *BIT) addServer(x int) {
	b.add(x+1, 1)
}

// --- Min Heap for busy servers ---

type busyItem struct {
	endTime int
	index   int
}

type busyHeap []*busyItem

func (h busyHeap) Len() int            { return len(h) }
func (h busyHeap) Less(i, j int) bool  { return h[i].endTime < h[j].endTime }
func (h busyHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *busyHeap) Push(x interface{}) { *h = append(*h, x.(*busyItem)) }
func (h *busyHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func (h *busyHeap) Peek() *busyItem { return (*h)[0] }

// --- Main solution ---

func busiestServers(k int, arrival []int, load []int) []int {
	counts := make([]int, k)
	available := newBIT(k)
	busy := &busyHeap{}
	heap.Init(busy)

	maxCount := 0

	for i := 0; i < len(arrival); i++ {
		t := arrival[i]

		// Free completed servers
		for busy.Len() > 0 && busy.Peek().endTime <= t {
			item := heap.Pop(busy).(*busyItem)
			available.addServer(item.index)
		}

		// Find server to assign
		serverIdx := available.ceil(i % k)
		if serverIdx == -1 {
			continue // all servers busy, drop request
		}

		available.remove(serverIdx)
		heap.Push(busy, &busyItem{endTime: t + load[i], index: serverIdx})
		counts[serverIdx]++
		if counts[serverIdx] > maxCount {
			maxCount = counts[serverIdx]
		}
	}

	result := []int{}
	for i, c := range counts {
		if c == maxCount {
			result = append(result, i)
		}
	}
	sort.Ints(result)
	return result
}

func main() {
	// Test case 1: k=3, arrival=[1,2,3,4,5], load=[5,2,3,3,3] -> [1]
	k := 3
	arrival := []int{1, 2, 3, 4, 5}
	load := []int{5, 2, 3, 3, 3}
	result := busiestServers(k, arrival, load)
	fmt.Printf("k=%d arrival=%v load=%v -> %v\n", k, arrival, load, result)

	// Test case 2: k=3, arrival=[1,2,3,4,8,9,10], load=[5,2,3,3,2,1,1] -> [1]
	arrival2 := []int{1, 2, 3, 4, 8, 9, 10}
	load2 := []int{5, 2, 3, 3, 2, 1, 1}
	result2 := busiestServers(3, arrival2, load2)
	fmt.Printf("k=3 arrival=%v load=%v -> %v\n", arrival2, load2, result2)

	// Test case 3: k=1, arrival=[1], load=[1] -> [0]
	result3 := busiestServers(1, []int{1}, []int{1})
	fmt.Printf("k=1 arrival=[1] load=[1] -> %v\n", result3)
}
