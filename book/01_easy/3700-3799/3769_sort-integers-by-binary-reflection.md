# 3769 — Sort Integers By Binary Reflection

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func SortIntegersByBinaryReflection(nums []int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3769: Sort Integers by Binary Reflection
// https://leetcode.com/problems/sort-integers-by-binary-reflection/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(SortIntegersByBinaryReflection([]int{3, 6, 5}))
	fmt.Println(SortIntegersByBinaryReflection([]int{1, 2, 3}))
}

// Time: O(n log n)
// Space: O(n)
func SortIntegersByBinaryReflection(nums []int) []int {
  // HashMap: O(1) lookup
	reflections := make(map[int]int)
	for _, v := range nums {
		reflections[v] = reverseBits(v)
	}

  // Custom sort
	sort.Slice(nums, func(i, j int) bool {
		if reflections[nums[i]] != reflections[nums[j]] {
			return reflections[nums[i]] < reflections[nums[j]]
		}
		return nums[i] < nums[j]
	})
	return nums
}

func reverseBits(n int) int {
	rev := 0
	for n > 0 {
		rev = (rev << 1) | (n & 1)
		n >>= 1
	}
	return rev
}
```
