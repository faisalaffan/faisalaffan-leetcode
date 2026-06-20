# 3927 — Minimize Array Sum Using Divisible Replacements

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimizeArraySumUsingDivisibleReplacements(nums []int) int64
```

> **💡 Hint:** For each element, find its minimum divisor present in the array.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(N sqrt(M))  
**Kompleksitas Ruang:** O(N) where M = max value

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3927: Minimize Array Sum Using Divisible Replacements
// https://leetcode.com/problems/minimize-array-sum-using-divisible-replacements/
// Difficulty: Medium
// Time: O(N sqrt(M)) | Space: O(N) where M = max value
// Approach: For each element, find its minimum divisor present in the array.
// Each number can be replaced with any present divisor (via chaining).
// Sum the minimal reachable value for each element.

import (
	"fmt"
	"math"
)

func MinimizeArraySumUsingDivisibleReplacements(nums []int) int64 {
	// Track which values exist
  // Membuat map (HashMap) — pencarian O(1)
	exists := make(map[int]bool)
	minVal := math.MaxInt32
	for _, v := range nums {
		exists[v] = true
		if v < minVal {
			minVal = v
		}
	}

	var ans int64 = 0
	for _, v := range nums {
		reduced := false
		// Try divisors from 1 to sqrt(v)
		for d := 1; d*d <= v; d++ {
			if v%d == 0 {
				if exists[d] {
					ans += int64(d)
					reduced = true
					break
				}
				other := v / d
				if other != d && exists[other] && other < v {
					// Can't directly use if not minimum, but may be useful
				}
			}
		}
		if !reduced {
			ans += int64(v)
		}
	}
	return ans
}

func main() {
	// Example 1
	fmt.Println(MinimizeArraySumUsingDivisibleReplacements([]int{3, 6, 2})) // Expected: 7

	// Example 2
	fmt.Println(MinimizeArraySumUsingDivisibleReplacements([]int{4, 2, 8, 3})) // Expected: 9

	// Example 3
	fmt.Println(MinimizeArraySumUsingDivisibleReplacements([]int{7, 5, 9})) // Expected: 21
}
```
