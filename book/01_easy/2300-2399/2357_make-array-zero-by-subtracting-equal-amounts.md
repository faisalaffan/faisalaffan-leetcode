# 2357 — Make Array Zero By Subtracting Equal Amounts

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MakeArrayZeroBySubtractingEqualAmounts(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2357: Make Array Zero by Subtracting Equal Amounts
// https://leetcode.com/problems/make-array-zero-by-subtracting-equal-amounts/
// Difficulty: Easy
// Time O(n log n) | Space O(1)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MakeArrayZeroBySubtractingEqualAmounts([]int{1, 5, 0, 3, 5})) // 3
	fmt.Println(MakeArrayZeroBySubtractingEqualAmounts([]int{0}))              // 0
}

func MakeArrayZeroBySubtractingEqualAmounts(nums []int) int {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(nums)
	count := 0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			count++
			sub := nums[i]
			for j := i; j < len(nums); j++ {
				nums[j] -= sub
			}
		}
	}
	return count
}
```
