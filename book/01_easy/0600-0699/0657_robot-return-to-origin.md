# 0657 — Robot Return To Origin

## Deskripsi

**Soal:** [0657. Robot Return To Origin](https://leetcode.com/problems/robot-return-to-origin/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n). Space: O(1).  
**Kompleksitas Ruang:** O(1).

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #657: Robot Return to Origin
// https://leetcode.com/problems/robot-return-to-origin/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(judgeCircle("UD"))   // true
	fmt.Println(judgeCircle("LL"))   // false
	fmt.Println(judgeCircle(""))     // true
}

// judgeCircle returns true if the robot returns to origin after executing all moves.
// Time: O(n). Space: O(1).
func judgeCircle(moves string) bool {
	x, y := 0, 0
	for _, c := range moves {
		switch c {
		case 'U':
			y++
		case 'D':
			y--
		case 'L':
			x--
		case 'R':
			x++
		}
	}
	return x == 0 && y == 0
}
```
