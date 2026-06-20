# 0353 — Design Snake Game

## Deskripsi

**Soal:** [0353. Design Snake Game](https://leetcode.com/problems/design-snake-game/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** LIS (Longest Increasing Subsequence)

**Fungsi Solusi:** `func Constructor(width int, height int, food [][]int) SnakeGame`

## Solusi Go

```go
package main

// LeetCode #353: Design Snake Game
// https://leetcode.com/problems/design-snake-game/
// Difficulty: Medium [Paid]
// Time O(1) per move | Space O(n)

import "fmt"

type SnakeGame struct {
	width, height int
	food          [][]int
	foodIdx       int
	snake         [][2]int // head is last element
	body          map[[2]int]struct{}
	score         int
}

func Constructor(width int, height int, food [][]int) SnakeGame {
	head := [2]int{0, 0}
	return SnakeGame{
		width: width, height: height,
		food:    food,
		foodIdx: 0,
		snake:   [][2]int{head},
		body:    map[[2]int]struct{}{head: {}},
		score:   0,
	}
}

func (sg *SnakeGame) Move(direction string) int {
	head := sg.snake[len(sg.snake)-1]
	var next [2]int
	switch direction {
	case "U":
		next = [2]int{head[0] - 1, head[1]}
	case "D":
		next = [2]int{head[0] + 1, head[1]}
	case "L":
		next = [2]int{head[0], head[1] - 1}
	case "R":
		next = [2]int{head[0], head[1] + 1}
	}

	// Check boundaries
	if next[0] < 0 || next[0] >= sg.height || next[1] < 0 || next[1] >= sg.width {
		return -1
	}

	// Check if eating food
	eat := sg.foodIdx < len(sg.food) && next[0] == sg.food[sg.foodIdx][0] && next[1] == sg.food[sg.foodIdx][1]

	if eat {
		sg.score++
		sg.foodIdx++
	} else {
		// Remove tail
		tail := sg.snake[0]
		delete(sg.body, tail)
		sg.snake = sg.snake[1:]
	}

	// Check collision with self
	if _, hit := sg.body[next]; hit {
		return -1
	}

	sg.snake = append(sg.snake, next)
	sg.body[next] = struct{}{}
	return sg.score
}

func main() {
	// Test case 1
	sg := Constructor(3, 2, [][]int{{1, 2}, {0, 1}})
	fmt.Println("Move R:", sg.Move("R")) // at (0,1)
	fmt.Println("Move D:", sg.Move("D")) // at (1,1)
	fmt.Println("Move R:", sg.Move("R")) // at (1,2), eat food → score 1
	fmt.Println("Move U:", sg.Move("U")) // at (0,2)
	fmt.Println("Move L:", sg.Move("L")) // at (0,1) - hits body
	// Expected: 0, 0, 1, 1, -1

	fmt.Println()

	// Test case 2: Hit wall
	sg2 := Constructor(1, 1, [][]int{})
	fmt.Println("Move R:", sg2.Move("R"))
	// Expected: -1
}
```
