# 0180 — Consecutive Numbers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func consecutiveNumbers(nums []int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #180: Consecutive Numbers
// https://leetcode.com/problems/consecutive-numbers/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

func consecutiveNumbers(nums []int) []int {
	if len(nums) < 3 {
		return nil
	}

	result := []int{}
  // HashMap: O(1) lookup
	seen := make(map[int]bool)

  // Linear scan O(n)
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] == nums[i+1] && nums[i] == nums[i+2] && !seen[nums[i]] {
			result = append(result, nums[i])
			seen[nums[i]] = true
		}
	}

	return result
}

func main() {
	fmt.Println(consecutiveNumbers([]int{1, 1, 1, 2, 2, 3, 3, 3}))
	fmt.Println(consecutiveNumbers([]int{1, 2, 3, 4}))
	fmt.Println(consecutiveNumbers([]int{1, 1, 1, 1, 2, 2, 2}))
}
```
