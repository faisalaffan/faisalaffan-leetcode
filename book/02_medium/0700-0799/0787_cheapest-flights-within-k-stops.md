# 0787 — Cheapest Flights Within K Stops

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(K * E)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #787: Cheapest Flights Within K Stops
// https://leetcode.com/problems/cheapest-flights-within-k-stops/
// Difficulty: Medium
// Time: O(K * E)
// Space: O(n)

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findCheapestPrice(3, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, 0, 2, 1))
	fmt.Println(findCheapestPrice(3, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, 0, 2, 0))
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
  // Alokasi slice integer
	prices := make([]int, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range prices {
		prices[i] = math.MaxInt32
	}
	prices[src] = 0

	for i := 0; i <= k; i++ {
  // Alokasi slice integer
		temp := make([]int, n)
		copy(temp, prices)
		for _, f := range flights {
			from, to, price := f[0], f[1], f[2]
			if prices[from] != math.MaxInt32 && prices[from]+price < temp[to] {
				temp[to] = prices[from] + price
			}
		}
		prices = temp
	}

	if prices[dst] == math.MaxInt32 {
		return -1
	}
	return prices[dst]
}
```
