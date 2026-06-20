# 0740 — Delete And Earn

## Deskripsi

**Soal:** [0740. Delete And Earn](https://leetcode.com/problems/delete-and-earn/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n + k) where k is max value  
**Kompleksitas Ruang:** O(k)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #740: Delete and Earn
// https://leetcode.com/problems/delete-and-earn/
// Difficulty: Medium
// Time: O(n + k) where k is max value
// Space: O(k)

import "fmt"

func main() {
	fmt.Println(deleteAndEarn([]int{3, 4, 2}))
	fmt.Println(deleteAndEarn([]int{2, 2, 3, 3, 3, 4}))
}

func deleteAndEarn(nums []int) int {
  // Edge case: input kosong
	if len(nums) == 0 {
		return 0
	}

	maxVal := 0
	for _, n := range nums {
		if n > maxVal {
			maxVal = n
		}
	}

  // Membuat slice untuk menyimpan hasil
	values := make([]int, maxVal+1)
	for _, n := range nums {
		values[n] += n
	}

	prev2, prev1 := 0, values[1]
	for i := 2; i <= maxVal; i++ {
		curr := max(prev1, prev2+values[i])
		prev2, prev1 = prev1, curr
	}

	return prev1
}
```
