# 2824 — Count Pairs Whose Sum Is Less Than Target

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func CountPairsWhoseSumIsLessThanTarget(nums []int, target int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n^2)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2824: Count Pairs Whose Sum is Less than Target
// https://leetcode.com/problems/count-pairs-whose-sum-is-less-than-target/
// Difficulty: Easy
// Time: O(n^2) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(CountPairsWhoseSumIsLessThanTarget([]int{-1, 1, 2, 3, 1}, 2))
	fmt.Println(CountPairsWhoseSumIsLessThanTarget([]int{-6, 2, 5, -2, -7, -1, 3}, -2))
}

func CountPairsWhoseSumIsLessThanTarget(nums []int, target int) int {
	count := 0
  // Linear scan O(n)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] < target {
				count++
			}
		}
	}
	return count
}
```
