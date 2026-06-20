# 1045 — Customers Who Bought All Products

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func customersWhoBoughtAllProducts(customer, product []int, totalProducts int) []int
```

> **💡 Hint:** Count distinct products per customer, compare to total products.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n) where n = len(customer_product)  
**Kompleksitas Ruang:** O(m) where m = number of customers

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1045: Customers Who Bought All Products
// https://leetcode.com/problems/customers-who-bought-all-products/
// Difficulty: Medium
//
// Approach: Count distinct products per customer, compare to total products.
// Time: O(n) where n = len(customer_product)
// Space: O(m) where m = number of customers

import "fmt"

func main() {
	// Simulated: customer_id, product_key
	customer := []int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
	product := []int{1, 1, 1, 1, 1, 2, 2, 2, 2, 2}
	fmt.Println(customersWhoBoughtAllProducts(customer, product, 2)) // [1,2,3,4,5]

	customer2 := []int{1, 1, 2}
	product2 := []int{1, 2, 1}
	fmt.Println(customersWhoBoughtAllProducts(customer2, product2, 2)) // [1]
}

func customersWhoBoughtAllProducts(customer, product []int, totalProducts int) []int {
  // Membuat map (HashMap) — pencarian O(1)
	bought := make(map[int]map[int]bool)
  // Membuat map (HashMap) — pencarian O(1)
	customerSet := make(map[int]bool)

  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(customer); i++ {
		c, p := customer[i], product[i]
		customerSet[c] = true
		if bought[c] == nil {
			bought[c] = make(map[int]bool)
		}
		bought[c][p] = true
	}

  // Alokasi slice integer
	result := make([]int, 0)
	for c := range customerSet {
		if len(bought[c]) == totalProducts {
			result = append(result, c)
		}
	}

	return result
}
```
