# 3074 — Apple Redistribution Into Boxes

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func AppleRedistributionIntoBoxes(apples []int, capacity []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3074: Apple Redistribution into Boxes
// https://leetcode.com/problems/apple-redistribution-into-boxes/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	// LeetCode name: minimumBoxes
	fmt.Println(AppleRedistributionIntoBoxes([]int{1, 3, 2}, []int{4, 3, 1, 5, 2})) // 2
	fmt.Println(AppleRedistributionIntoBoxes([]int{5, 5, 5}, []int{2, 4, 2, 7}))    // 4
}

// Time: O(n log n) | Space: O(1)
// LeetCode submission name: minimumBoxes
func AppleRedistributionIntoBoxes(apples []int, capacity []int) int {
	totalApples := 0
	for _, a := range apples {
		totalApples += a
	}
	sort.Sort(sort.Reverse(sort.IntSlice(capacity)))
	boxes := 0
	for _, c := range capacity {
		boxes++
		totalApples -= c
		if totalApples <= 0 {
			return boxes
		}
	}
	return boxes
}
```
