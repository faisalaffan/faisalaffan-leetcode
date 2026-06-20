# 0985 — Sum Of Even Numbers After Queries

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func sumEvenAfterQueries(nums []int, queries [][]int) []int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n + q) where n = len(nums), q = len(queries)  |  **Ruang:** O(1) excluding output


## 💻 Solusi Go

```go
package main

// LeetCode #985: Sum of Even Numbers After Queries
// https://leetcode.com/problems/sum-of-even-numbers-after-queries/
// Difficulty: Medium
//
// Approach: Maintain running sum of even numbers, update incrementally
// Time: O(n + q) where n = len(nums), q = len(queries)
// Space: O(1) excluding output

import "fmt"

func main() {
	fmt.Println(sumEvenAfterQueries([]int{1, 2, 3, 4}, [][]int{{1, 0}, {-3, 1}, {-4, 0}, {2, 3}})) // [8,6,2,4]
	fmt.Println(sumEvenAfterQueries([]int{1}, [][]int{{4, 0}}))                                  // [0]
}

func sumEvenAfterQueries(nums []int, queries [][]int) []int {
  // Alokasi slice
	result := make([]int, len(queries))
	sum := 0

	for _, n := range nums {
		if n%2 == 0 {
			sum += n
		}
	}

	for i, q := range queries {
		val, idx := q[0], q[1]
		if nums[idx]%2 == 0 {
			sum -= nums[idx]
		}
		nums[idx] += val
		if nums[idx]%2 == 0 {
			sum += nums[idx]
		}
		result[i] = sum
	}

	return result
}
```
