package main

// LeetCode #3009: Maximum Number of Intersections on the Chart
// https://leetcode.com/problems/maximum-number-of-intersections-on-the-chart/
// Difficulty: Hard [Paid]

import (
	"fmt"
	"sort"
)

func maxIntersectionCount(y []int) int {
	n := len(y)
	type event struct {
		x     int
		delta int
	}
	events := make([]event, 0, 2*n)
	for i := 1; i < n; i++ {
		s := 2 * y[i-1]
		e := 2 * y[i]
		if i != n-1 {
			if y[i-1] < y[i] {
				e--
			} else {
				e++
			}
		}
		if s > e {
			s, e = e, s
		}
		events = append(events, event{s, 1}, event{e + 1, -1})
	}
	sort.Slice(events, func(i, j int) bool {
		if events[i].x != events[j].x {
			return events[i].x < events[j].x
		}
		return events[i].delta < events[j].delta
	})
	ans := 0
	cur := 0
	for _, ev := range events {
		cur += ev.delta
		if cur > ans {
			ans = cur
		}
	}
	return ans
}

func main() {
	fmt.Println(maxIntersectionCount([]int{1, 2, 1, 2, 1, 3, 2}))
	fmt.Println(maxIntersectionCount([]int{2, 1, 3, 4, 5}))
}
