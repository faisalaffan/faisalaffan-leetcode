# 0683 — K Empty Slots

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func kEmptySlots(bulbs []int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, Sliding Window

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
	"math"
)

// LeetCode #683: K Empty Slots
// https://leetcode.com/problems/k-empty-slots/
// Difficulty: Hard
//
// bulbs[i] = position that lights up on day i (1-indexed).
// Find earliest day with two lit bulbs having exactly k unlit bulbs between them.
// Sliding window over positions using day-of-lighting.

func main() {
	// [1,3,2], k=1 => 2 (day 2: bulbs at 1 and 3 are on)
	fmt.Println(kEmptySlots([]int{1, 3, 2}, 1))
	// [1,2,3], k=1 => -1
	fmt.Println(kEmptySlots([]int{1, 2, 3}, 1))
	// [3,1,5,4,2], k=1 => 4 (day 4: bulbs at 3 and 5 on, 4 between)
	fmt.Println(kEmptySlots([]int{3, 1, 5, 4, 2}, 1))
	// [6,5,8,9,7,1,10,2,3,4], k=2 => 8
	fmt.Println(kEmptySlots([]int{6, 5, 8, 9, 7, 1, 10, 2, 3, 4}, 2))
	// [2,1,3], k=1 => 2
	fmt.Println(kEmptySlots([]int{2, 1, 3}, 1))
}

func kEmptySlots(bulbs []int, k int) int {
	n := len(bulbs)
  // Alokasi slice
	day := make([]int, n)
	for i, pos := range bulbs {
		day[pos-1] = i + 1
	}

	result := math.MaxInt32
	left, right := 0, k+1

	for right < n {
		i := left + 1
		valid := true
		for i < right {
			if day[i] < day[left] || day[i] < day[right] {
				valid = false
				left = i
				right = i + k + 1
				break
			}
			i++
		}
		if valid {
			cur := max(day[left], day[right])
			if cur < result {
				result = cur
			}
			left = right
			right = left + k + 1
		}
	}

	if result == math.MaxInt32 {
		return -1
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```
