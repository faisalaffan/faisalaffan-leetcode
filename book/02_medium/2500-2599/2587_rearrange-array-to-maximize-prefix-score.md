# 2587 — Rearrange Array To Maximize Prefix Score

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxScore(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2587: Rearrange Array to Maximize Prefix Score
// https://leetcode.com/problems/rearrange-array-to-maximize-prefix-score/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func maxScore(nums []int) int {
  // Custom sort dengan comparator
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	var sum int64
	count := 0
	for _, v := range nums {
		sum += int64(v)
		if sum > 0 {
			count++
		} else {
			break
		}
	}
	return count
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", maxScore([]int{2, -1, 0, 1, -3, 3, -3}))
	// Expected: 6

	// Test case 2
	fmt.Println("Test 2:", maxScore([]int{-2, -3, 0}))
	// Expected: 0

	// Test case 3
	fmt.Println("Test 3:", maxScore([]int{1, 2, 3, 4, 5}))
	// Expected: 5
}
```
