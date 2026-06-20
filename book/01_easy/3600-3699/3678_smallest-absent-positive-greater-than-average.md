# 3678 — Smallest Absent Positive Greater Than Average

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func SmallestAbsentPositiveGreaterThanAverage(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n + m) where m is the range of candidate values  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3678: Smallest Absent Positive Greater Than Average
// https://leetcode.com/problems/smallest-absent-positive-greater-than-average/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(SmallestAbsentPositiveGreaterThanAverage([]int{3, 5}))
	fmt.Println(SmallestAbsentPositiveGreaterThanAverage([]int{-1, 1, 2}))
	fmt.Println(SmallestAbsentPositiveGreaterThanAverage([]int{4, -1}))
}

// Time: O(n + m) where m is the range of candidate values
// Space: O(n)
func SmallestAbsentPositiveGreaterThanAverage(nums []int) int {
  // Membuat map (HashMap) — pencarian O(1)
	has := make(map[int]bool)
	sum := 0
	for _, x := range nums {
		has[x] = true
		sum += x
	}

	avg := float64(sum) / float64(len(nums))
	ans := 1
	if int(avg)+1 > ans {
		ans = int(avg) + 1
	}

	for has[ans] {
		ans++
	}
	return ans
}
```
