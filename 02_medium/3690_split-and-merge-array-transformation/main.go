package main

// LeetCode #3690: Split and Merge Array Transformation
// https://leetcode.com/problems/split-and-merge-array-transformation/
// Difficulty: Medium
// Time: O(n! * n^4) | Space: O(n! * n)

import "fmt"

func splitAndMergeArrayTransformation(nums1 []int, nums2 []int) int {
	n := len(nums1)
	target := make([]int, n)
	copy(target, nums2)

	type state struct {
		arr   []int
		steps int
	}

	queue := []state{{arr: append([]int(nil), nums1...), steps: 0}}
	visited := make(map[string]bool)

	key := func(arr []int) string {
		b := make([]byte, len(arr)*4)
		for i, v := range arr {
			b[i*4] = byte(v >> 24)
			b[i*4+1] = byte(v >> 16)
			b[i*4+2] = byte(v >> 8)
			b[i*4+3] = byte(v)
		}
		return string(b)
	}

	visited[key(nums1)] = true

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if equal(cur.arr, target) {
			return cur.steps
		}

		// Try all subarrays [L, R]
		for L := 0; L < n; L++ {
			for R := L; R < n; R++ {
				sub := make([]int, R-L+1)
				copy(sub, cur.arr[L:R+1])

				remain := make([]int, 0, n-(R-L+1))
				remain = append(remain, cur.arr[:L]...)
				remain = append(remain, cur.arr[R+1:]...)

				// Insert sub at all positions in remain
				for pos := 0; pos <= len(remain); pos++ {
					next := make([]int, 0, n)
					next = append(next, remain[:pos]...)
					next = append(next, sub...)
					next = append(next, remain[pos:]...)

					k := key(next)
					if !visited[k] {
						visited[k] = true
						queue = append(queue, state{arr: next, steps: cur.steps + 1})
					}
				}
			}
		}
	}

	return -1
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(splitAndMergeArrayTransformation([]int{3, 1, 2}, []int{1, 2, 3}))
	fmt.Println(splitAndMergeArrayTransformation([]int{1, 2, 3}, []int{1, 2, 3}))
	fmt.Println(splitAndMergeArrayTransformation([]int{2, 1}, []int{1, 2}))
}
