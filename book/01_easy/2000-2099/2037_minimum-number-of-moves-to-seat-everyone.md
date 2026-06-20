# 2037 — Minimum Number Of Moves To Seat Everyone

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func MinimumNumberOfMovesToSeatEveryone(seats []int, students []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n), Space: O(1)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2037: Minimum Number of Moves to Seat Everyone
// https://leetcode.com/problems/minimum-number-of-moves-to-seat-everyone/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MinimumNumberOfMovesToSeatEveryone([]int{3, 1, 5}, []int{2, 7, 4}))   // 4
	fmt.Println(MinimumNumberOfMovesToSeatEveryone([]int{4, 1, 5, 9}, []int{1, 3, 2, 6})) // 7
}

// Time: O(n log n), Space: O(1)
func MinimumNumberOfMovesToSeatEveryone(seats []int, students []int) int {
  // Sort O(n log n)
	sort.Ints(seats)
  // Sort O(n log n)
	sort.Ints(students)
	moves := 0
  // Linear scan O(n)
	for i := 0; i < len(seats); i++ {
		diff := seats[i] - students[i]
		if diff < 0 {
			diff = -diff
		}
		moves += diff
	}
	return moves
}
```
