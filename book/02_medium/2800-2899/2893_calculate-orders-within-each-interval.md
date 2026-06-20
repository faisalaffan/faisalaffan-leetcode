# 2893 — Calculate Orders Within Each Interval

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func CalculateOrdersWithinEachInterval(orders []Order, interval int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2893: Calculate Orders Within Each Interval
// https://leetcode.com/problems/calculate-orders-within-each-interval/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

type Order struct {
	Time   int
	Amount int
}

func CalculateOrdersWithinEachInterval(orders []Order, interval int) []int {
	if len(orders) == 0 {
		return []int{}
	}

	// Group orders by interval
	maxTime := orders[len(orders)-1].Time
	bucketCount := maxTime/interval + 1
  // Alokasi slice integer
	buckets := make([]int, bucketCount)

	for _, o := range orders {
		idx := o.Time / interval
		buckets[idx] += o.Amount
	}

	return buckets
}

func main() {
	orders := []Order{
		{0, 10}, {1, 20}, {4, 30}, {6, 40},
	}
	fmt.Println(CalculateOrdersWithinEachInterval(orders, 3))

	orders2 := []Order{
		{0, 5}, {2, 10},
	}
	fmt.Println(CalculateOrdersWithinEachInterval(orders2, 5))
}
```
