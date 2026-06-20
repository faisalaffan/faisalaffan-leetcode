# 1827 — Minimum Operations To Make The Array Increasing

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MinOperations(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1827: Minimum Operations to Make the Array Increasing
// https://leetcode.com/problems/minimum-operations-to-make-the-array-increasing/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func MinOperations(nums []int) int {
	ops := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			needed := nums[i-1] - nums[i] + 1
			nums[i] += needed
			ops += needed
		}
	}
	return ops
}

func main() {
	fmt.Println(MinOperations([]int{1, 1, 1}))
	fmt.Println(MinOperations([]int{1, 5, 2, 4, 1}))
	fmt.Println(MinOperations([]int{8}))
}
```
