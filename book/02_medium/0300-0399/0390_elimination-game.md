# 0390 — Elimination Game

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func lastRemaining(n int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(log n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #390: Elimination Game
// https://leetcode.com/problems/elimination-game/
// Difficulty: Medium
// Time: O(log n) | Space: O(1)

import "fmt"

func lastRemaining(n int) int {
	head := 1
	remaining := n
	step := 1
	leftToRight := true

	for remaining > 1 {
		if leftToRight || remaining%2 == 1 {
			head += step
		}
		remaining /= 2
		step *= 2
		leftToRight = !leftToRight
	}
	return head
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", lastRemaining(9))
	// Expected: 6

	// Test case 2
	fmt.Println("Test 2:", lastRemaining(1))
	// Expected: 1

	// Test case 3
	fmt.Println("Test 3:", lastRemaining(100))
	// Expected: 54
}
```
