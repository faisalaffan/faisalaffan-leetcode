# 0180 — Consecutive Numbers

## Deskripsi

**Soal:** [0180. Consecutive Numbers](https://leetcode.com/problems/consecutive-numbers/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func consecutiveNumbers(nums []int) []int`

## Solusi Go

```go
package main

// LeetCode #180: Consecutive Numbers
// https://leetcode.com/problems/consecutive-numbers/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

func consecutiveNumbers(nums []int) []int {
	if len(nums) < 3 {
		return nil
	}

	result := []int{}
  // Membuat map untuk pencarian O(1): key → value
	seen := make(map[int]bool)

  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] == nums[i+1] && nums[i] == nums[i+2] && !seen[nums[i]] {
			result = append(result, nums[i])
			seen[nums[i]] = true
		}
	}

	return result
}

func main() {
	fmt.Println(consecutiveNumbers([]int{1, 1, 1, 2, 2, 3, 3, 3}))
	fmt.Println(consecutiveNumbers([]int{1, 2, 3, 4}))
	fmt.Println(consecutiveNumbers([]int{1, 1, 1, 1, 2, 2, 2}))
}
```
