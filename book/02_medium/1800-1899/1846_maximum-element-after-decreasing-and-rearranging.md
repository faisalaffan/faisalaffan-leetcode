# 1846 — Maximum Element After Decreasing And Rearranging

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MaximumElementAfterDecreasingAndRearranging(arr []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n), Space: O(1) (ignoring sort space)  
**Kompleksitas Ruang:** O(1) (ignoring sort space)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1846: Maximum Element After Decreasing and Rearranging
// https://leetcode.com/problems/maximum-element-after-decreasing-and-rearranging/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MaximumElementAfterDecreasingAndRearranging([]int{2, 2, 1, 2, 1}))
	fmt.Println(MaximumElementAfterDecreasingAndRearranging([]int{100, 1, 1000}))
	fmt.Println(MaximumElementAfterDecreasingAndRearranging([]int{1, 2, 3, 4, 5}))
}

// Time: O(n log n), Space: O(1) (ignoring sort space)
func MaximumElementAfterDecreasingAndRearranging(arr []int) int {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(arr)
	arr[0] = 1
	for i := 1; i < len(arr); i++ {
		if arr[i]-arr[i-1] > 1 {
			arr[i] = arr[i-1] + 1
		}
	}
	return arr[len(arr)-1]
}
```
