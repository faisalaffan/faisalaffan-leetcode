package main

// LeetCode #702: Search in a Sorted Array of Unknown Size
// https://leetcode.com/problems/search-in-a-sorted-array-of-unknown-size/
// Difficulty: Medium [Paid]
// Time: O(log n)
// Space: O(1)

import "fmt"

func main() {
	reader := &ArrayReader{arr: []int{-1, 0, 3, 5, 9, 12}}
	fmt.Println(search(reader, 9))
	fmt.Println(search(reader, 2))
}

type ArrayReader struct {
	arr []int
}

func (r *ArrayReader) get(index int) int {
	if index >= len(r.arr) {
		return 1 << 31 - 1
	}
	return r.arr[index]
}

func search(reader *ArrayReader, target int) int {
	// Find upper bound
	left, right := 0, 1
	for reader.get(right) < target {
		left = right
		right <<= 1
	}

	// Binary search
	for left <= right {
		mid := left + (right-left)/2
		val := reader.get(mid)
		if val == target {
			return mid
		} else if val < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
