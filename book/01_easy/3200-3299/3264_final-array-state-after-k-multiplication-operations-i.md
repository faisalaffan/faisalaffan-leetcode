# 3264 — Final Array State After K Multiplication Operations I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func FinalArrayStateAfterKMultiplicationOperationsI(nums []int, k int, multiplier int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(k * n). Space: O(1).  
**Kompleksitas Ruang:** O(1).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3264: Final Array State After K Multiplication Operations I
// https://leetcode.com/problems/final-array-state-after-k-multiplication-operations-i/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FinalArrayStateAfterKMultiplicationOperationsI([]int{2, 1, 3, 5, 6}, 5, 2))
	fmt.Println(FinalArrayStateAfterKMultiplicationOperationsI([]int{1, 2}, 3, 4))
}

// FinalArrayStateAfterKMultiplicationOperationsI finds the minimum element each time, multiplies it by multiplier, and repeats k times.
// Time: O(k * n). Space: O(1).
func FinalArrayStateAfterKMultiplicationOperationsI(nums []int, k int, multiplier int) []int {
	for t := 0; t < k; t++ {
		// Find index of minimum element
		minIdx := 0
		for i := 1; i < len(nums); i++ {
			if nums[i] < nums[minIdx] {
				minIdx = i
			}
		}
		nums[minIdx] *= multiplier
	}
	return nums
}
```
