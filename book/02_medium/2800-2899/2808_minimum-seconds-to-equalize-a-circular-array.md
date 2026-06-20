# 2808 — Minimum Seconds To Equalize A Circular Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func MinimumSecondsToEqualizeACircularArray(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2808: Minimum Seconds to Equalize a Circular Array
// https://leetcode.com/problems/minimum-seconds-to-equalize-a-circular-array/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func MinimumSecondsToEqualizeACircularArray(nums []int) int {
	n := len(nums)
  // HashMap: O(1) lookup
	pos := make(map[int][]int)
	for i, v := range nums {
		pos[v] = append(pos[v], i)
	}

	best := n / 2
	for _, positions := range pos {
		if len(positions) == 0 {
			continue
		}
		maxGap := 0
  // Linear scan O(n)
		for i := 0; i < len(positions); i++ {
			curr := positions[i]
			var prev int
			if i > 0 {
				prev = positions[i-1]
			} else {
				prev = positions[len(positions)-1] - n
			}
			gap := curr - prev
			if gap > maxGap {
				maxGap = gap
			}
		}
		seconds := maxGap / 2
		if seconds < best {
			best = seconds
		}
	}

	return best
}

func main() {
	fmt.Println(MinimumSecondsToEqualizeACircularArray([]int{1, 2, 1, 2}))
	fmt.Println(MinimumSecondsToEqualizeACircularArray([]int{2, 1, 3, 3, 2}))
}
```
