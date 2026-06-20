# 0561 — Array Partition

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func ArrayPartition(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #561: Array Partition
// https://leetcode.com/problems/array-partition/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

// Time: O(n log n), Space: O(1)
func ArrayPartition(nums []int) int {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(nums)
	sum := 0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(nums); i += 2 {
		sum += nums[i]
	}
	return sum
}

func main() {
	fmt.Println(ArrayPartition([]int{1, 4, 3, 2}))
	fmt.Println(ArrayPartition([]int{6, 2, 6, 5, 1, 2}))
}
```
