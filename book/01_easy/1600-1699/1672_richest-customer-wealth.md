# 1672 — Richest Customer Wealth

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MaximumWealth(accounts [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(m*n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1672: Richest Customer Wealth
// https://leetcode.com/problems/richest-customer-wealth/
// Difficulty: Easy

import "fmt"

// Time: O(m*n), Space: O(1)
func MaximumWealth(accounts [][]int) int {
	maxWealth := 0
	for _, customer := range accounts {
		sum := 0
		for _, amount := range customer {
			sum += amount
		}
		if sum > maxWealth {
			maxWealth = sum
		}
	}
	return maxWealth
}

func main() {
	fmt.Println(MaximumWealth([][]int{{1, 2, 3}, {3, 2, 1}}))
	fmt.Println(MaximumWealth([][]int{{1, 5}, {7, 3}, {3, 5}}))
	fmt.Println(MaximumWealth([][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}}))
}
```
