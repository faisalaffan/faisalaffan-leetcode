# 0682 — Baseball Game

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func calPoints(operations []string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Stack

**Waktu:** O(n). Space: O(n).  |  **Ruang:** O(n).

> 🎓 **Fresh Grad Tips:** Kuasai **Stack** — sering muncul di interview!

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
  // Alokasi slice
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
