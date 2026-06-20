# 2438 — Range Product Queries Of Powers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func productQueries(n int, queries [][]int) []int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n + q * n)  |  **Ruang:** O(log n)


## 💻 Solusi Go

```go
package main

// LeetCode #2438: Range Product Queries of Powers
// https://leetcode.com/problems/range-product-queries-of-powers/
// Difficulty: Medium
// Time: O(n + q * n) | Space: O(log n)
// Decompose n into powers of 2. For each query, multiply range.

import "fmt"

func main() {
	fmt.Println(productQueries(15, [][]int{{0, 1}, {2, 2}, {0, 3}})) // [2,4,64]
	fmt.Println(productQueries(2, [][]int{{0, 0}}))                  // [2]
}

const MOD = 1000000007

func productQueries(n int, queries [][]int) []int {
  // Alokasi slice
	powers := make([]int, 0)
	pow := 1
	for n > 0 {
		if n&1 == 1 {
			powers = append(powers, pow)
		}
		n >>= 1
		pow <<= 1
	}

  // Alokasi slice
	ans := make([]int, len(queries))
	for i, q := range queries {
		prod := 1
		for j := q[0]; j <= q[1]; j++ {
			prod = (prod * powers[j]) % MOD
		}
		ans[i] = prod
	}
	return ans
}
```
