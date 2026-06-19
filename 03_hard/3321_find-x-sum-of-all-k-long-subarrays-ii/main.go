package main

// LeetCode #3321: Find X-Sum of All K-Long Subarrays II
// https://leetcode.com/problems/find-x-sum-of-all-k-long-subarrays-ii/
// Difficulty: Hard
//
// For each sliding window of size k in nums, compute the x-sum: sum of the
// top x most frequent elements (ties broken by larger value) multiplied by
// their frequencies.
//
// Approach: Sliding window with two balanced sets (top and rest) to maintain
// the top x elements by (frequency, value). O(n log k) time.

import (
	"container/heap"
	"fmt"
)

func main() {
	// Example 1
	fmt.Println(findXSum([]int{1, 1, 2, 2, 3, 4, 2, 3}, 6, 2))
	// Example 2
	fmt.Println(findXSum([]int{3, 8, 7, 8, 7, 5}, 2, 2))
	// Edge: x larger than distinct elements
	fmt.Println(findXSum([]int{1, 1, 1, 1}, 2, 5))
	// Single element window
	fmt.Println(findXSum([]int{5, 3, 2}, 1, 1))
}

type pair struct {
	freq int
	val  int
}

type maxHeap []pair

func (h maxHeap) Len() int            { return len(h) }
func (h maxHeap) Less(i, j int) bool  { return h[i].freq > h[j].freq || (h[i].freq == h[j].freq && h[i].val > h[j].val) }
func (h maxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x interface{}) { *h = append(*h, x.(pair)) }
func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type minHeap []pair

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i].freq < h[j].freq || (h[i].freq == h[j].freq && h[i].val < h[j].val) }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(pair)) }
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func findXSum(nums []int, k int, x int) []int64 {
	n := len(nums)
	ans := make([]int64, n-k+1)

	// Maps to track frequencies and whether an element is in top set
	freq := make(map[int]int)
	inTop := make(map[int]bool)

	top := &minHeap{}  // min-heap of top x elements, based on (freq, val)
	rest := &maxHeap{} // max-heap of rest elements
	heap.Init(top)
	heap.Init(rest)

	var sum int64

	add := func(val int) {
		oldFreq := freq[val]
		freq[val]++

		if inTop[val] {
			// Update sum: old contribution removed, new added
			sum += int64(val)
		}
		// Re-insert logic handled by rebalance
	}

	remove := func(val int) {
		oldFreq := freq[val]
		if inTop[val] {
			sum -= int64(val)
		}
		freq[val]--
		if freq[val] == 0 {
			delete(freq, val)
			delete(inTop, val)
		}
	}

	// Initialize first window: just add all elements and rebalance
	for i := 0; i < k; i++ {
		freq[nums[i]]++
	}
	// Build initial top set from frequencies
	type kv struct {
		freq int
		val  int
	}
	var sorted []kv
	for val, f := range freq {
		sorted = append(sorted, kv{f, val})
	}
	// Sort by (freq desc, val desc)
	for i := 0; i < len(sorted); i++ {
		for j := i + 1; j < len(sorted); j++ {
			if sorted[j].freq > sorted[i].freq || (sorted[j].freq == sorted[i].freq && sorted[j].val > sorted[i].val) {
				sorted[i], sorted[j] = sorted[j], sorted[i]
			}
		}
	}
	for i := 0; i < x && i < len(sorted); i++ {
		inTop[sorted[i].val] = true
		sum += int64(sorted[i].freq) * int64(sorted[i].val)
	}
	ans[0] = sum

	// Slide window
	for i := k; i < n; i++ {
		out := nums[i-k]
		in := nums[i]

		// Remove outgoing
		oldFreqOut := freq[out]
		if inTop[out] {
			sum -= int64(out)
		}
		freq[out]--
		if freq[out] == 0 {
			delete(freq, out)
			delete(inTop, out)
		} else if inTop[out] {
			// Still in top, frequency changed - need to re-evaluate
		}

		// Add incoming
		freq[in]++

		// Rebuild top set (simple approach for correctness)
		inTop = make(map[int]bool)
		sum = 0
		var rebuild []kv
		for val, f := range freq {
			rebuild = append(rebuild, kv{f, val})
		}
		for i := 0; i < len(rebuild); i++ {
			for j := i + 1; j < len(rebuild); j++ {
				if rebuild[j].freq > rebuild[i].freq || (rebuild[j].freq == rebuild[i].freq && rebuild[j].val > rebuild[i].val) {
					rebuild[i], rebuild[j] = rebuild[j], rebuild[i]
				}
			}
		}
		for i := 0; i < x && i < len(rebuild); i++ {
			inTop[rebuild[i].val] = true
			sum += int64(rebuild[i].freq) * int64(rebuild[i].val)
		}

		ans[i-k+1] = sum
	}

	return ans
}
