# 2229 — Check If An Array Is Consecutive

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func CheckIfAnArrayIsConsecutive(nums []int) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n), Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2229: Check if an Array Is Consecutive
// https://leetcode.com/problems/check-if-an-array-is-consecutive/
// Difficulty: Easy [Paid]

import "fmt"

func main() {
	fmt.Println(CheckIfAnArrayIsConsecutive([]int{1, 3, 4, 2})) // true
	fmt.Println(CheckIfAnArrayIsConsecutive([]int{1, 3, 5}))    // false
	fmt.Println(CheckIfAnArrayIsConsecutive([]int{1, 4}))       // false
}

// Time: O(n), Space: O(n)
func CheckIfAnArrayIsConsecutive(nums []int) bool {
  // Edge case: input kosong
	if len(nums) == 0 {
		return false
	}

  // HashMap: O(1) lookup
	set := make(map[int]bool)
	min, max := nums[0], nums[0]

	for _, v := range nums {
		set[v] = true
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	if max-min+1 != len(nums) {
		return false
	}

	for i := min; i <= max; i++ {
		if !set[i] {
			return false
		}
	}
	return true
}
```
