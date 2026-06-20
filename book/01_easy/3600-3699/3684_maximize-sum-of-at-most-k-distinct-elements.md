# 3684 — Maximize Sum Of At Most K Distinct Elements

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MaximizeSumOfAtMostKDistinctElements(nums []int, k int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3684: Maximize Sum of At Most K Distinct Elements
// https://leetcode.com/problems/maximize-sum-of-at-most-k-distinct-elements/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MaximizeSumOfAtMostKDistinctElements([]int{84, 93, 100, 77, 90}, 3))
	fmt.Println(MaximizeSumOfAtMostKDistinctElements([]int{84, 93, 100, 77, 93}, 3))
	fmt.Println(MaximizeSumOfAtMostKDistinctElements([]int{1, 1, 1, 2, 2, 2}, 6))
}

// Time: O(n log n)
// Space: O(n)
func MaximizeSumOfAtMostKDistinctElements(nums []int, k int) []int {
  // Membuat map (HashMap) — pencarian O(1)
	seen := make(map[int]bool)
  // Alokasi slice integer
	unique := make([]int, 0)
	for _, v := range nums {
		if !seen[v] {
			seen[v] = true
			unique = append(unique, v)
		}
	}

  // Custom sort dengan comparator
	sort.Slice(unique, func(i, j int) bool {
		return unique[i] > unique[j]
	})

	size := k
	if size > len(unique) {
		size = len(unique)
	}
	return unique[:size]
}
```
