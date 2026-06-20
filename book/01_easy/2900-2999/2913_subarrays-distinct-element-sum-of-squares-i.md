# 2913 — Subarrays Distinct Element Sum Of Squares I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func SubarraysDistinctElementSumOfSquaresI(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2913: Subarrays Distinct Element Sum of Squares I
// https://leetcode.com/problems/subarrays-distinct-element-sum-of-squares-i/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: sumCounts
	fmt.Println(SubarraysDistinctElementSumOfSquaresI([]int{1, 2, 1})) // 15
	fmt.Println(SubarraysDistinctElementSumOfSquaresI([]int{2, 2}))    // 3
}

// Time: O(n^2) | Space: O(n)
// LeetCode submission name: sumCounts
func SubarraysDistinctElementSumOfSquaresI(nums []int) int {
	n := len(nums)
	total := 0
	for i := 0; i < n; i++ {
  // Membuat map (HashMap) — pencarian O(1)
		seen := make(map[int]bool)
		distinct := 0
		for j := i; j < n; j++ {
			if !seen[nums[j]] {
				seen[nums[j]] = true
				distinct++
			}
			total += distinct * distinct
		}
	}
	return total
}
```
