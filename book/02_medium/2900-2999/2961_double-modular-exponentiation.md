# 2961 — Double Modular Exponentiation

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func getGoodIndices(variables [][]int, target int) (ans []int)`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n log m)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2961: Double Modular Exponentiation
// https://leetcode.com/problems/double-modular-exponentiation/
// Difficulty: Medium
// Time: O(n log m) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(getGoodIndices([][]int{{2, 3, 3, 10}, {3, 3, 3, 1}, {6, 1, 1, 4}}, 2))
	fmt.Println(getGoodIndices([][]int{{39, 3, 1000, 1000}}, 17))
}

func getGoodIndices(variables [][]int, target int) (ans []int) {
	qpow := func(a, n, mod int) int {
		ans := 1
		for ; n > 0; n >>= 1 {
			if n&1 == 1 {
				ans = ans * a % mod
			}
			a = a * a % mod
		}
		return ans
	}
	for i, e := range variables {
		a, b, c, m := e[0], e[1], e[2], e[3]
		if qpow(qpow(a, b, 10), c, m) == target {
			ans = append(ans, i)
		}
	}
	return
}
```
