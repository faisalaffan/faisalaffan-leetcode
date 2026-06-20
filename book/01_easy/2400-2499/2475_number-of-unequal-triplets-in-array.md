# 2475 — Number Of Unequal Triplets In Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func NumberOfUnequalTripletsInArray(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #2475: Number of Unequal Triplets in Array
// https://leetcode.com/problems/number-of-unequal-triplets-in-array/
// Difficulty: Easy
// Time O(n^3) | Space O(1)

import "fmt"

func main() {
	fmt.Println(NumberOfUnequalTripletsInArray([]int{4, 4, 2, 4, 3})) // 3
	fmt.Println(NumberOfUnequalTripletsInArray([]int{1, 1, 1, 1, 1}))  // 0
}

func NumberOfUnequalTripletsInArray(nums []int) int {
	n := len(nums)
	count := 0
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				if nums[i] != nums[j] && nums[i] != nums[k] && nums[j] != nums[k] {
					count++
				}
			}
		}
	}
	return count
}
```
