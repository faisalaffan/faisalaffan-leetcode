# 0682 — Baseball Game

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan aturan permainan dan harus menentukan siapa yang menang atau berapa skor maksimal. Tugasmu adalah menganalisis permainan dan membuat keputusan optimal di setiap langkah.

Soal game theory menguji kemampuanmu berpikir beberapa langkah ke depan (minimax). Seringkali diselesaikan dengan DP (Dynamic Programming) untuk menyimpan hasil subproblem.

**Konsep kunci:** minimax, optimal play, game state, DP memoization, win/lose positions.

**Fungsi yang perlu kamu implementasikan:**
```go
func calPoints(operations []string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Stack

**Kompleksitas Waktu:** O(n). Space: O(n).  
**Kompleksitas Ruang:** O(n).

> **Untuk fresh graduate:** Kuasai dulu teknik **Stack** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

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
  // Alokasi slice integer
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
