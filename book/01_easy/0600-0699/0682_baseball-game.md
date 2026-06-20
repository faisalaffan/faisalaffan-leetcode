# 0682 — Baseball Game

## Deskripsi

**Soal:** [0682. Baseball Game](https://leetcode.com/problems/baseball-game/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n). Space: O(n).  
**Kompleksitas Ruang:** O(n).

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #682: Baseball Game
// https://leetcode.com/problems/baseball-game/
// Difficulty: Easy

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(calPoints([]string{"5", "2", "C", "D", "+"})) // 30
	fmt.Println(calPoints([]string{"5", "-2", "4", "C", "D", "9", "+", "+"})) // 27
	fmt.Println(calPoints([]string{"1", "C"})) // 0
}

// calPoints calculates the total score for a baseball game based on operations.
// Time: O(n). Space: O(n).
func calPoints(operations []string) int {
  // Membuat slice untuk menyimpan hasil
	stack := make([]int, 0, len(operations))
	for _, op := range operations {
		switch op {
		case "C":
			stack = stack[:len(stack)-1]
		case "D":
			stack = append(stack, 2*stack[len(stack)-1])
		case "+":
			stack = append(stack, stack[len(stack)-1]+stack[len(stack)-2])
		default:
			n, _ := strconv.Atoi(op)
			stack = append(stack, n)
		}
	}
	sum := 0
	for _, v := range stack {
		sum += v
	}
	return sum
}
```
