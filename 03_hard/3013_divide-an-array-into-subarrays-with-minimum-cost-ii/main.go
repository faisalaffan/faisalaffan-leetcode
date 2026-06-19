package main

// LeetCode #3013: Divide an Array Into Subarrays With Minimum Cost II
// https://leetcode.com/problems/divide-an-array-into-subarrays-with-minimum-cost-ii/
// Difficulty: Hard
//
// Approach: Sliding window median (min-heap + max-heap)
// We need the first subarray to start at index 0. The remaining k-1 subarrays
// each contribute the smallest element from their respective window.
// We slide a window of size 'dist' and maintain k-1 smallest elements using
// two heaps (a max-heap for the "small" group and a min-heap for the "large"
// group), tracking the sum of the small group.

import (
	"container/heap"
	"fmt"
)

type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type maxHeap []int

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func minimumCost(nums []int, k int, dist int) int64 {
	n := len(nums)
	m := k - 1 // number of elements to select from each window
	small := &maxHeap{}
	large := &minHeap{}
	heap.Init(small)
	heap.Init(large)
	sum := int64(0)
	balance := make(map[int]int)

	add := func(val int) {
		if small.Len() > 0 && val <= (*small)[0] {
			heap.Push(small, val)
			sum += int64(val)
		} else {
			heap.Push(large, val)
		}
		if small.Len() > m {
			popped := heap.Pop(small).(int)
			sum -= int64(popped)
			heap.Push(large, popped)
		} else if small.Len() < m && large.Len() > 0 {
			popped := heap.Pop(large).(int)
			sum += int64(popped)
			heap.Push(small, popped)
		}
	}

	remove := func(val int) {
		if small.Len() > 0 && val <= (*small)[0] {
			balance[val]--
			sum -= int64(val)
			for small.Len() > 0 && balance[(*small)[0]] < 0 {
				popped := heap.Pop(small).(int)
				balance[popped]++
				if balance[popped] == 0 {
					delete(balance, popped)
				}
			}
			for small.Len() < m && large.Len() > 0 {
				popped := heap.Pop(large).(int)
				sum += int64(popped)
				heap.Push(small, popped)
			}
		} else {
			balance[val]--
		}
	}

	right := dist + 1
	if right > n-1 {
		right = n - 1
	}
	for i := 1; i <= right; i++ {
		add(nums[i])
	}
	ans := int64(nums[0]) + sum

	for left := 1; left+m-1 < n; left++ {
		right = left + dist
		if right >= n {
			right = n - 1
		}
		if right < left {
			break
		}
		remove(nums[left])
		if right+1 < n {
			add(nums[right+1])
		}
		if small.Len() == m {
			candidate := int64(nums[0]) + sum
			if candidate < ans {
				ans = candidate
			}
		}
	}
	return ans
}

func main() {
	// Example: nums=[1,3,2,6,4,2], k=3, cost=2 -> 6
	fmt.Println(minimumCost([]int{1, 3, 2, 6, 4, 2}, 3, 2))
	// Another test
	fmt.Println(minimumCost([]int{1, 3, 2, 6, 4, 5}, 3, 2))
}
