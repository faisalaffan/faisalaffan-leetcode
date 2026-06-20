# 3649 — Number Of Perfect Pairs

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func numberOfPerfectPairs(nums []int) int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3649: Number of Perfect Pairs
// https://leetcode.com/problems/number-of-perfect-pairs/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func numberOfPerfectPairs(nums []int) int64 {
	n := len(nums)
  // Alokasi slice
	arr := make([]int, n)
	for i, v := range nums {
		if v < 0 {
			arr[i] = -v
		} else {
			arr[i] = v
		}
	}

  // Sort O(n log n)
	sort.Ints(arr)

	var ans int64 = 0
	j := 0
	for i := 1; i < n; i++ {
		for 2*arr[j] < arr[i] {
			j++
		}
		ans += int64(i - j)
	}
	return ans
}

func main() {
	fmt.Println(numberOfPerfectPairs([]int{1, 2, 3, 4}))
	fmt.Println(numberOfPerfectPairs([]int{-1, 1, -2, 2}))
	fmt.Println(numberOfPerfectPairs([]int{5, 1, 2}))
}
```
