# 1310 — Xor Queries Of A Subarray

## Deskripsi

**Soal:** [1310. Xor Queries Of A Subarray](https://leetcode.com/problems/xor-queries-of-a-subarray/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n + m) where n = len(arr), m = len(queries)  
**Kompleksitas Ruang:** O(n) for prefix XOR array

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1310: XOR Queries of a Subarray
// https://leetcode.com/problems/xor-queries-of-a-subarray/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(xorQueries([]int{1, 3, 4, 8}, [][]int{{0, 1}, {1, 2}, {0, 3}, {3, 3}}))
	// [2,7,14,8]

	// Test case 2
	fmt.Println(xorQueries([]int{4, 8, 2, 10}, [][]int{{2, 3}, {1, 3}, {0, 0}, {0, 3}}))
	// [8,0,4,4]

	// Test case 3
	fmt.Println(xorQueries([]int{2}, [][]int{{0, 0}}))
	// [2]
}

// Time: O(n + m) where n = len(arr), m = len(queries)
// Space: O(n) for prefix XOR array
func xorQueries(arr []int, queries [][]int) []int {
	n := len(arr)
  // Membuat slice untuk menyimpan hasil
	prefix := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] ^ arr[i]
	}

  // Membuat slice untuk menyimpan hasil
	result := make([]int, len(queries))
	for i, q := range queries {
		result[i] = prefix[q[1]+1] ^ prefix[q[0]]
	}
	return result
}
```
