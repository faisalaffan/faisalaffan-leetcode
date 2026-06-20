# 1833 — Maximum Ice Cream Bars

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxIceCream(costs []int, coins int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1833: Maximum Ice Cream Bars
// https://leetcode.com/problems/maximum-ice-cream-bars/
// Difficulty: Medium
// Time: O(n log n), Space: O(1)

import (
	"fmt"
	"sort"
)

func maxIceCream(costs []int, coins int) int {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(costs)
	count := 0
	for _, c := range costs {
		if coins >= c {
			coins -= c
			count++
		} else {
			break
		}
	}
	return count
}

func main() {
	fmt.Println(maxIceCream([]int{1, 3, 2, 4, 1}, 7)) // Expected: 4
	fmt.Println(maxIceCream([]int{10, 6, 8, 7, 7, 8}, 5)) // Expected: 0
	fmt.Println(maxIceCream([]int{1, 6, 3, 1, 2, 5}, 20)) // Expected: 6
}
```
