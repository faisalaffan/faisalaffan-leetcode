# 0228 — Summary Ranges

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func SummaryRanges(nums []int) []string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1) excluding output


## 💻 Solusi Go

```go
package main

// LeetCode #228: Summary Ranges
// https://leetcode.com/problems/summary-ranges/
// Difficulty: Easy

import (
	"fmt"
	"strconv"
)

// Time: O(n) | Space: O(1) excluding output
func SummaryRanges(nums []int) []string {
	var res []string
	i := 0
	for i < len(nums) {
		start := nums[i]
		for i+1 < len(nums) && nums[i+1] == nums[i]+1 {
			i++
		}
		if start == nums[i] {
			res = append(res, strconv.Itoa(start))
		} else {
			res = append(res, strconv.Itoa(start)+"->"+strconv.Itoa(nums[i]))
		}
		i++
	}
	return res
}

func main() {
	fmt.Println(SummaryRanges([]int{0, 1, 2, 4, 5, 7}))
	fmt.Println(SummaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))
	fmt.Println(SummaryRanges([]int{}))
}
```
