# 2240 — Number Of Ways To Buy Pens And Pencils

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func waysToBuyPensPencils(total int, cost1 int, cost2 int) int64`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(total / cost1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2240: Number of Ways to Buy Pens and Pencils
// https://leetcode.com/problems/number-of-ways-to-buy-pens-and-pencils/
// Difficulty: Medium
// Time: O(total / cost1) | Space: O(1)

import "fmt"

func waysToBuyPensPencils(total int, cost1 int, cost2 int) int64 {
	var ways int64 = 0
	for pens := 0; pens*cost1 <= total; pens++ {
		remaining := total - pens*cost1
		ways += int64(remaining/cost2) + 1
		if cost2 == 0 {
			break
		}
	}
	return ways
}

func main() {
	// Test case 1
	fmt.Println(waysToBuyPensPencils(20, 10, 5))
	// Expected: 9

	// Test case 2
	fmt.Println(waysToBuyPensPencils(5, 10, 10))
	// Expected: 1

	// Test case 3
	fmt.Println(waysToBuyPensPencils(100, 1, 1))
	// Expected: 5151
}
```
