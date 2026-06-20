# 1418 — Display Table Of Food Orders In A Restaurant

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan aturan permainan dan harus menentukan siapa yang menang atau berapa skor maksimal. Tugasmu adalah menganalisis permainan dan membuat keputusan optimal di setiap langkah.

Soal game theory menguji kemampuanmu berpikir beberapa langkah ke depan (minimax). Seringkali diselesaikan dengan DP (Dynamic Programming) untuk menyimpan hasil subproblem.

**Konsep kunci:** minimax, optimal play, game state, DP memoization, win/lose positions.

**Fungsi yang perlu kamu implementasikan:**
```go
func displayTable(orders [][]string) [][]string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n log n) where n = number of orders  
**Kompleksitas Ruang:** O(n) for maps

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1418: Display Table of Food Orders in a Restaurant
// https://leetcode.com/problems/display-table-of-food-orders-in-a-restaurant/
// Difficulty: Medium

import "fmt"
import "sort"
import "strconv"

func main() {
	// Test case 1
	fmt.Println(displayTable([][]string{
		{"David", "3", "Ceviche"},
		{"Corina", "10", "Beef Burrito"},
		{"David", "3", "Fried Chicken"},
		{"Carla", "5", "Water"},
		{"Carla", "5", "Ceviche"},
		{"Rous", "3", "Ceviche"},
	}))

	// Test case 2
	fmt.Println(displayTable([][]string{
		{"James", "12", "Fried Chicken"},
		{"Ratesh", "12", "Fried Chicken"},
		{"Amadeus", "12", "Fried Chicken"},
		{"Adam", "1", "Canadian Waffles"},
		{"Brianna", "1", "Canadian Waffles"},
	}))
}

// Time: O(n log n) where n = number of orders
// Space: O(n) for maps
func displayTable(orders [][]string) [][]string {
  // Membuat map (HashMap) — pencarian O(1)
	foodItems := make(map[string]bool)
  // Membuat map (HashMap) — pencarian O(1)
	tableOrders := make(map[int]map[string]int)

	for _, o := range orders {
		table, _ := strconv.Atoi(o[1])
		food := o[2]

		foodItems[food] = true
		if tableOrders[table] == nil {
			tableOrders[table] = make(map[string]int)
		}
		tableOrders[table][food]++
	}

	// Sort food items (excluding "Table" header)
	foods := make([]string, 0, len(foodItems))
	for f := range foodItems {
		foods = append(foods, f)
	}
	sort.Strings(foods)

	// Sort table numbers
  // Alokasi slice integer
	tables := make([]int, 0, len(tableOrders))
	for t := range tableOrders {
		tables = append(tables, t)
	}
  // Urutkan secara ascending — O(n log n)
	sort.Ints(tables)

	// Build result
  // Membuat matriks/slice 2D untuk DP
	result := make([][]string, 0, len(tables)+1)
	header := make([]string, 0, len(foods)+1)
	header = append(header, "Table")
	header = append(header, foods...)
	result = append(result, header)

	for _, t := range tables {
		row := make([]string, 0, len(foods)+1)
		row = append(row, strconv.Itoa(t))
		for _, f := range foods {
			row = append(row, strconv.Itoa(tableOrders[t][f]))
		}
		result = append(result, row)
	}

	return result
}
```
