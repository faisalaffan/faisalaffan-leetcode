# 2006 — Count Number Of Pairs With Absolute Difference K

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func CountNumberOfPairsWithAbsoluteDifferenceK(nums []int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n), Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2006: Count Number of Pairs With Absolute Difference K
// https://leetcode.com/problems/count-number-of-pairs-with-absolute-difference-k/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CountNumberOfPairsWithAbsoluteDifferenceK([]int{1, 2, 2, 1}, 1))   // 4
	fmt.Println(CountNumberOfPairsWithAbsoluteDifferenceK([]int{1, 3}, 3))          // 0
	fmt.Println(CountNumberOfPairsWithAbsoluteDifferenceK([]int{3, 2, 1, 5, 4}, 2)) // 3
}

// Time: O(n), Space: O(n)
func CountNumberOfPairsWithAbsoluteDifferenceK(nums []int, k int) int {
  // HashMap: O(1) lookup
	freq := make(map[int]int)
	count := 0
	for _, v := range nums {
		count += freq[v-k] + freq[v+k]
		freq[v]++
	}
	return count
}
```
