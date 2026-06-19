package main

// LeetCode #2013: Detect Squares
// https://leetcode.com/problems/detect-squares/
// Difficulty: Medium

import "fmt"

func main() {
	obj := Constructor2013()
	obj.Add([]int{3, 10})
	obj.Add([]int{11, 2})
	obj.Add([]int{3, 2})
	fmt.Println(obj.Count([]int{11, 10}))
	fmt.Println(obj.Count([]int{14, 8}))
	obj.Add([]int{11, 2})
	fmt.Println(obj.Count([]int{11, 10}))
}

// DetectSquares struct
type DetectSquares struct {
	points map[[2]int]int
	xMap   map[int]map[int]int
}

// Constructor2013 initializes DetectSquares
func Constructor2013() DetectSquares {
	return DetectSquares{
		points: make(map[[2]int]int),
		xMap:   make(map[int]map[int]int),
	}
}

// Add adds a point
func (this *DetectSquares) Add(point []int) {
	key := [2]int{point[0], point[1]}
	this.points[key]++
	if this.xMap[point[0]] == nil {
		this.xMap[point[0]] = make(map[int]int)
	}
	this.xMap[point[0]][point[1]]++
}

// Count counts number of ways to form a square
func (this *DetectSquares) Count(point []int) int {
	x1, y1 := point[0], point[1]
	ans := 0

	for y2 := range this.xMap[x1] {
		if y2 == y1 {
			continue
		}
		side := y2 - y1
		if side < 0 {
			side = -side
		}

		// Two possible squares: x1 +/- side
		for _, x2 := range []int{x1 - side, x1 + side} {
			if this.xMap[x2] != nil {
				ans += this.xMap[x1][y2] * this.xMap[x2][y1] * this.xMap[x2][y2]
			}
		}
	}

	return ans
}
