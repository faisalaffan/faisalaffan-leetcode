# 1713 — Minimum Operations To Make A Subsequence

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minOperations(target []int, arr []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1713: Minimum Operations to Make a Subsequence
// https://leetcode.com/problems/minimum-operations-to-make-a-subsequence/
// Difficulty: Hard
// Strategy: Since target has distinct elements, map target values to indices.
// Then find LIS of those indices in arr.

import (
	"fmt"
	"sort"
)

func minOperations(target []int, arr []int) int {
	// Map target values to their indices
  // Membuat map (HashMap) — pencarian O(1)
	pos := make(map[int]int)
	for i, v := range target {
		pos[v] = i
	}

	// Build list of indices in arr that also appear in target
	// This becomes the LIS problem
  // Alokasi slice integer
	indices := make([]int, 0)
	for _, v := range arr {
		if idx, ok := pos[v]; ok {
			indices = append(indices, idx)
		}
	}

	// LIS on indices using patience sorting O(n log n)
  // Alokasi slice integer
	tails := make([]int, 0)
	for _, idx := range indices {
		// Find first element >= idx in tails
		j := sort.SearchInts(tails, idx)
		if j == len(tails) {
			tails = append(tails, idx)
		} else {
			tails[j] = idx
		}
	}

	// Minimum operations = len(target) - longest common subsequence length
	return len(target) - len(tails)
}

func main() {
	// Example 1: target=[5,1,3], arr=[9,4,2,3,4] -> 2
	target1 := []int{5, 1, 3}
	arr1 := []int{9, 4, 2, 3, 4}
	fmt.Printf("minOperations(%v, %v) = %d (expected 2)\n", target1, arr1, minOperations(target1, arr1))

	// Example 2: target=[6,4,8,1,3,2], arr=[4,7,6,2,3,8,6,1] -> 3
	target2 := []int{6, 4, 8, 1, 3, 2}
	arr2 := []int{4, 7, 6, 2, 3, 8, 6, 1}
	fmt.Printf("minOperations(%v, %v) = %d (expected 3)\n", target2, arr2, minOperations(target2, arr2))

	// Example 3: target=[1,2,3], arr=[1,2,3] -> 0
	target3 := []int{1, 2, 3}
	arr3 := []int{1, 2, 3}
	fmt.Printf("minOperations(%v, %v) = %d (expected 0)\n", target3, arr3, minOperations(target3, arr3))

	// Example 4: target=[1,2,3], arr=[4,5,6] -> 3
	target4 := []int{1, 2, 3}
	arr4 := []int{4, 5, 6}
	fmt.Printf("minOperations(%v, %v) = %d (expected 3)\n", target4, arr4, minOperations(target4, arr4))
}
```
