# 3683 — Earliest Time To Finish One Task

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func EarliestTimeToFinishOneTask(tasks [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3683: Earliest Time to Finish One Task
// https://leetcode.com/problems/earliest-time-to-finish-one-task/
// Difficulty: Easy

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(EarliestTimeToFinishOneTask([][]int{{1, 6}, {2, 3}}))
	fmt.Println(EarliestTimeToFinishOneTask([][]int{{100, 100}, {100, 100}, {100, 100}}))
}

// Time: O(n)
// Space: O(1)
func EarliestTimeToFinishOneTask(tasks [][]int) int {
	ans := math.MaxInt
	for _, t := range tasks {
		finish := t[0] + t[1]
		if finish < ans {
			ans = finish
		}
	}
	return ans
}
```
