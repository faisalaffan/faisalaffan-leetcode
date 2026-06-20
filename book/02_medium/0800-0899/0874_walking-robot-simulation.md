# 0874 — Walking Robot Simulation

## Deskripsi

**Soal:** [0874. Walking Robot Simulation](https://leetcode.com/problems/walking-robot-simulation/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n + m) where n = len(commands), m = len(obstacles)  
**Kompleksitas Ruang:** O(m)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #874: Walking Robot Simulation
// https://leetcode.com/problems/walking-robot-simulation/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(WalkingRobotSimulation([]int{4, -1, 3}, [][]int{}))
	fmt.Println(WalkingRobotSimulation([]int{4, -1, 4, -2, 4}, [][]int{{2, 4}}))
	fmt.Println(WalkingRobotSimulation([]int{6, -1, -1, 6}, [][]int{}))
}

// Time: O(n + m) where n = len(commands), m = len(obstacles) | Space: O(m)
func WalkingRobotSimulation(commands []int, obstacles [][]int) int {
  // Membuat map untuk pencarian O(1): key → value
	obsSet := make(map[[2]int]bool)
	for _, o := range obstacles {
		obsSet[[2]int{o[0], o[1]}] = true
	}

	dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	x, y, dir, maxDist := 0, 0, 0, 0

	for _, cmd := range commands {
		if cmd == -1 {
			dir = (dir + 1) % 4
		} else if cmd == -2 {
			dir = (dir + 3) % 4
		} else {
			for step := 0; step < cmd; step++ {
				nx, ny := x+dirs[dir][0], y+dirs[dir][1]
				if obsSet[[2]int{nx, ny}] {
					break
				}
				x, y = nx, ny
				dist := x*x + y*y
				if dist > maxDist {
					maxDist = dist
				}
			}
		}
	}

	return maxDist
}
```
