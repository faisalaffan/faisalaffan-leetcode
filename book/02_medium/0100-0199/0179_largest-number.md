# 0179 — Largest Number

## Deskripsi

**Soal:** [0179. Largest Number](https://leetcode.com/problems/largest-number/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func largestNumber(nums []int) string`

## Solusi Go

```go
package main

// LeetCode #179: Largest Number
// https://leetcode.com/problems/largest-number/
// Difficulty: Medium
// Time: O(n log n), Space: O(n)

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
  // Membuat slice untuk menyimpan hasil
	strs := make([]string, len(nums))
	for i, num := range nums {
		strs[i] = strconv.Itoa(num)
	}

	sort.Slice(strs, func(i, j int) bool {
		return strs[i]+strs[j] > strs[j]+strs[i]
	})

	if strs[0] == "0" {
		return "0"
	}

	return strings.Join(strs, "")
}

func main() {
	fmt.Println(largestNumber([]int{10, 2}))
	fmt.Println(largestNumber([]int{3, 30, 34, 5, 9}))
	fmt.Println(largestNumber([]int{0, 0}))
}
```
