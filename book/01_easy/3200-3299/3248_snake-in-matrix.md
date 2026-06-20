# 3248 — Snake In Matrix

## Deskripsi

**Soal:** [3248. Snake In Matrix](https://leetcode.com/problems/snake-in-matrix/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(m). Space: O(1).  
**Kompleksitas Ruang:** O(1).

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3248: Snake in Matrix
// https://leetcode.com/problems/snake-in-matrix/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(SnakeInMatrix(3, []string{"RIGHT", "DOWN"}))
	fmt.Println(SnakeInMatrix(2, []string{"DOWN", "RIGHT", "UP"}))
}

// SnakeInMatrix returns the final position of the snake in an n x n matrix after following commands.
// Time: O(m). Space: O(1).
func SnakeInMatrix(n int, commands []string) int {
	r, c := 0, 0
	for _, cmd := range commands {
		switch cmd {
		case "UP":
			r--
		case "DOWN":
			r++
		case "LEFT":
			c--
		case "RIGHT":
			c++
		}
	}
	return r*n + c
}
```
