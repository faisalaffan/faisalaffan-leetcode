# 1014 — Best Sightseeing Pair

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func maxScoreSightseeingPair(values []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1014: Best Sightseeing Pair
// https://leetcode.com/problems/best-sightseeing-pair/
// Difficulty: Medium
//
// Approach: Track max value of (values[i] + i) seen so far
// Score = values[i] + values[j] + i - j = (values[i] + i) + (values[j] - j)
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(maxScoreSightseeingPair([]int{8, 1, 5, 2, 6})) // 11
	fmt.Println(maxScoreSightseeingPair([]int{1, 2}))           // 2
}

func maxScoreSightseeingPair(values []int) int {
	maxI := values[0]
	result := 0

	for j := 1; j < len(values); j++ {
		if maxI+values[j]-j > result {
			result = maxI + values[j] - j
		}
		if values[j]+j > maxI {
			maxI = values[j] + j
		}
	}

	return result
}
```
