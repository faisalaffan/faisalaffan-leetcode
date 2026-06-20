# 3462 — Maximum Sum With At Most K Elements

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxSum(grid [][]int, limits []int, k int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n*m*log(m) + total*log(k)) Space: O(total)  
**Kompleksitas Ruang:** O(total)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3462: Maximum Sum With at Most K Elements
// https://leetcode.com/problems/maximum-sum-with-at-most-k-elements/
// Difficulty: Medium
// Time: O(n*m*log(m) + total*log(k)) Space: O(total)

import (
	"fmt"
	"sort"
)

func maxSum(grid [][]int, limits []int, k int) int64 {
	var candidates []int
	for i, row := range grid {
  // Urutkan secara ascending — O(n log n)
		sort.Ints(row)
		lim := limits[i]
		m := len(row)
		for j := m - lim; j < m; j++ {
			candidates = append(candidates, row[j])
		}
	}
  // Custom sort dengan comparator
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] > candidates[j]
	})
	var sum int64
	for i := 0; i < k && i < len(candidates); i++ {
		sum += int64(candidates[i])
	}
	return sum
}

func main() {
	fmt.Println(maxSum([][]int{{5, 3, 7}, {8, 2, 6}}, []int{2, 2}, 3)) // 21
	fmt.Println(maxSum([][]int{{1, 2}, {3, 4}}, []int{1, 1}, 2)) // 7
}
```
