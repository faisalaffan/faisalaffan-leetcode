# 0384 — Shuffle An Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func Constructor(nums []int) Solution
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n) per shuffle  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #384: Shuffle an Array
// https://leetcode.com/problems/shuffle-an-array/
// Difficulty: Medium
// Time: O(n) per shuffle | Space: O(n)

import (
	"fmt"
	"math/rand"
)

type Solution struct {
	original []int
}

func Constructor(nums []int) Solution {
  // Alokasi slice integer
	orig := make([]int, len(nums))
	copy(orig, nums)
	return Solution{original: orig}
}

func (s *Solution) Reset() []int {
  // Alokasi slice integer
	result := make([]int, len(s.original))
	copy(result, s.original)
	return result
}

func (s *Solution) Shuffle() []int {
  // Alokasi slice integer
	result := make([]int, len(s.original))
	copy(result, s.original)
	// Fisher-Yates
	for i := len(result) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		result[i], result[j] = result[j], result[i]
	}
	return result
}

func main() {
	sol := Constructor([]int{1, 2, 3})
	fmt.Println("Reset:", sol.Reset())
	fmt.Println("Shuffle:", sol.Shuffle())
	fmt.Println("Shuffle:", sol.Shuffle())
	fmt.Println("Reset:", sol.Reset())
}
```
