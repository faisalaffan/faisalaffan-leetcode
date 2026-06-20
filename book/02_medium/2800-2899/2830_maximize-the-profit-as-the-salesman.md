# 2830 — Maximize The Profit As The Salesman

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MaximizeTheProfitAsTheSalesman(n int, offers [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** O(n + m)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2830: Maximize the Profit as the Salesman
// https://leetcode.com/problems/maximize-the-profit-as-the-salesman/
// Difficulty: Medium
// Time: O(n + m) | Space: O(n)

import "fmt"

func MaximizeTheProfitAsTheSalesman(n int, offers [][]int) int {
	// Group offers by end position
  // Membuat matriks/slice 2D untuk DP
	byEnd := make([][][]int, n)
	for _, offer := range offers {
		start, end, gold := offer[0], offer[1], offer[2]
		byEnd[end] = append(byEnd[end], []int{start, gold})
	}

  // Alokasi slice integer
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		if i > 0 {
			dp[i] = dp[i-1]
		}
		for _, offer := range byEnd[i] {
			start, gold := offer[0], offer[1]
			val := gold
			if start > 0 {
				val += dp[start-1]
			}
			if val > dp[i] {
				dp[i] = val
			}
		}
	}

	return dp[n-1]
}

func main() {
	fmt.Println(MaximizeTheProfitAsTheSalesman(5, [][]int{{0, 0, 1}, {0, 2, 2}, {1, 3, 2}}))
	fmt.Println(MaximizeTheProfitAsTheSalesman(3, [][]int{{0, 0, 5}, {1, 1, 3}, {2, 2, 4}}))
}
```
