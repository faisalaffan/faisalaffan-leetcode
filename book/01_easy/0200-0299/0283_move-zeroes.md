# 0283 — Move Zeroes

## Deskripsi

**Soal:** [0283. Move Zeroes](https://leetcode.com/problems/move-zeroes/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func MoveZeroes(nums []int) `

## Solusi Go

```go
package main

// LeetCode #283: Move Zeroes
// https://leetcode.com/problems/move-zeroes/
// Difficulty: Easy

import "fmt"

// Time: O(n) | Space: O(1)
func MoveZeroes(nums []int) {
	lastNonZero := 0
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[lastNonZero] = nums[lastNonZero], nums[i]
			lastNonZero++
		}
	}
}

func main() {
	n1 := []int{0, 1, 0, 3, 12}
	MoveZeroes(n1)
	fmt.Println(n1)
	n2 := []int{0}
	MoveZeroes(n2)
	fmt.Println(n2)
}
```
