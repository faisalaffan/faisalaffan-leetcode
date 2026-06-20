# 2367 — Number Of Arithmetic Triplets

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func NumberOfArithmeticTriplets(nums []int, diff int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2367: Number of Arithmetic Triplets
// https://leetcode.com/problems/number-of-arithmetic-triplets/
// Difficulty: Easy
// Time O(n) | Space O(n)

import "fmt"

func main() {
	fmt.Println(NumberOfArithmeticTriplets([]int{0, 1, 4, 6, 7, 10}, 3)) // 2
	fmt.Println(NumberOfArithmeticTriplets([]int{4, 5, 6, 7, 8, 9}, 2))  // 2
}

func NumberOfArithmeticTriplets(nums []int, diff int) int {
  // HashMap: O(1) lookup
	seen := make(map[int]bool, len(nums))
	for _, n := range nums {
		seen[n] = true
	}
	count := 0
	for _, n := range nums {
		if seen[n+diff] && seen[n+2*diff] {
			count++
		}
	}
	return count
}
```
