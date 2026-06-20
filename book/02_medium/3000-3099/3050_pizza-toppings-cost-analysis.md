# 3050 — Pizza Toppings Cost Analysis

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan tabel database dan diminta untuk menulis query SQL. Karena repo ini menggunakan Go, query SQL disimulasikan dengan struktur data Go (map untuk grouping, slice untuk sorting, struct untuk representasi row).

Soal tipe ini menguji kemampuanmu menganalisis data relasional — seperti yang kamu lakukan dengan SQL di pekerjaan backend sehari-hari.

**Konsep kunci:** GROUP BY, JOIN, aggregate (SUM, COUNT, AVG), window function (RANK, ROW_NUMBER), HAVING.

**Fungsi yang perlu kamu implementasikan:**
```go
func pizzaToppingsCostAnalysis(toppings []Topping) []PizzaCombo
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n^3)  
**Kompleksitas Ruang:** O(C(n,3))

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3050: Pizza Toppings Cost Analysis
// https://leetcode.com/problems/pizza-toppings-cost-analysis/
// Difficulty: Medium (SQL problem — implemented in Go)
// Time: O(n^3) | Space: O(C(n,3))

import (
	"fmt"
	"sort"
)

type Topping struct {
	Name string
	Cost float64
}

type PizzaCombo struct {
	Pizza     string
	TotalCost float64
}

func pizzaToppingsCostAnalysis(toppings []Topping) []PizzaCombo {
	n := len(toppings)
	var results []PizzaCombo

	// Generate all C(n,3) combinations
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				total := toppings[i].Cost + toppings[j].Cost + toppings[k].Cost
				pizzaName := toppings[i].Name + "," + toppings[j].Name + "," + toppings[k].Name
				results = append(results, PizzaCombo{
					Pizza:     pizzaName,
					TotalCost: total,
				})
			}
		}
	}

	// Sort by total_cost DESC, then pizza name ASC (which is topping1,topping2,topping3 ASC
	// since we generate in sorted order)
  // Custom sort dengan comparator
	sort.Slice(results, func(a, b int) bool {
		if results[a].TotalCost != results[b].TotalCost {
			return results[a].TotalCost > results[b].TotalCost // DESC
		}
		return results[a].Pizza < results[b].Pizza // ASC
	})

	return results
}

func main() {
	toppings := []Topping{
		{Name: "Pepperoni", Cost: 0.50},
		{Name: "Sausage", Cost: 0.70},
		{Name: "Chicken", Cost: 0.55},
		{Name: "Extra Cheese", Cost: 0.40},
		{Name: "Mushrooms", Cost: 0.60},
	}

	fmt.Println("Pizza Toppings Cost Analysis")
	fmt.Println("============================")
	fmt.Printf("%-40s %s\n", "Pizza", "Total Cost")
	fmt.Println("----------------------------------------")

	results := pizzaToppingsCostAnalysis(toppings)
	for _, r := range results {
		fmt.Printf("%-40s $%.2f\n", r.Pizza, r.TotalCost)
	}
}
```
