# 2032 — Two Out Of Three

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func TwoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2032: Two Out of Three
// https://leetcode.com/problems/two-out-of-three/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(TwoOutOfThree([]int{1, 1, 3, 2}, []int{2, 3}, []int{3}))    // [3 2]
	fmt.Println(TwoOutOfThree([]int{3, 1}, []int{2, 3}, []int{1, 2}))       // [2 3 1]
	fmt.Println(TwoOutOfThree([]int{1, 2, 2}, []int{4, 3, 3}, []int{5}))    // []
}

// Time: O(n), Space: O(n)
func TwoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
  // Membuat map (HashMap) — pencarian O(1)
	set1 := make(map[int]bool)
  // Membuat map (HashMap) — pencarian O(1)
	set2 := make(map[int]bool)
  // Membuat map (HashMap) — pencarian O(1)
	set3 := make(map[int]bool)

	for _, v := range nums1 {
		set1[v] = true
	}
	for _, v := range nums2 {
		set2[v] = true
	}
	for _, v := range nums3 {
		set3[v] = true
	}

  // Membuat map (HashMap) — pencarian O(1)
	count := make(map[int]int)
	for v := range set1 {
		count[v]++
	}
	for v := range set2 {
		count[v]++
	}
	for v := range set3 {
		count[v]++
	}

	var result []int
	for v, c := range count {
		if c >= 2 {
			result = append(result, v)
		}
	}
  // Urutkan secara ascending — O(n log n)
	sort.Ints(result)
	return result
}
```
