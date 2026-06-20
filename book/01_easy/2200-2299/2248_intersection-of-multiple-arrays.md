# 2248 — Intersection Of Multiple Arrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func IntersectionOfMultipleArrays(nums [][]int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n * m), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2248: Intersection of Multiple Arrays
// https://leetcode.com/problems/intersection-of-multiple-arrays/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(IntersectionOfMultipleArrays([][]int{{3, 1, 2, 4, 5}, {1, 2, 3, 4}, {3, 4, 5, 6}})) // [3 4]
	fmt.Println(IntersectionOfMultipleArrays([][]int{{1, 2, 3}, {4, 5, 6}}))                        // []
}

// Time: O(n * m), Space: O(n)
func IntersectionOfMultipleArrays(nums [][]int) []int {
  // Edge case: input kosong — langsung return
	if len(nums) == 0 {
		return []int{}
	}

  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[int]int)
	for _, v := range nums[0] {
		freq[v] = 1
	}

	for i := 1; i < len(nums); i++ {
  // Membuat map (HashMap) — pencarian O(1)
		seen := make(map[int]bool)
		for _, v := range nums[i] {
			if !seen[v] {
				freq[v]++
				seen[v] = true
			}
		}
	}

	var result []int
	for v, c := range freq {
		if c == len(nums) {
			result = append(result, v)
		}
	}
  // Urutkan secara ascending — O(n log n)
	sort.Ints(result)
	return result
}
```
