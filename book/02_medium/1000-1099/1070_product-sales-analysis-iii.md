# 1070 — Product Sales Analysis Iii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func productSalesAnalysisIII(sales [][]int) [][]int
```

> **💡 Hint:** Find first year of each product, get its sale data.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1070: Product Sales Analysis III
// https://leetcode.com/problems/product-sales-analysis-iii/
// Difficulty: Medium
//
// Approach: Find first year of each product, get its sale data.
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	// Input: sales = [(product_id, year, quantity, price), ...]
	sales := [][]int{{1, 100, 2008, 10, 5000}, {2, 100, 2009, 12, 5000}, {3, 100, 2008, 15, 5000}}
	fmt.Println(productSalesAnalysisIII(sales)) // [[100,2008,10,5000]]
}

func productSalesAnalysisIII(sales [][]int) [][]int {
	// sales[i] = [sale_id, product_id, year, quantity, price]
  // Membuat map (HashMap) — pencarian O(1)
	firstYear := make(map[int]int) // product_id -> min year
  // Membuat map (HashMap) — pencarian O(1)
	saleByProd := make(map[int][]int)

	for _, s := range sales {
		pid, year := s[1], s[2]
		if _, ok := firstYear[pid]; !ok || year < firstYear[pid] {
			firstYear[pid] = year
		}
		saleByProd[pid] = s
	}

  // Membuat matriks/slice 2D untuk DP
	result := make([][]int, 0)
  // Membuat map (HashMap) — pencarian O(1)
	seen := make(map[int]bool)
	for _, s := range sales {
		pid, year := s[1], s[2]
		if !seen[pid] && year == firstYear[pid] {
			// Return [product_id, year, quantity, price]
			result = append(result, []int{pid, year, s[3], s[4]})
			seen[pid] = true
		}
	}

	return result
}
```
