# 1164 — Product Price At A Given Date

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func productPriceAtDate(products [][]int) []productPrice
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #1164: Product Price at a Given Date
// https://leetcode.com/problems/product-price-at-a-given-date/
// Difficulty: Medium

// For each product, find its price on 2019-08-16.
// If no change after date, price is 10 (default).
// Find the most recent price change on or before the date.

// Time: O(n log n)
// Space: O(n)

type productPrice struct {
	productID int
	price     int
}

func productPriceAtDate(products [][]int) []productPrice {
	// products[i] = [product_id, new_price, change_date]
	// Find price of each product on 2019-08-16

	// Group by product
  // Membuat map (HashMap) — pencarian O(1)
	changes := make(map[int][][2]int) // productID -> [(date, price)]
	for _, p := range products {
		id, price, date := p[0], p[1], p[2]
		changes[id] = append(changes[id], [2]int{date, price})
	}

	// Sort each product's changes by date
	for id := range changes {
  // Custom sort dengan comparator
		sort.Slice(changes[id], func(i, j int) bool {
			return changes[id][i][0] < changes[id][j][0]
		})
	}

	targetDate := 20190816
	result := make([]productPrice, 0, len(changes))

	for id, vals := range changes {
		price := 10 // default price
		for _, v := range vals {
			if v[0] <= targetDate {
				price = v[1]
			} else {
				break
			}
		}
		result = append(result, productPrice{id, price})
	}

  // Custom sort dengan comparator
	sort.Slice(result, func(i, j int) bool {
		return result[i].productID < result[j].productID
	})

	return result
}

func main() {
	products := [][]int{
		{1, 20, 20190801},
		{2, 50, 20190801},
		{1, 30, 20190815},
		{1, 40, 20190817},
		{2, 80, 20190814},
	}
	fmt.Printf("%v (expected: [{1 30} {2 80}])\n", productPriceAtDate(products))

	products2 := [][]int{
		{1, 20, 20190817},
		{2, 50, 20190817},
	}
	fmt.Printf("%v (expected: [{1 10} {2 10}])\n", productPriceAtDate(products2))
}
```
