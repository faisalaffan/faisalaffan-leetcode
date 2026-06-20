# 0870 — Advantage Shuffle

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func AdvantageShuffle(nums1 []int, nums2 []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #870: Advantage Shuffle
// https://leetcode.com/problems/advantage-shuffle/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(AdvantageShuffle([]int{2, 7, 11, 15}, []int{1, 10, 4, 11}))
	fmt.Println(AdvantageShuffle([]int{12, 24, 8, 32}, []int{13, 25, 32, 11}))
	fmt.Println(AdvantageShuffle([]int{2, 0, 4, 1, 2}, []int{1, 3, 0, 0, 2}))
}

// Time: O(n log n) | Space: O(n)
func AdvantageShuffle(nums1 []int, nums2 []int) []int {
	n := len(nums1)
  // Urutkan secara ascending — O(n log n)
	sort.Ints(nums1)

  // Alokasi slice integer
	idx := make([]int, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range idx {
		idx[i] = i
	}
  // Custom sort dengan comparator
	sort.Slice(idx, func(i, j int) bool {
		return nums2[idx[i]] < nums2[idx[j]]
	})

  // Alokasi slice integer
	ans := make([]int, n)
	left, right := 0, n-1
	for _, x := range nums1 {
		if x > nums2[idx[left]] {
			ans[idx[left]] = x
			left++
		} else {
			ans[idx[right]] = x
			right--
		}
	}

	return ans
}
```
