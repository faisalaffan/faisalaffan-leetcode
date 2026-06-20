# 1509 — Minimum Difference Between Largest And Smallest Value In Three Moves

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MinDifference(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(N log N), Space: O(1) if ignoring sort space  
**Kompleksitas Ruang:** O(1) if ignoring sort space

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1509: Minimum Difference Between Largest and Smallest Value in Three Moves
// https://leetcode.com/problems/minimum-difference-between-largest-and-smallest-value-in-three-moves/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MinDifference([]int{5, 3, 2, 4}))
	fmt.Println(MinDifference([]int{1, 5, 0, 10, 14}))
	fmt.Println(MinDifference([]int{3, 100, 20}))
}

func MinDifference(nums []int) int {
	// Time: O(N log N), Space: O(1) if ignoring sort space
	if len(nums) <= 4 {
		return 0
	}

  // Urutkan secara ascending — O(n log n)
	sort.Ints(nums)
	n := len(nums)

	// After 3 moves, we can change up to 3 values.
	// The minimum difference will be between some combination
	// of removing 0-3 from left and 3-0 from right.
	minDiff := nums[n-1] - nums[0]
	for i := 0; i <= 3; i++ {
		diff := nums[n-1-(3-i)] - nums[i]
		if diff < minDiff {
			minDiff = diff
		}
	}

	return minDiff
}
```
