# 2815 — Max Pair Sum In An Array

## Deskripsi

**Soal:** [2815. Max Pair Sum In An Array](https://leetcode.com/problems/max-pair-sum-in-an-array/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2815: Max Pair Sum in an Array
// https://leetcode.com/problems/max-pair-sum-in-an-array/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(MaxPairSumInAnArray([]int{51, 71, 17, 24, 42}))
	fmt.Println(MaxPairSumInAnArray([]int{1, 2, 3, 4}))
}

func maxDigit(n int) int {
	maxD := 0
	for n > 0 {
		d := n % 10
		if d > maxD {
			maxD = d
		}
		n /= 10
	}
	return maxD
}

func MaxPairSumInAnArray(nums []int) int {
  // Membuat slice untuk menyimpan hasil
	maxVal := make([]int, 10) // digits 0-9
	ans := -1
	for _, n := range nums {
		md := maxDigit(n)
		if maxVal[md] > 0 {
			sum := maxVal[md] + n
			if sum > ans {
				ans = sum
			}
		}
		if n > maxVal[md] {
			maxVal[md] = n
		}
	}
	return ans
}
```
