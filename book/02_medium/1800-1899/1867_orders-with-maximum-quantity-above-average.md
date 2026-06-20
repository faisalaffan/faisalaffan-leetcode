# 1867 — Orders With Maximum Quantity Above Average

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func OrdersAboveAverage(orders [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1867: Orders With Maximum Quantity Above Average
// https://leetcode.com/problems/orders-with-maximum-quantity-above-average/
// Difficulty: Medium [Paid]

import "fmt"

func main() {
	// Sample: order_id, quantity
	orders := [][]int{{1, 10}, {2, 5}, {3, 8}, {4, 3}, {5, 12}}
	fmt.Println(OrdersAboveAverage(orders))
}

// Time: O(n), Space: O(1)
func OrdersAboveAverage(orders [][]int) int {
	if len(orders) == 0 {
		return 0
	}
	sum := 0
	for _, o := range orders {
		sum += o[1]
	}
	avg := float64(sum) / float64(len(orders))
	count := 0
	for _, o := range orders {
		if float64(o[1]) > avg {
			count++
		}
	}
	// Find max quantity among orders above average
	maxQty := 0
	for _, o := range orders {
		if float64(o[1]) > avg && o[1] > maxQty {
			maxQty = o[1]
		}
	}
	return maxQty
}
```
