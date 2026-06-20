# 3371 — Identify The Largest Outlier In An Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func getLargestOutlier(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n) Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3371: Identify the Largest Outlier in an Array
// https://leetcode.com/problems/identify-the-largest-outlier-in-an-array/
// Difficulty: Medium
// Time: O(n) Space: O(n)

import "fmt"

func main() {
	fmt.Println(getLargestOutlier([]int{2, 3, 5, 10}))   // 10
	fmt.Println(getLargestOutlier([]int{-2, -1, -3, -6, 4})) // 4
	fmt.Println(getLargestOutlier([]int{1, 1, 1, 1, 1, 5, 5}))  // 5
}

func getLargestOutlier(nums []int) int {
  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[int]int)
	total := 0
	for _, v := range nums {
		freq[v]++
		total += v
	}

	ans := -(1 << 60)
	for _, v := range nums {
		outlier := total - 2*v
		if outlier == v && freq[v] < 2 {
			continue
		}
		if _, ok := freq[outlier]; ok && outlier > ans {
			ans = outlier
		}
	}
	return ans
}
```
