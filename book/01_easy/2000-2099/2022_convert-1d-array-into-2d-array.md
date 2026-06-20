# 2022 — Convert 1D Array Into 2D Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func ConvertOneDArrayIntoTwoDArray(original []int, m int, n int) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(m*n), Space: O(m*n)  
**Kompleksitas Ruang:** O(m*n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2022: Convert 1D Array Into 2D Array
// https://leetcode.com/problems/convert-1d-array-into-2d-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(ConvertOneDArrayIntoTwoDArray([]int{1, 2, 3, 4}, 2, 2)) // [[1 2] [3 4]]
	fmt.Println(ConvertOneDArrayIntoTwoDArray([]int{1, 2, 3}, 1, 3))    // [[1 2 3]]
	fmt.Println(ConvertOneDArrayIntoTwoDArray([]int{1, 2}, 1, 1))       // []
}

// Time: O(m*n), Space: O(m*n)
func ConvertOneDArrayIntoTwoDArray(original []int, m int, n int) [][]int {
	if len(original) != m*n {
		return [][]int{}
	}

  // Membuat matriks/slice 2D untuk DP
	result := make([][]int, m)
	for i := 0; i < m; i++ {
		result[i] = make([]int, n)
		for j := 0; j < n; j++ {
			result[i][j] = original[i*n+j]
		}
	}
	return result
}
```
