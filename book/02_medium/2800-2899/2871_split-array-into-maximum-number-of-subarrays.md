# 2871 — Split Array Into Maximum Number Of Subarrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func SplitArrayIntoMaximumNumberOfSubarrays(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2871: Split Array Into Maximum Number of Subarrays
// https://leetcode.com/problems/split-array-into-maximum-number-of-subarrays/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func SplitArrayIntoMaximumNumberOfSubarrays(nums []int) int {
	// We need to split such that AND of each subarray's bitwise AND is minimum
	// Minimum possible AND of any subarray is the AND of entire array
	minAnd := nums[0]
	for _, v := range nums[1:] {
		minAnd &= v
	}

	if minAnd != 0 {
		return 1
	}

	// Count subarrays whose AND equals 0
	count := 0
	curAnd := nums[0]
	for _, v := range nums[1:] {
		if curAnd == 0 {
			count++
			curAnd = v
		} else {
			curAnd &= v
		}
	}
	if curAnd == 0 {
		count++
	}

	return count
}

func main() {
	fmt.Println(SplitArrayIntoMaximumNumberOfSubarrays([]int{1, 0, 2, 0, 1, 0}))
	fmt.Println(SplitArrayIntoMaximumNumberOfSubarrays([]int{1, 2, 3}))
}
```
