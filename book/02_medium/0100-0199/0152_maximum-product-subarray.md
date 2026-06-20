# 0152 — Maximum Product Subarray

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxProduct(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #152: Maximum Product Subarray
// https://leetcode.com/problems/maximum-product-subarray/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

func maxProduct(nums []int) int {
  // Edge case: input kosong — langsung return
	if len(nums) == 0 {
		return 0
	}

	maxProd, minProd, result := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			maxProd, minProd = minProd, maxProd
		}

		maxProd = max(nums[i], maxProd*nums[i])
		minProd = min(nums[i], minProd*nums[i])

		result = max(result, maxProd)
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxProduct([]int{2, 3, -2, 4}))
	fmt.Println(maxProduct([]int{-2, 0, -1}))
	fmt.Println(maxProduct([]int{-2, 3, -4}))
}
```
