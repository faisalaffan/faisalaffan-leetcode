# 1174 — Immediate Food Delivery Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func immediateFoodDeliveryII(deliveries []delivery) float64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
	"math"
)

// LeetCode #1174: Immediate Food Delivery II
// https://leetcode.com/problems/immediate-food-delivery-ii/
// Difficulty: Medium

// For each customer's first order, find percentage that is "immediate"
// (order_date == customer_pref_delivery_date).

// Time: O(n)
// Space: O(n)

type delivery struct {
	ID         int
	customerID int
	orderDate  int
	prefDate   int
}

func immediateFoodDeliveryII(deliveries []delivery) float64 {
	// For each customer, track first order (earliest orderDate, tie-break by smallest ID)
	type firstOrder struct {
		date     int
		prefDate int
	}
  // Membuat map (HashMap) — pencarian O(1)
	custFirst := make(map[int]firstOrder)

	for _, d := range deliveries {
		fo, exists := custFirst[d.customerID]
		if !exists || d.orderDate < fo.date {
			custFirst[d.customerID] = firstOrder{d.orderDate, d.prefDate}
		}
	}

	immediate := 0
	total := 0
	for _, fo := range custFirst {
		total++
		if fo.date == fo.prefDate {
			immediate++
		}
	}

	if total == 0 {
		return 0
	}
	return math.Round(float64(immediate)/float64(total)*100) / 100
}

func main() {
	d1 := []delivery{
		{1, 1, 20190801, 20190801},
		{2, 2, 20190802, 20190803},
		{3, 1, 20190803, 20190804},
	}
	fmt.Printf("%.2f (expected: 0.50)\n", immediateFoodDeliveryII(d1))

	d2 := []delivery{
		{1, 1, 20190801, 20190801},
		{2, 1, 20190802, 20190802},
	}
	fmt.Printf("%.2f (expected: 1.00)\n", immediateFoodDeliveryII(d2))
}
```
