# 0713 — Subarray Product Less Than K

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func numSubarrayProductLessThanK(nums []int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #713: Subarray Product Less Than K
// https://leetcode.com/problems/subarray-product-less-than-k/
// Difficulty: Medium
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(numSubarrayProductLessThanK([]int{10, 5, 2, 6}, 100))
	fmt.Println(numSubarrayProductLessThanK([]int{1, 2, 3}, 0))
}

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}

	count := 0
	product := 1
	left := 0

	for right := 0; right < len(nums); right++ {
		product *= nums[right]

		for product >= k {
			product /= nums[left]
			left++
		}

		count += right - left + 1
	}

	return count
}
```
