# 2393 — Count Strictly Increasing Subarrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func countIncreasing(nums []int) int64`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2393: Count Strictly Increasing Subarrays
// https://leetcode.com/problems/count-strictly-increasing-subarrays/
// Difficulty: Medium
// Time: O(n) | Space: O(1)
// Track length of current increasing run. Each position adds run_len subarrays.

import "fmt"

func main() {
	fmt.Println(countIncreasing([]int{1, 3, 5, 4, 4, 6})) // 10
	fmt.Println(countIncreasing([]int{1, 2, 3, 4, 5}))    // 15
}

func countIncreasing(nums []int) int64 {
	var ans int64
	run := 0
  // Linear scan O(n)
	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i] > nums[i-1] {
			run++
		} else {
			run = 1
		}
		ans += int64(run)
	}
	return ans
}
```
