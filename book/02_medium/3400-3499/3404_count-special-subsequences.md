# 3404 — Count Special Subsequences

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func numberOfSubsequences(nums []int) int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n^2) Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3404: Count Special Subsequences
// https://leetcode.com/problems/count-special-subsequences/
// Difficulty: Medium
// Time: O(n^2) Space: O(n)

import "fmt"

func numberOfSubsequences(nums []int) int64 {
	n := len(nums)
	var ans int64
  // HashMap: O(1) lookup
	cnt := make(map[float64]int)

	// For each r, q = r-2. Accumulate (p,q) pairs as r increases.
	for r := 4; r < n-2; r++ {
		q := r - 2
		b := float64(nums[q])
		for _, aVal := range nums[:q-1] {
			ratio := float64(aVal) / b
			cnt[ratio]++
		}

		c := float64(nums[r])
		for _, dVal := range nums[r+2:] {
			ratio := float64(dVal) / c
			ans += int64(cnt[ratio])
		}
	}
	return ans
}

func main() {
	fmt.Println(numberOfSubsequences([]int{1, 2, 3, 4, 3, 6, 1})) // 1
	fmt.Println(numberOfSubsequences([]int{3, 4, 3, 4, 3, 4, 3, 4})) // 3
}
```
