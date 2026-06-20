# 1814 — Count Nice Pairs In An Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func countNicePairs(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n log M) where M = max digit length, Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1814: Count Nice Pairs in an Array
// https://leetcode.com/problems/count-nice-pairs-in-an-array/
// Difficulty: Medium
// Time: O(n log M) where M = max digit length, Space: O(n)

import "fmt"

const mod = 1_000_000_007

func countNicePairs(nums []int) int {
  // HashMap: O(1) lookup
	count := make(map[int]int)
	result := 0

	for _, v := range nums {
		key := v - rev(v)
		result = (result + count[key]) % mod
		count[key]++
	}
	return result
}

func rev(x int) int {
	r := 0
	for x > 0 {
		r = r*10 + x%10
		x /= 10
	}
	return r
}

func main() {
	fmt.Println(countNicePairs([]int{42, 11, 1, 97})) // Expected: 2
	fmt.Println(countNicePairs([]int{13, 10, 35, 24, 76})) // Expected: 4
	fmt.Println(countNicePairs([]int{1, 1, 1, 1})) // Expected: 6
}
```
