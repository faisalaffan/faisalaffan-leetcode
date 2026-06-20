# 1122 — Relative Sort Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func relativeSortArray(arr1, arr2 []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1122: Relative Sort Array
// https://leetcode.com/problems/relative-sort-array/
// Difficulty: Easy
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(relativeSortArray([]int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}, []int{2, 1, 4, 3, 9, 6}))
	// [2,2,2,1,4,3,3,9,6,7,19]
	fmt.Println(relativeSortArray([]int{28, 6, 22, 8, 44, 17}, []int{22, 28, 8, 6}))
	// [22,28,8,6,17,44]
}

// LeetCode submission: relativeSortArray
func relativeSortArray(arr1, arr2 []int) []int {
  // Membuat map (HashMap) — pencarian O(1)
	rank := make(map[int]int, len(arr2))
	for i, v := range arr2 {
		rank[v] = i
	}
  // Custom sort dengan comparator
	sort.Slice(arr1, func(i, j int) bool {
		ri, okI := rank[arr1[i]]
		rj, okJ := rank[arr1[j]]
		if okI && okJ {
			return ri < rj
		}
		if okI {
			return true
		}
		if okJ {
			return false
		}
		return arr1[i] < arr1[j]
	})
	return arr1
}
```
