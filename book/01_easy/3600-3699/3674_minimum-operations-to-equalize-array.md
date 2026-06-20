# 3674 — Minimum Operations To Equalize Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumOperationsToEqualizeArray(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3674: Minimum Operations to Equalize Array
// https://leetcode.com/problems/minimum-operations-to-equalize-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MinimumOperationsToEqualizeArray([]int{1, 2, 3}))
	fmt.Println(MinimumOperationsToEqualizeArray([]int{5, 5, 5}))
}

// Time: O(n)
// Space: O(1)
func MinimumOperationsToEqualizeArray(nums []int) int {
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[0] {
			return 1
		}
	}
	return 0
}
```
