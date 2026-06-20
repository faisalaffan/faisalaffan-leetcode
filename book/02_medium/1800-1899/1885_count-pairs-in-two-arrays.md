# 1885 — Count Pairs In Two Arrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func CountPairs(nums1 []int, nums2 []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(n log n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1885: Count Pairs in Two Arrays
// https://leetcode.com/problems/count-pairs-in-two-arrays/
// Difficulty: Medium [Paid]

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(CountPairs([]int{1, 3, 4}, []int{1, 3, 4}))
	fmt.Println(CountPairs([]int{1, 2, 3}, []int{2, 3, 4}))
	fmt.Println(CountPairs([]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}))
}

// Time: O(n log n), Space: O(n)
func CountPairs(nums1 []int, nums2 []int) int {
	n := len(nums1)
  // Alokasi slice integer
	diff := make([]int, n)
	for i := 0; i < n; i++ {
		diff[i] = nums1[i] - nums2[i]
	}
  // Urutkan secara ascending — O(n log n)
	sort.Ints(diff)

	left, right := 0, n-1
	count := 0
  // Two-pointer: gerakkan kiri atau kanan
	for left < right {
		if diff[left]+diff[right] > 0 {
			count += right - left
			right--
		} else {
			left++
		}
	}
	return count
}
```
