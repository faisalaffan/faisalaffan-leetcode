# 2994 — Friday Purchases Ii

## Deskripsi

**Soal:** [2994. Friday Purchases Ii](https://leetcode.com/problems/friday-purchases-ii/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

**Fungsi Solusi:** `func fridayPurchasesII(purchases []Purchase) []string`

## Solusi Go

```go
package main

// LeetCode #2994: Friday Purchases II (SQL simulation)
// https://leetcode.com/problems/friday-purchases-ii/
// Difficulty: Hard [Paid]
//
// For each Friday, calculate total amount and distinct user count.
// Only consider purchases made on Fridays.
// Sort by week starting date (the Friday date) ascending.

import (
	"fmt"
	"sort"
	"time"
)

type Purchase struct {
	UserID       int
	PurchaseDate string
	Amount       float64
}

func fridayPurchasesII(purchases []Purchase) []string {
	// Filter only Fridays
	var fridayPurchases []Purchase
	for _, p := range purchases {
		t, err := time.Parse("2006-01-02", p.PurchaseDate[:10])
		if err != nil {
			continue
		}
		if t.Weekday() == time.Friday {
			fridayPurchases = append(fridayPurchases, p)
		}
	}

	type weeklyData struct {
		total float64
		users map[int]bool
	}
  // Membuat map untuk pencarian O(1): key → value
	weekly := make(map[string]*weeklyData)
  // Membuat slice untuk menyimpan hasil
	weekOrder := make([]string, 0)

	for _, p := range fridayPurchases {
		wk := p.PurchaseDate[:10]
		if _, ok := weekly[wk]; !ok {
			weekly[wk] = &weeklyData{users: make(map[int]bool)}
			weekOrder = append(weekOrder, wk)
		}
		weekly[wk].total += p.Amount
		weekly[wk].users[p.UserID] = true
	}

	sort.Strings(weekOrder)

	var result []string
	for _, wk := range weekOrder {
		s := weekly[wk]
		result = append(result, fmt.Sprintf("%s|%.2f|%d", wk, s.total, len(s.users)))
	}
	return result
}

func main() {
	// Test 1: Basic Friday purchases
	purchases := []Purchase{
		{1, "2023-11-24", 100.50},
		{2, "2023-11-24", 50.25},
		{1, "2023-12-01", 200.00},
		{3, "2023-12-01", 75.00},
	}
	fmt.Println("Test 1:")
	for _, r := range fridayPurchasesII(purchases) {
		fmt.Println(r)
	}

	// Test 2: Mix of Friday and non-Friday
	fmt.Println("\nTest 2 (mix of days):")
	purchases2 := []Purchase{
		{1, "2023-11-22", 10.00}, // Wednesday - ignored
		{1, "2023-11-24", 100.00}, // Friday
		{2, "2023-11-23", 20.00}, // Thursday - ignored
		{3, "2023-11-24", 50.00}, // Friday
	}
	for _, r := range fridayPurchasesII(purchases2) {
		fmt.Println(r)
	}

	// Test 3: Same user multiple times same Friday
	fmt.Println("\nTest 3 (same user multiple times):")
	purchases3 := []Purchase{
		{1, "2023-11-24", 30.00},
		{1, "2023-11-24", 40.00},
		{2, "2023-11-24", 50.00},
	}
	for _, r := range fridayPurchasesII(purchases3) {
		fmt.Println(r)
	}

	// Test 4: No Friday purchases
	fmt.Println("\nTest 4 (no Fridays):")
	purchases4 := []Purchase{
		{1, "2023-11-22", 10.00},
	}
	fmt.Println(fridayPurchasesII(purchases4))
}
```
