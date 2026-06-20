# 2677 — Chunk Array

## Deskripsi

**Soal:** [2677. Chunk Array](https://leetcode.com/problems/chunk-array/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2677: Chunk Array
// https://leetcode.com/problems/chunk-array/
// Difficulty: Easy
// Time: O(n) | Space: O(n)
// Note: JavaScript problem, adapted to Go. Splits array into chunks of given size.

import "fmt"

func main() {
	fmt.Println(ChunkArray([]int{1, 2, 3, 4, 5}, 1))
	fmt.Println(ChunkArray([]int{1, 9, 6, 3, 2}, 3))
}

func ChunkArray(arr []int, size int) [][]int {
	var result [][]int
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(arr); i += size {
		end := i + size
		if end > len(arr) {
			end = len(arr)
		}
		result = append(result, arr[i:end])
	}
	return result
}
```
