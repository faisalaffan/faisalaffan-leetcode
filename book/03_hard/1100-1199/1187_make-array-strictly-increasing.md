# 1187 — Make Array Strictly Increasing

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func makeArrayIncreasing(arr1 []int, arr2 []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1187: Make Array Strictly Increasing
// https://leetcode.com/problems/make-array-strictly-increasing/
// Difficulty: Hard

import (
	"fmt"
	"math"
	"sort"
)

func makeArrayIncreasing(arr1 []int, arr2 []int) int {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(arr2)

	// dp maps "last value in arr1 after operations" → min operations
	dp := map[int]int{-1: 0}

	for _, x := range arr1 {
  // Membuat map (HashMap) — pencarian O(1)
		ndp := make(map[int]int)

		for last, ops := range dp {
			// Option 1: keep x if it's > last
			if x > last {
				if v, ok := ndp[x]; !ok || ops < v {
					ndp[x] = ops
				}
			}

			// Option 2: replace x with smallest arr2 element > last
			idx := sort.Search(len(arr2), func(i int) bool { return arr2[i] > last })
			if idx < len(arr2) {
				if v, ok := ndp[arr2[idx]]; !ok || ops+1 < v {
					ndp[arr2[idx]] = ops + 1
				}
			}
		}

		dp = ndp
	}

	ans := math.MaxInt32
	for _, ops := range dp {
		if ops < ans {
			ans = ops
		}
	}
	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}

func main() {
	// Test case 1
	fmt.Println(makeArrayIncreasing([]int{1, 5, 3, 6, 7}, []int{1, 3, 2, 4})) // 1

	// Test case 2
	fmt.Println(makeArrayIncreasing([]int{1, 2, 3}, []int{4, 5, 6})) // 0 (already strictly increasing)

	// Test case 3: impossible
	fmt.Println(makeArrayIncreasing([]int{1, 2, 1}, []int{1, 2})) // -1

	// Test case 4: all need replacement
	fmt.Println(makeArrayIncreasing([]int{10, 20, 30}, []int{1, 2, 3, 4})) // 3
}
```
