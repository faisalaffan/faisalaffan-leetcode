# 2751 — Robot Collisions

## Deskripsi

**Soal:** [2751. Robot Collisions](https://leetcode.com/problems/robot-collisions/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Stack (tumpukan LIFO), LIS (Longest Increasing Subsequence)

> **Ide Kunci:** Stack simulation.

## Solusi Go

```go
package main

// LeetCode #2751: Robot Collisions
// https://leetcode.com/problems/robot-collisions/
// Difficulty: Hard
//
// Approach: Stack simulation.
// Sort robots by position. Use a stack to track surviving robots.
// When a left-moving robot encounters a right-moving robot, they collide.
// The robot with lower health is destroyed; the survivor's health decreases by 1.
// If healths are equal, both are destroyed.

import (
	"fmt"
	"sort"
)

func main() {
	// Example 1: positions=[5,4,3,2,1], healths=[2,17,9,15,10], directions="RRRRR" -> [2,17,9,15,10]
	fmt.Println(survivedRobotsHealths([]int{5, 4, 3, 2, 1}, []int{2, 17, 9, 15, 10}, "RRRRR"))
	// Example 2: positions=[3,5,2,6], healths=[10,10,15,12], directions="RLRL" -> [14]
	fmt.Println(survivedRobotsHealths([]int{3, 5, 2, 6}, []int{10, 10, 15, 12}, "RLRL"))
	// Example 3: positions=[1,2,5,6], healths=[10,10,11,11], directions="RLRL" -> []
	fmt.Println(survivedRobotsHealths([]int{1, 2, 5, 6}, []int{10, 10, 11, 11}, "RLRL"))
}

func survivedRobotsHealths(positions []int, healths []int, direction string) []int {
	n := len(positions)

	type robot struct {
		pos, health, idx int
		dir              byte
	}

  // Membuat slice untuk menyimpan hasil
	robots := make([]robot, n)
	for i := 0; i < n; i++ {
		robots[i] = robot{positions[i], healths[i], i, direction[i]}
	}

	sort.Slice(robots, func(i, j int) bool {
		return robots[i].pos < robots[j].pos
	})

  // Membuat slice untuk menyimpan hasil
	stack := make([]int, 0, n) // indices into robots (surviving, sorted by pos)

	for i := 0; i < n; i++ {
		if robots[i].dir == 'R' {
			stack = append(stack, i)
			continue
		}

		// Left-moving robot: fight right-moving robots on its left
		for len(stack) > 0 && robots[stack[len(stack)-1]].dir == 'R' {
			top := &robots[stack[len(stack)-1]]
			cur := &robots[i]

			if top.health > cur.health {
				top.health--
				cur.health = 0
				break
			} else if top.health < cur.health {
				cur.health--
				top.health = 0
				stack = stack[:len(stack)-1]
			} else { // equal health
				top.health = 0
				cur.health = 0
				stack = stack[:len(stack)-1]
				break
			}
		}

		// If current robot survived, push to stack
		if robots[i].health > 0 {
			stack = append(stack, i)
		}
	}

	// Collect survivors by original index
  // Membuat map untuk pencarian O(1): key → value
	survivorByIndex := make(map[int]int)
	for _, idx := range stack {
		survivorByIndex[robots[idx].idx] = robots[idx].health
	}

  // Membuat slice untuk menyimpan hasil
	result := make([]int, 0, len(stack))
	for i := 0; i < n; i++ {
		if h, ok := survivorByIndex[i]; ok {
			result = append(result, h)
		}
	}

	return result
}
```
