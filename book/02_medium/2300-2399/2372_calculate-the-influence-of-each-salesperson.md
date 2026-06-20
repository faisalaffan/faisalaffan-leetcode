# 2372 — Calculate The Influence Of Each Salesperson

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan tabel database dan diminta untuk menulis query SQL. Karena repo ini menggunakan Go, query SQL disimulasikan dengan struktur data Go (map untuk grouping, slice untuk sorting, struct untuk representasi row).

Soal tipe ini menguji kemampuanmu menganalisis data relasional — seperti yang kamu lakukan dengan SQL di pekerjaan backend sehari-hari.

**Konsep kunci:** GROUP BY, JOIN, aggregate (SUM, COUNT, AVG), window function (RANK, ROW_NUMBER), HAVING.

**Fungsi yang perlu kamu implementasikan:**
```go
func calculateInfluence(sales [][]int, priceMap map[int]int) map[int]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n + m)  
**Kompleksitas Ruang:** O(n + m)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2372: Calculate the Influence of Each Salesperson
// https://leetcode.com/problems/calculate-the-influence-of-each-salesperson/
// Difficulty: Medium
// Time: O(n + m) | Space: O(n + m)
// Simulate: track total price by salesperson from sales and product tables.

import "fmt"

func main() {
	sales := [][]int{
		{1, 1, 10}, // salesperson_id, product_id, quantity
		{2, 1, 5},
		{1, 2, 8},
		{3, 2, 3},
		{2, 2, 2},
	}
	products := map[int]int{
		1: 100, // product_id -> price
		2: 50,
	}
	fmt.Println(calculateInfluence(sales, products)) // {1: 1400, 2: 600, 3: 150}

	sales2 := [][]int{
		{1, 1, 1},
		{2, 1, 1},
	}
	products2 := map[int]int{1: 500}
	fmt.Println(calculateInfluence(sales2, products2)) // {1: 500, 2: 500}
}

func calculateInfluence(sales [][]int, priceMap map[int]int) map[int]int {
  // Membuat map (HashMap) — pencarian O(1)
	result := make(map[int]int)
	for _, s := range sales {
		sp, prodID, qty := s[0], s[1], s[2]
		result[sp] += qty * priceMap[prodID]
	}
	return result
}
```
