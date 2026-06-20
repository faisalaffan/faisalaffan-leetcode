# 1646 — Get Maximum In Generated Array

## Deskripsi

**Soal:** [1646. Get Maximum In Generated Array](https://leetcode.com/problems/get-maximum-in-generated-array/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func GetMaximumGenerated(n int) int`

## Solusi Go

```go
package main

// LeetCode #1646: Get Maximum in Generated Array
// https://leetcode.com/problems/get-maximum-in-generated-array/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(n)
func GetMaximumGenerated(n int) int {
  // Edge case: input kosong
	if n == 0 {
		return 0
	}
  // Membuat slice untuk menyimpan hasil
	nums := make([]int, n+1)
	nums[0] = 0
	nums[1] = 1
	maxVal := 1
	for i := 2; i <= n; i++ {
		if i%2 == 0 {
			nums[i] = nums[i/2]
		} else {
			nums[i] = nums[i/2] + nums[i/2+1]
		}
		if nums[i] > maxVal {
			maxVal = nums[i]
		}
	}
	return maxVal
}

func main() {
	fmt.Println(GetMaximumGenerated(7))
	fmt.Println(GetMaximumGenerated(2))
	fmt.Println(GetMaximumGenerated(3))
}
```
