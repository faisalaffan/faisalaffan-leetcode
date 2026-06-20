# 0826 — Most Profit Assigning Work

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MostProfitAssigningWork(difficulty []int, profit []int, worker []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n + m log m)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #826: Most Profit Assigning Work
// https://leetcode.com/problems/most-profit-assigning-work/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MostProfitAssigningWork([]int{2, 4, 6, 8, 10}, []int{10, 20, 30, 40, 50}, []int{4, 5, 6, 7}))
	fmt.Println(MostProfitAssigningWork([]int{85, 47, 57}, []int{24, 66, 99}, []int{40, 25, 25}))
	fmt.Println(MostProfitAssigningWork([]int{68, 35, 52, 47, 86}, []int{67, 17, 1, 81, 3}, []int{92, 10, 85, 84, 82}))
}

// Time: O(n log n + m log m) | Space: O(n)
func MostProfitAssigningWork(difficulty []int, profit []int, worker []int) int {
	n := len(difficulty)
  // Alokasi slice integer
	jobs := make([][2]int, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range difficulty {
		jobs[i] = [2]int{difficulty[i], profit[i]}
	}

  // Custom sort dengan comparator
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i][0] < jobs[j][0]
	})
  // Urutkan secara ascending — O(n log n)
	sort.Ints(worker)

	ans := 0
	idx := 0
	maxProfit := 0

	for _, w := range worker {
		for idx < n && jobs[idx][0] <= w {
			if jobs[idx][1] > maxProfit {
				maxProfit = jobs[idx][1]
			}
			idx++
		}
		ans += maxProfit
	}

	return ans
}
```
