# 1503 — Last Moment Before All Ants Fall Out Of A Plank

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func GetLastMoment(n int, left []int, right []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(N), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1503: Last Moment Before All Ants Fall Out of a Plank
// https://leetcode.com/problems/last-moment-before-all-ants-fall-out-of-a-plank/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(GetLastMoment(4, []int{4, 3}, []int{0, 1}))
	fmt.Println(GetLastMoment(7, []int{}, []int{0, 1, 2, 3, 4, 5, 6, 7}))
	fmt.Println(GetLastMoment(7, []int{0, 1, 2, 3, 4, 5, 6, 7}, []int{}))
}

func GetLastMoment(n int, left []int, right []int) int {
	// Time: O(N), Space: O(1)
	// When ants meet, they reverse direction. This is equivalent to
	// ants passing through each other (identity swap).
	// So the last moment is max of:
	//   - ants moving left: their starting position (time to reach 0)
	//   - ants moving right: n - their starting position (time to reach n)
	maxTime := 0

	for _, pos := range left {
		if pos > maxTime {
			maxTime = pos
		}
	}

	for _, pos := range right {
		t := n - pos
		if t > maxTime {
			maxTime = t
		}
	}

	return maxTime
}
```
