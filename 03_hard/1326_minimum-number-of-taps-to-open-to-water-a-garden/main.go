package main

// LeetCode #1326: Minimum Number of Taps to Open to Water a Garden
// https://leetcode.com/problems/minimum-number-of-taps-to-open-to-water-a-garden/
// Difficulty: Hard

import "fmt"

func main() {
	fmt.Println("1326. Minimum Number of Taps to Open to Water a Garden")
	fmt.Println("n=5, ranges=[3,4,1,1,0,0]:", minTaps(5, []int{3, 4, 1, 1, 0, 0}), "(expected 1)")
	fmt.Println("n=5, ranges=[3,4,1,1,2,0]:", minTaps(5, []int{3, 4, 1, 1, 2, 0}), "(expected 1)")
	fmt.Println("n=3, ranges=[0,0,0,0]:", minTaps(3, []int{0, 0, 0, 0}), "(expected -1)")
}

func minTaps(n int, ranges []int) int {
	// farthest[p] = rightmost reachable point starting from position p.
	farthest := make([]int, n+1)
	for i, r := range ranges {
		left := i - r
		if left < 0 {
			left = 0
		}
		right := i + r
		if right > n {
			right = n
		}
		if right > farthest[left] {
			farthest[left] = right
		}
	}

	// Greedy jump-game style traversal.
	taps := 0
	curEnd := 0
	nextEnd := 0

	for i := 0; i <= n; i++ {
		if farthest[i] > nextEnd {
			nextEnd = farthest[i]
		}
		if i == curEnd {
			if curEnd == nextEnd {
				// We haven't made progress — unreachable.
				break
			}
			taps++
			curEnd = nextEnd
			if curEnd >= n {
				return taps
			}
		}
	}

	return -1
}
