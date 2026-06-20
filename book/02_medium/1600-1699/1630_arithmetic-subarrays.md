# 1630 — Arithmetic Subarrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func CheckArithmeticSubarrays(nums []int, l []int, r []int) []bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(M * N log N), Space: O(N)  
**Kompleksitas Ruang:** O(N)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1630: Arithmetic Subarrays
// https://leetcode.com/problems/arithmetic-subarrays/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(CheckArithmeticSubarrays([]int{4, 6, 5, 9, 3, 7}, []int{0, 0, 2}, []int{2, 3, 5}))
	fmt.Println(CheckArithmeticSubarrays([]int{-12, -9, -3, -12, -6, 15, 20, -25, -20, -15, -10}, []int{0, 1, 6, 4, 8, 7}, []int{4, 4, 9, 7, 9, 10}))
}

func CheckArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	// Time: O(M * N log N), Space: O(N)
	result := make([]bool, len(l))

  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(l); i++ {
  // Alokasi slice integer
		sub := make([]int, r[i]-l[i]+1)
		copy(sub, nums[l[i]:r[i]+1])
		result[i] = isArithmetic(sub)
	}

	return result
}

func isArithmetic(arr []int) bool {
	if len(arr) <= 2 {
		return true
	}

  // Urutkan secara ascending — O(n log n)
	sort.Ints(arr)
	diff := arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != diff {
			return false
		}
	}
	return true
}
```
