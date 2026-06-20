# 3075 — Maximize Happiness Of Selected Children

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maximumHappinessSum(happiness []int, k int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3075: Maximize Happiness of Selected Children
// https://leetcode.com/problems/maximize-happiness-of-selected-children/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximumHappinessSum([]int{1, 2, 3}, 2))
	fmt.Println(maximumHappinessSum([]int{1, 1, 1, 1}, 2))
	fmt.Println(maximumHappinessSum([]int{2, 3, 4, 5}, 1))
}

func maximumHappinessSum(happiness []int, k int) int64 {
  // Custom sort dengan comparator
	sort.Slice(happiness, func(i, j int) bool {
		return happiness[i] > happiness[j]
	})
	ans := int64(0)
	for i := 0; i < k; i++ {
		val := happiness[i] - i
		if val > 0 {
			ans += int64(val)
		}
	}
	return ans
}
```
