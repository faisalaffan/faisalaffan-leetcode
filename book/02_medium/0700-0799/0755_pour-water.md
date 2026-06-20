# 0755 — Pour Water

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func pourWater(heights []int, volume int, k int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer

**Waktu:** O(V * N)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #755: Pour Water
// https://leetcode.com/problems/pour-water/
// Difficulty: Medium [Paid]
// Time: O(V * N)
// Space: O(1)

import "fmt"

func main() {
	heights := []int{2, 1, 1, 2, 1, 2, 2}
	result := pourWater(heights, 4, 3)
	fmt.Println(result)
}

func pourWater(heights []int, volume int, k int) []int {
	n := len(heights)

	for v := 0; v < volume; v++ {
		pos := k

		// Try left
		left := k
		for left > 0 && heights[left] >= heights[left-1] {
			left--
		}
		for left < k && heights[left] == heights[left+1] {
			left++
		}
		if heights[left] < heights[pos] {
			pos = left
		}

		if pos == k {
			// Try right
			right := k
			for right < n-1 && heights[right] >= heights[right+1] {
				right++
			}
			for right > k && heights[right] == heights[right-1] {
				right--
			}
			if heights[right] < heights[pos] {
				pos = right
			}
		}

		heights[pos]++
	}

	return heights
}
```
