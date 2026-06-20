# 0274 — H Index

## Deskripsi

**Soal:** [0274. H Index](https://leetcode.com/problems/h-index/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func hIndex(citations []int) int`

## Solusi Go

```go
package main

// LeetCode #274: H-Index
// https://leetcode.com/problems/h-index/
// Difficulty: Medium
// Time: O(n), Space: O(n)

import "fmt"

func hIndex(citations []int) int {
	n := len(citations)
  // Membuat slice untuk menyimpan hasil
	buckets := make([]int, n+1)

	for _, c := range citations {
		if c >= n {
			buckets[n]++
		} else {
			buckets[c]++
		}
	}

	count := 0
	for i := n; i >= 0; i-- {
		count += buckets[i]
		if count >= i {
			return i
		}
	}

	return 0
}

func main() {
	fmt.Println(hIndex([]int{3, 0, 6, 1, 5}))
	fmt.Println(hIndex([]int{1, 3, 1}))
	fmt.Println(hIndex([]int{0}))
}
```
