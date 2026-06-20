# 3427 — Sum Of Variable Length Subarrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func SumOfVariableLengthSubarrays(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n^2). Space: O(1).  
**Kompleksitas Ruang:** O(1).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3427: Sum of Variable Length Subarrays
// https://leetcode.com/problems/sum-of-variable-length-subarrays/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(SumOfVariableLengthSubarrays([]int{2, 3, 1}))
	fmt.Println(SumOfVariableLengthSubarrays([]int{3, 1, 1, 2}))
}

// SumOfVariableLengthSubarrays computes sum of subarrays where each subarray starts at i-(nums[i]%something) and ends at i.
// Time: O(n^2). Space: O(1).
func SumOfVariableLengthSubarrays(nums []int) int {
	n := len(nums)
	total := 0
	for i := 0; i < n; i++ {
		start := i - nums[i]
		if start < 0 {
			start = 0
		}
		for j := start; j <= i; j++ {
			total += nums[j]
		}
	}
	return total
}
```
