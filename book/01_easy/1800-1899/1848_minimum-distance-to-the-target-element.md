# 1848 — Minimum Distance To The Target Element

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func GetMinDistance(nums []int, target int, start int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1848: Minimum Distance to the Target Element
// https://leetcode.com/problems/minimum-distance-to-the-target-element/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func GetMinDistance(nums []int, target int, start int) int {
	minDist := len(nums)
	for i, num := range nums {
		if num == target {
			dist := start - i
			if dist < 0 {
				dist = -dist
			}
			if dist < minDist {
				minDist = dist
			}
		}
	}
	return minDist
}

func main() {
	fmt.Println(GetMinDistance([]int{1, 2, 3, 4, 5}, 5, 3))
	fmt.Println(GetMinDistance([]int{1}, 1, 0))
	fmt.Println(GetMinDistance([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 1, 0))
}
```
