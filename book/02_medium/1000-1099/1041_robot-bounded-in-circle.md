# 1041 — Robot Bounded In Circle

## Deskripsi

**Soal:** [1041. Robot Bounded In Circle](https://leetcode.com/problems/robot-bounded-in-circle/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

> **Ide Kunci:** Simulate robot movement. Robot is bounded in circle iff

## Solusi Go

```go
package main

// LeetCode #1041: Robot Bounded In Circle
// https://leetcode.com/problems/robot-bounded-in-circle/
// Difficulty: Medium
//
// Approach: Simulate robot movement. Robot is bounded in circle iff
//           final position is (0,0) OR direction != North.
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(isRobotBounded("GGLLGG")) // true
	fmt.Println(isRobotBounded("GG"))     // false
	fmt.Println(isRobotBounded("GL"))     // true
}

func isRobotBounded(instructions string) bool {
	x, y := 0, 0
	dirX, dirY := 0, 1 // facing north

	for _, c := range instructions {
		switch c {
		case 'G':
			x += dirX
			y += dirY
		case 'L':
			dirX, dirY = -dirY, dirX
		case 'R':
			dirX, dirY = dirY, -dirX
		}
	}

	return (x == 0 && y == 0) || !(dirX == 0 && dirY == 1)
}
```
