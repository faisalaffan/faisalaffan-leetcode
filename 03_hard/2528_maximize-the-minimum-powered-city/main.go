package main

// LeetCode #2528: Maximize the Minimum Powered City
// https://leetcode.com/problems/maximize-the-minimum-powered-city/
// Difficulty: Hard
//
// Given an array stations where stations[i] is the power of the station
// at city i, a radius r, and k additional stations that can be placed,
// maximize the minimum total power across all cities.
// Each station powers cities within distance r.
//
// Approach: Binary search on the minimum power. Use difference array
// to check if we can achieve at least x power at every city.

import "fmt"

func main() {
	// Example 1
	fmt.Println(maxPower([]int{1, 2, 4, 5, 0}, 1, 2))
	// Example 2
	fmt.Println(maxPower([]int{4, 4, 4, 4}, 0, 3))
	// Example 3
	fmt.Println(maxPower([]int{5, 10, 5}, 1, 2))
	// Edge: single city
	fmt.Println(maxPower([]int{10}, 0, 5))
}

func maxPower(stations []int, r int, k int) int64 {
	n := len(stations)
	d := make([]int, n+1)
	s := make([]int, n+1)
	for i, v := range stations {
		left, right := max(0, i-r), min(i+r, n-1)
		d[left] += v
		d[right+1] -= v
	}
	s[0] = d[0]
	for i := 1; i < n+1; i++ {
		s[i] = s[i-1] + d[i]
	}
	check := func(x, k int) bool {
		d := make([]int, n+1)
		t := 0
		for i := range stations {
			t += d[i]
			dist := x - (s[i] + t)
			if dist > 0 {
				if k < dist {
					return false
				}
				k -= dist
				j := min(i+r, n-1)
				left, right := max(0, j-r), min(j+r, n-1)
				d[left] += dist
				d[right+1] -= dist
				t += dist
			}
		}
		return true
	}
	left, right := 0, 1<<40
	for left < right {
		mid := (left + right + 1) >> 1
		if check(mid, k) {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return int64(left)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
