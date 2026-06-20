# 1968 — Array With Elements Not Equal To Average Of Neighbors

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func RearrangeArray(nums []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(n log n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1968: Array With Elements Not Equal to Average of Neighbors
// https://leetcode.com/problems/array-with-elements-not-equal-to-average-of-neighbors/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(RearrangeArray([]int{1, 2, 3, 4, 5}))
	fmt.Println(RearrangeArray([]int{6, 2, 0, 9, 7}))
}

// Time: O(n log n), Space: O(n)
func RearrangeArray(nums []int) []int {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(nums)
	n := len(nums)
  // Alokasi slice integer
	result := make([]int, n)
	left, right := 0, n-1
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			result[i] = nums[left]
			left++
		} else {
			result[i] = nums[right]
			right--
		}
	}
	return result
}
```
