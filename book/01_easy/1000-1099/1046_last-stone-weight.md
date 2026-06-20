# 1046 — Last Stone Weight

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func lastStoneWeight(stones []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1046: Last Stone Weight
// https://leetcode.com/problems/last-stone-weight/
// Difficulty: Easy
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(lastStoneWeight([]int{2, 7, 4, 1, 8, 1})) // 1
	fmt.Println(lastStoneWeight([]int{1}))                 // 1
	fmt.Println(lastStoneWeight([]int{2, 2}))              // 0
}

// LeetCode submission: lastStoneWeight
func lastStoneWeight(stones []int) int {
	for len(stones) > 1 {
  // Urutkan secara ascending — O(n log n)
		sort.Ints(stones)
		n := len(stones)
		if stones[n-1] == stones[n-2] {
			stones = stones[:n-2]
		} else {
			stones[n-2] = stones[n-1] - stones[n-2]
			stones = stones[:n-1]
		}
	}
	if len(stones) == 1 {
		return stones[0]
	}
	return 0
}
```
