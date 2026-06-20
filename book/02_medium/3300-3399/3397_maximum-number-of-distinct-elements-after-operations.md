# 3397 — Maximum Number Of Distinct Elements After Operations

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func maxDistinctElements(nums []int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n log n) Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3397: Maximum Number of Distinct Elements After Operations
// https://leetcode.com/problems/maximum-number-of-distinct-elements-after-operations/
// Difficulty: Medium
// Time: O(n log n) Space: O(1)

import (
	"fmt"
	"math"
	"slices"
)

func maxDistinctElements(nums []int, k int) int {
	n := len(nums)
	if k*2+1 >= n {
		return n
	}

	slices.Sort(nums)
	pre := math.MinInt
	ans := 0
	for _, x := range nums {
		candidate := max(x-k, pre+1)
		if candidate <= x+k {
			ans++
			pre = candidate
		}
	}
	return ans
}

func main() {
	fmt.Println(maxDistinctElements([]int{1, 2, 2, 3, 3, 4}, 2)) // 6
	fmt.Println(maxDistinctElements([]int{4, 4, 4, 4}, 1))        // 3
	fmt.Println(maxDistinctElements([]int{1, 1, 1, 1}, 0))        // 1
}
```
