package main

// LeetCode #2069: Walking Robot Simulation II
// https://leetcode.com/problems/walking-robot-simulation-ii/
// Difficulty: Medium
// Time: O(1) per operation | Space: O(1)

import "fmt"

type Robot struct {
	w, h   int
	x, y   int
	dir    int
	dirs   [][2]int
	dirStr []string
	moved  bool
}

func Constructor(width int, height int) Robot {
	return Robot{
		w:      width,
		h:      height,
		x:      0,
		y:      0,
		dir:    0,
		dirs:   [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}},
		dirStr: []string{"East", "North", "West", "South"},
		moved:  false,
	}
}

func (r *Robot) Step(num int) {
	r.moved = true
	perimeter := 2 * (r.w + r.h - 2)
	if num >= perimeter {
		num %= perimeter
		if r.x == 0 && r.y == 0 {
			r.dir = 0 // Reset direction to East
		}
	}

	for i := 0; i < num; i++ {
		nx := r.x + r.dirs[r.dir][0]
		ny := r.y + r.dirs[r.dir][1]
		if nx < 0 || nx >= r.w || ny < 0 || ny >= r.h {
			r.dir = (r.dir + 1) % 4
			nx = r.x + r.dirs[r.dir][0]
			ny = r.y + r.dirs[r.dir][1]
		}
		r.x, r.y = nx, ny
	}
}

func (r *Robot) GetPos() []int {
	return []int{r.x, r.y}
}

func (r *Robot) GetDir() string {
	if !r.moved || (r.x == 0 && r.y == 0) {
		return "East"
	}
	return r.dirStr[r.dir]
}

func main() {
	robot := Constructor(6, 3)
	robot.Step(2)
	fmt.Println("Test 1 Pos:", robot.GetPos()) // [2, 0]
	fmt.Println("Test 1 Dir:", robot.GetDir()) // East
	robot.Step(2)
	fmt.Println("Test 2 Pos:", robot.GetPos()) // [4, 0]
	fmt.Println("Test 2 Dir:", robot.GetDir()) // East
	robot.Step(2)
	fmt.Println("Test 3 Pos:", robot.GetPos()) // [5, 1]
	fmt.Println("Test 3 Dir:", robot.GetDir()) // North

	robot2 := Constructor(3, 2)
	robot2.Step(2)
	robot2.Step(3) // goes around perimeter
	fmt.Println("Test 4 Pos:", robot2.GetPos())
	fmt.Println("Test 4 Dir:", robot2.GetDir())
}
