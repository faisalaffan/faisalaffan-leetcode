# 3432 — Count Partitions With Even Sum Difference

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func CountPartitionsWithEvenSumDifference(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer

**Waktu:** O(n). Space: O(1).  |  **Ruang:** O(1).

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3432: Count Partitions with Even Sum Difference
// https://leetcode.com/problems/count-partitions-with-even-sum-difference/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CountPartitionsWithEvenSumDifference([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(CountPartitionsWithEvenSumDifference([]int{10, 10, 10, 10, 10}))
}

// CountPartitionsWithEvenSumDifference counts partitions where the difference between left and right sums is even.
// Time: O(n). Space: O(1).
func CountPartitionsWithEvenSumDifference(nums []int) int {
	totalSum := 0
	for _, v := range nums {
		totalSum += v
	}
	leftSum := 0
	count := 0
  // Linear scan O(n)
	for i := 0; i < len(nums)-1; i++ {
		leftSum += nums[i]
		rightSum := totalSum - leftSum
		if (leftSum-rightSum)%2 == 0 {
			count++
		}
	}
	return count
}
```
