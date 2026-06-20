# 3502 — Minimum Cost To Reach Every Position

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumCostToReachEveryPosition(cost []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n). Space: O(n).  
**Kompleksitas Ruang:** O(n).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3502: Minimum Cost to Reach Every Position
// https://leetcode.com/problems/minimum-cost-to-reach-every-position/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MinimumCostToReachEveryPosition([]int{5, 3, 4, 1, 3, 2}))
	fmt.Println(MinimumCostToReachEveryPosition([]int{1, 2, 3, 2, 1}))
}

// MinimumCostToReachEveryPosition returns an array where ans[i] is the minimum cost to reach position i.
// You can travel from j to i (j < i) at cost cost[i], or from i to j at cost cost[j].
// Time: O(n). Space: O(n).
func MinimumCostToReachEveryPosition(cost []int) []int {
	n := len(cost)
  // Alokasi slice integer
	result := make([]int, n)
	result[0] = cost[0]
	for i := 1; i < n; i++ {
		if cost[i] < result[i-1] {
			result[i] = cost[i]
		} else {
			result[i] = result[i-1]
		}
	}
	return result
}
```
