# 1685 — Sum Of Absolute Differences In A Sorted Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func getSumAbsoluteDifferences(nums []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** O(n), Space: O(1) (excluding output)  
**Kompleksitas Ruang:** O(1) (excluding output)

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1685: Sum of Absolute Differences in a Sorted Array
// https://leetcode.com/problems/sum-of-absolute-differences-in-a-sorted-array/
// Difficulty: Medium
// Time: O(n), Space: O(1) (excluding output)

import "fmt"

func getSumAbsoluteDifferences(nums []int) []int {
	n := len(nums)
	total := 0
	for _, v := range nums {
		total += v
	}

  // Alokasi slice integer
	result := make([]int, n)
	prefix := 0
	for i, v := range nums {
		// Left side: v * i - prefix_sum_left
		// Right side: (total - prefix_sum_left - v) - v * (n-1-i)
		left := v*i - prefix
		right := (total - prefix - v) - v*(n-1-i)
		result[i] = left + right
		prefix += v
	}
	return result
}

func main() {
	fmt.Println(getSumAbsoluteDifferences([]int{2, 3, 5}))    // Expected: [4, 3, 5]
	fmt.Println(getSumAbsoluteDifferences([]int{1, 4, 6, 8, 10})) // Expected: [24, 15, 13, 15, 21]
	fmt.Println(getSumAbsoluteDifferences([]int{1, 2}))       // Expected: [1, 1]
}
```
