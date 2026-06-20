# 3400 — Maximum Number Of Matching Indices After Right Shifts

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func maximumMatchingIndices(nums []int, x int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n^2) Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3400: Maximum Number of Matching Indices After Right Shifts
// https://leetcode.com/problems/maximum-number-of-matching-indices-after-right-shifts/
// Difficulty: Medium [Paid]
// Time: O(n^2) Space: O(1)

import "fmt"

func maximumMatchingIndices(nums []int, x int) int {
	n := len(nums)
	ans := 0
	for shift := 0; shift < n; shift++ {
		cnt := 0
		for i := 0; i < n; i++ {
			j := (i - shift + n) % n
			if nums[i] == nums[j] {
				cnt++
			}
		}
		if cnt > ans {
			ans = cnt
		}
	}
	return ans
}

func main() {
	fmt.Println(maximumMatchingIndices([]int{1, 2, 3, 4}, 0)) // 4
	fmt.Println(maximumMatchingIndices([]int{1, 2, 3, 1}, 0)) // 3
	fmt.Println(maximumMatchingIndices([]int{1, 1, 1, 1}, 0)) // 4
}
```
