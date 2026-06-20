# 1045 — Customers Who Bought All Products

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func customersWhoBoughtAllProducts(customer, product []int, totalProducts int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n) where n = len(customer_product)  |  **Ruang:** O(m) where m = number of customers

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

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
  // HashMap: O(1) lookup
	bought := make(map[int]map[int]bool)
  // HashMap: O(1) lookup
	customerSet := make(map[int]bool)

  // Linear scan O(n)
	for i := 0; i < len(customer); i++ {
		c, p := customer[i], product[i]
		customerSet[c] = true
		if bought[c] == nil {
			bought[c] = make(map[int]bool)
		}
		bought[c][p] = true
	}

  // Alokasi slice
	result := make([]int, 0)
	for c := range customerSet {
		if len(bought[c]) == totalProducts {
			result = append(result, c)
		}
	}

	return result
}
```
