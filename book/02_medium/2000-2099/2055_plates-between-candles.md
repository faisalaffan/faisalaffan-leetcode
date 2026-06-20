# 2055 — Plates Between Candles

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func platesBetweenCandles(s string, queries [][]int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, Prefix Sum

**Waktu:** O(n + q)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2055: Plates Between Candles
// https://leetcode.com/problems/plates-between-candles/
// Difficulty: Medium
// Time: O(n + q) | Space: O(n)

import "fmt"

func platesBetweenCandles(s string, queries [][]int) []int {
	n := len(s)
	// Prefix sum of plates
  // Alokasi slice
	prefix := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i]
		if s[i] == '*' {
			prefix[i+1]++
		}
	}

	// Nearest candle to the left
  // Alokasi slice
	leftCandle := make([]int, n)
	last := -1
	for i := 0; i < n; i++ {
		if s[i] == '|' {
			last = i
		}
		leftCandle[i] = last
	}

	// Nearest candle to the right
  // Alokasi slice
	rightCandle := make([]int, n)
	last = -1
	for i := n - 1; i >= 0; i-- {
		if s[i] == '|' {
			last = i
		}
		rightCandle[i] = last
	}

  // Alokasi slice
	result := make([]int, len(queries))
	for i, q := range queries {
		left, right := q[0], q[1]
		l := rightCandle[left]
		r := leftCandle[right]
		if l == -1 || r == -1 || l >= r {
			result[i] = 0
		} else {
			result[i] = prefix[r] - prefix[l]
		}
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", platesBetweenCandles("**|**|***|", [][]int{{2, 5}, {5, 9}}))
	// Expected: [2, 3]

	// Test case 2
	fmt.Println("Test 2:", platesBetweenCandles("***|**|*****|**||**|*", [][]int{{1, 17}, {4, 5}, {14, 17}, {5, 11}, {15, 16}}))
	// Expected: [9, 0, 0, 0, 0]

	// Test case 3
	fmt.Println("Test 3:", platesBetweenCandles("|*|", [][]int{{0, 2}}))
	// Expected: [1]
}
```
