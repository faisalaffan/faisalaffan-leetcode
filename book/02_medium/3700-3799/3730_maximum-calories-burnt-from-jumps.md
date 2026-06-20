# 3730 — Maximum Calories Burnt From Jumps

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maximumCaloriesBurntFromJumps(heights []int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3730: Maximum Calories Burnt from Jumps
// https://leetcode.com/problems/maximum-calories-burnt-from-jumps/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func maximumCaloriesBurntFromJumps(heights []int) int64 {
	n := len(heights)
  // Urutkan secara ascending — O(n log n)
	sort.Ints(heights)

  // Alokasi slice integer
	seq := make([]int, 0, n)
	l, r := 0, n-1
	for l <= r {
		seq = append(seq, heights[r])
		r--
		if l <= r {
			seq = append(seq, heights[l])
			l++
		}
	}

	total := int64(seq[0]) * int64(seq[0])
	for i := 1; i < n; i++ {
		diff := seq[i] - seq[i-1]
		total += int64(diff) * int64(diff)
	}
	return total
}

func main() {
	fmt.Println(maximumCaloriesBurntFromJumps([]int{1, 7, 9}))
	fmt.Println(maximumCaloriesBurntFromJumps([]int{5, 2, 4}))
	fmt.Println(maximumCaloriesBurntFromJumps([]int{3, 3}))
}
```
