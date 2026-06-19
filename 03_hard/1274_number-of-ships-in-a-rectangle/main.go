package main

// LeetCode #1274: Number of Ships in a Rectangle
// https://leetcode.com/problems/number-of-ships-in-a-rectangle/
// Difficulty: Hard [Paid]
//
// Approach: Divide-and-conquer (quadtree).
// The Sea API tells whether a rectangle contains at least one ship.
// Recursively split the rectangle into 4 quadrants until reaching
// single points, counting each ship.

import "fmt"

// Sea is a mock of the LeetCode API.
type Sea struct {
	hasShipsFunc func(topRight, bottomLeft []int) bool
}

func (s *Sea) hasShips(topRight, bottomLeft []int) bool {
	return s.hasShipsFunc(topRight, bottomLeft)
}

// countShips returns the number of ships in the given rectangle.
func countShips(sea *Sea, topRight, bottomLeft []int) int {
	x1, y1 := bottomLeft[0], bottomLeft[1]
	x2, y2 := topRight[0], topRight[1]
	if x1 > x2 || y1 > y2 {
		return 0
	}
	if !sea.hasShips(topRight, bottomLeft) {
		return 0
	}
	if x1 == x2 && y1 == y2 {
		return 1
	}
	midX := (x1 + x2) / 2
	midY := (y1 + y2) / 2
	cnt := 0
	// bottom-left
	cnt += countShips(sea, []int{midX, midY}, []int{x1, y1})
	// bottom-right
	cnt += countShips(sea, []int{x2, midY}, []int{midX + 1, y1})
	// top-left
	cnt += countShips(sea, []int{midX, y2}, []int{x1, midY + 1})
	// top-right
	cnt += countShips(sea, []int{x2, y2}, []int{midX + 1, midY + 1})
	return cnt
}

func main() {
	// Test: 5 ships at known positions
	ships := [][2]int{{1, 1}, {2, 2}, {3, 3}, {5, 5}, {4, 5}}
	sea := &Sea{
		hasShipsFunc: func(topRight, bottomLeft []int) bool {
			for _, s := range ships {
				if s[0] >= bottomLeft[0] && s[0] <= topRight[0] &&
					s[1] >= bottomLeft[1] && s[1] <= topRight[1] {
					return true
				}
			}
			return false
		},
	}
	fmt.Println(countShips(sea, []int{4, 4}, []int{1, 1})) // 3
	fmt.Println(countShips(sea, []int{6, 6}, []int{1, 1})) // 5
	fmt.Println(countShips(sea, []int{0, 0}, []int{0, 0})) // 0
	fmt.Println(countShips(sea, []int{5, 5}, []int{4, 4})) // 2
}
