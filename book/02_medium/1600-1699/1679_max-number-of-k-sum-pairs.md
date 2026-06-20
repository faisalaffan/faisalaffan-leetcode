# 1679 — Max Number Of K Sum Pairs

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func maxOperations(nums []int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n), Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1679: Max Number of K-Sum Pairs
// https://leetcode.com/problems/max-number-of-k-sum-pairs/
// Difficulty: Medium
// Time: O(n), Space: O(n)

import "fmt"

func maxOperations(nums []int, k int) int {
  // HashMap: O(1) lookup
	counts := make(map[int]int)
	ops := 0

	for _, num := range nums {
		complement := k - num
		if counts[complement] > 0 {
			ops++
			counts[complement]--
		} else {
			counts[num]++
		}
	}
	return ops
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", maxOperations([]int{1, 2, 3, 4}, 5)) // Expected: 2

	// Test case 2
	fmt.Println("Test 2:", maxOperations([]int{3, 1, 3, 4, 3}, 6)) // Expected: 1

	// Test case 3
	fmt.Println("Test 3:", maxOperations([]int{1, 2, 3, 4, 5, 6}, 7)) // Expected: 3
}
```
