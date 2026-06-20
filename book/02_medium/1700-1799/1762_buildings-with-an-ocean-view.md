# 1762 — Buildings With An Ocean View

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func findBuildings(heights []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1) excluding output  
**Kompleksitas Ruang:** O(1) excluding output

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1762: Buildings With an Ocean View
// https://leetcode.com/problems/buildings-with-an-ocean-view/
// Difficulty: Medium [Paid]
// Time: O(n), Space: O(1) excluding output

import "fmt"

func findBuildings(heights []int) []int {
  // Alokasi slice integer
	result := make([]int, 0)
	maxHeight := -1

	for i := len(heights) - 1; i >= 0; i-- {
		if heights[i] > maxHeight {
			result = append([]int{i}, result...)
			maxHeight = heights[i]
		}
	}
	return result
}

func main() {
	fmt.Println(findBuildings([]int{4, 2, 3, 1}))    // Expected: [0, 2, 3]
	fmt.Println(findBuildings([]int{4, 3, 2, 1}))    // Expected: [0, 1, 2, 3]
	fmt.Println(findBuildings([]int{1, 3, 2, 4}))    // Expected: [3]
}
```
