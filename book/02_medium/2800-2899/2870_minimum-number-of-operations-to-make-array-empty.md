# 2870 — Minimum Number Of Operations To Make Array Empty

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func MinimumNumberOfOperationsToMakeArrayEmpty(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2870: Minimum Number of Operations to Make Array Empty
// https://leetcode.com/problems/minimum-number-of-operations-to-make-array-empty/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func MinimumNumberOfOperationsToMakeArrayEmpty(nums []int) int {
  // HashMap: O(1) lookup
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}

	ops := 0
	for _, c := range freq {
		if c == 1 {
			return -1
		}
		// Use as many 3s as possible
		ops += c / 3
		if c%3 != 0 {
			ops++
		}
	}

	return ops
}

func main() {
	fmt.Println(MinimumNumberOfOperationsToMakeArrayEmpty([]int{2, 3, 3, 3, 3, 2}))
	fmt.Println(MinimumNumberOfOperationsToMakeArrayEmpty([]int{1, 1, 1, 1}))
	fmt.Println(MinimumNumberOfOperationsToMakeArrayEmpty([]int{1, 2, 3}))
}
```
