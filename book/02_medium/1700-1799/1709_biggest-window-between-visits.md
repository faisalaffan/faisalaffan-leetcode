# 1709 — Biggest Window Between Visits

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan aturan permainan dan harus menentukan siapa yang menang atau berapa skor maksimal. Tugasmu adalah menganalisis permainan dan membuat keputusan optimal di setiap langkah.

Soal game theory menguji kemampuanmu berpikir beberapa langkah ke depan (minimax). Seringkali diselesaikan dengan DP (Dynamic Programming) untuk menyimpan hasil subproblem.

**Konsep kunci:** minimax, optimal play, game state, DP memoization, win/lose positions.

**Fungsi yang perlu kamu implementasikan:**
```go
func biggestWindow(visits []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Sliding Window

**Kompleksitas Waktu:** O(n log n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Sliding Window** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1709: Biggest Window Between Visits
// https://leetcode.com/problems/biggest-window-between-visits/
// Difficulty: Medium [Paid] (SQL)
// This is a SQL problem. In Go we simulate the logic.
// Time: O(n log n), Space: O(n)

import (
	"fmt"
	"sort"
)

func biggestWindow(visits []int) int {
	if len(visits) == 0 {
		return 0
	}
  // Urutkan secara ascending — O(n log n)
	sort.Ints(visits)
	maxGap := 0
	for i := 1; i < len(visits); i++ {
		gap := visits[i] - visits[i-1]
		if gap > maxGap {
			maxGap = gap
		}
	}
	return maxGap
}

func main() {
	fmt.Println(biggestWindow([]int{1, 3, 7, 10})) // Expected: 4 (between 3 and 7)
	fmt.Println(biggestWindow([]int{1, 2, 3, 4}))  // Expected: 1
	fmt.Println(biggestWindow([]int{5}))            // Expected: 0
}
```
