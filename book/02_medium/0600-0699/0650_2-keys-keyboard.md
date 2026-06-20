# 0650 — 2 Keys Keyboard

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func minSteps(n int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n sqrt(n))  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #650: 2 Keys Keyboard
// https://leetcode.com/problems/2-keys-keyboard/
// Difficulty: Medium
// Time: O(n sqrt(n))
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(minSteps(3))
	fmt.Println(minSteps(1))
	fmt.Println(minSteps(10))
}

func minSteps(n int) int {
	result := 0
	d := 2

	for n > 1 {
		for n%d == 0 {
			result += d
			n /= d
		}
		d++
	}

	return result
}
```
