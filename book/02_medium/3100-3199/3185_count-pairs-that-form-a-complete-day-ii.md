# 3185 — Count Pairs That Form A Complete Day Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func countCompleteDayPairs(hours []int) int64`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(24)


## 💻 Solusi Go

```go
package main

// LeetCode #3185: Count Pairs That Form a Complete Day II
// https://leetcode.com/problems/count-pairs-that-form-a-complete-day-ii/
// Difficulty: Medium
// Time: O(n) | Space: O(24)

import "fmt"

func countCompleteDayPairs(hours []int) int64 {
  // Alokasi slice
	count := make([]int, 24)
	var ans int64

	for _, h := range hours {
		r := h % 24
		need := (24 - r) % 24
		ans += int64(count[need])
		count[r]++
	}
	return ans
}

func main() {
	fmt.Println(countCompleteDayPairs([]int{12, 12, 30, 24, 24})) // Expected: 2
	fmt.Println(countCompleteDayPairs([]int{72, 48, 24, 3}))       // Expected: 3
}
```
