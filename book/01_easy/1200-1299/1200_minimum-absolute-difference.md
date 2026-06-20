# 1200 — Minimum Absolute Difference

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumAbsDifference(arr []int) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1200: Minimum Absolute Difference
// https://leetcode.com/problems/minimum-absolute-difference/
// Difficulty: Easy
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minimumAbsDifference([]int{4, 2, 1, 3}))       // [[1,2],[2,3],[3,4]]
	fmt.Println(minimumAbsDifference([]int{1, 3, 6, 10, 15}))  // [[1,3]]
}

// LeetCode submission: minimumAbsDifference
func minimumAbsDifference(arr []int) [][]int {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(arr)
	minDiff := 1 << 31
	for i := 1; i < len(arr); i++ {
		diff := arr[i] - arr[i-1]
		if diff < minDiff {
			minDiff = diff
		}
	}
	var ans [][]int
	for i := 1; i < len(arr); i++ {
		if arr[i]-arr[i-1] == minDiff {
			ans = append(ans, []int{arr[i-1], arr[i]})
		}
	}
	return ans
}
```
