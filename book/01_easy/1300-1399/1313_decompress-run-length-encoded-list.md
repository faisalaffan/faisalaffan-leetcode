# 1313 — Decompress Run Length Encoded List

## Deskripsi

**Soal:** [1313. Decompress Run Length Encoded List](https://leetcode.com/problems/decompress-run-length-encoded-list/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n + totalLen), Space: O(totalLen)  
**Kompleksitas Ruang:** O(totalLen)

**Algoritma:** LIS (Longest Increasing Subsequence)

**Fungsi Solusi:** `func decompressRLElist(nums []int) []int`

## Solusi Go

```go
package main

// LeetCode #1313: Decompress Run-Length Encoded List
// https://leetcode.com/problems/decompress-run-length-encoded-list/
// Difficulty: Easy
//
// LeetCode submission: func decompressRLElist(nums []int) []int

import "fmt"

func main() {
	fmt.Println(DecompressRunLengthEncodedList([]int{1, 2, 3, 4}))       // [2 4 4 4]
	fmt.Println(DecompressRunLengthEncodedList([]int{1, 1, 2, 3}))       // [1 3 3]
	fmt.Println(DecompressRunLengthEncodedList([]int{2, 5, 1, 7, 3, 9})) // [5 5 7 9 9 9]
}

// Time: O(n + totalLen), Space: O(totalLen)
func DecompressRunLengthEncodedList(nums []int) []int {
	totalLen := 0
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(nums); i += 2 {
		totalLen += nums[i]
	}
  // Membuat slice untuk menyimpan hasil
	res := make([]int, 0, totalLen)
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(nums); i += 2 {
		freq, val := nums[i], nums[i+1]
		for j := 0; j < freq; j++ {
			res = append(res, val)
		}
	}
	return res
}
```
