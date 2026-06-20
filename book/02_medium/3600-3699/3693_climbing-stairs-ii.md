# 3693 — Climbing Stairs Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func climbingStairsIi(n int, costs []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3693: Climbing Stairs II
// https://leetcode.com/problems/climbing-stairs-ii/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func climbingStairsIi(n int, costs []int) int {
	dp0, dp1, dp2 := 0, 0, 0
	for j := 0; j < n; j++ {
		cur := dp2 + costs[j] + 1
		if j >= 1 {
			cand := dp1 + costs[j] + 4
			if cand < cur {
				cur = cand
			}
		}
		if j >= 2 {
			cand := dp0 + costs[j] + 9
			if cand < cur {
				cur = cand
			}
		}
		dp0, dp1, dp2 = dp1, dp2, cur
	}
	return dp2
}

func main() {
	fmt.Println(climbingStairsIi(3, []int{1, 2, 3}))
	fmt.Println(climbingStairsIi(2, []int{5, 10}))
	fmt.Println(climbingStairsIi(5, []int{1, 1, 1, 1, 1}))
}
```
