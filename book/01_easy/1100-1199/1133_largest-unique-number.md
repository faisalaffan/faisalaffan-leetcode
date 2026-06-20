# 1133 — Largest Unique Number

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func largestUniqueNumber(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1133: Largest Unique Number
// https://leetcode.com/problems/largest-unique-number/
// Difficulty: Easy [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(largestUniqueNumber([]int{5, 7, 3, 9, 4, 9, 8, 3, 1})) // 8
	fmt.Println(largestUniqueNumber([]int{9, 9, 8, 8}))                // -1
}

// LeetCode submission: largestUniqueNumber
func largestUniqueNumber(nums []int) int {
  // HashMap: O(1) lookup
	count := make(map[int]int)
	for _, v := range nums {
		count[v]++
	}
	ans := -1
	for k, v := range count {
		if v == 1 && k > ans {
			ans = k
		}
	}
	return ans
}
```
