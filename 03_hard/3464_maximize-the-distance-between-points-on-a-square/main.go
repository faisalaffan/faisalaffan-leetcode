package main

// LeetCode #3464: Maximize the Distance Between Points on a Square
// https://leetcode.com/problems/maximize-the-distance-between-points-on-a-square/
// Difficulty: Hard
//
// Given the side length of a square and a set of points placed on its boundary,
// choose k points to maximize the minimum Manhattan distance between any two
// chosen points.
//
// Approach: Binary search on answer D. Check if we can select k points with
// minimum distance >= D using greedy placement on the square perimeter.

import "fmt"
	
	"sort"


func main() {
	// Example 1
	fmt.Println(maxDistance(3, [][]int{{0, 0}, {3, 0}, {3, 3}, {0, 3}}, 3))
	// Example 2
	fmt.Println(maxDistance(2, [][]int{{0, 0}, {2, 0}, {2, 2}, {0, 2}}, 2))
	// Edge: two points
	fmt.Println(maxDistance(5, [][]int{{0, 0}, {5, 0}}, 2))
	// Edge: k = 1
	fmt.Println(maxDistance(10, [][]int{{0, 0}, {5, 5}}, 1))
}

func maxDistance(side int, points [][]int, k int) int {
	n := len(points)
	// Convert points to perimeter positions (distance from (0,0) along perimeter)
	pos := make([]int, n)
	for i, p := range points {
		x, y := p[0], p[1]
		if y == 0 {
			pos[i] = x
		} else if x == side {
			pos[i] = side + y
		} else if y == side {
			pos[i] = 3*side - x
		} else {
			pos[i] = 4*side - y
		}
	}

	sort.Ints(pos)
	perimeter := 4 * side

	// Binary search on answer
	left, right := 0, perimeter
	for left < right {
		mid := (left + right + 1) / 2
		if canPlace(pos, perimeter, k, mid) {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return left
}

func canPlace(pos []int, perimeter int, k int, minDist int) bool {
	n := len(pos)
	// Try each starting point
	for start := 0; start < n; start++ {
		count := 1
		last := pos[start]
		// Greedily place points
		for i := 1; i < k; i++ {
			// Find next point at distance >= minDist
			found := false
			for j := 0; j < n; j++ {
				idx := (start + j) % n
				dist := pos[idx] - last
				if dist < 0 {
					dist += perimeter
				}
				if dist >= minDist {
					last = pos[idx]
					count++
					found = true
					break
				}
			}
			if !found {
				break
			}
		}
		if count >= k {
			// Verify the wrap-around distance
			wrapDist := pos[start] - last
			if wrapDist < 0 {
				wrapDist += perimeter
			}
			if wrapDist >= minDist || k <= 1 {
				return true
			}
			if wrapDist >= minDist {
				return true
			}
		}
	}
	return false
}
