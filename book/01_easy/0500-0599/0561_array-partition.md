# 0561 — Array Partition

## Deskripsi

**Soal:** [0561. Array Partition](https://leetcode.com/problems/array-partition/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n log n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func ArrayPartition(nums []int) int`

## Solusi Go

```go
package main

// LeetCode #561: Array Partition
// https://leetcode.com/problems/array-partition/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

// Time: O(n log n), Space: O(1)
func ArrayPartition(nums []int) int {
	sort.Ints(nums)
	sum := 0
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(nums); i += 2 {
		sum += nums[i]
	}
	return sum
}

func main() {
	fmt.Println(ArrayPartition([]int{1, 4, 3, 2}))
	fmt.Println(ArrayPartition([]int{6, 2, 6, 5, 1, 2}))
}
```
