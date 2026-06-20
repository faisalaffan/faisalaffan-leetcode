# 2677 — Chunk Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func ChunkArray(arr []int, size int) [][]int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

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
  // Linear scan O(n)
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
