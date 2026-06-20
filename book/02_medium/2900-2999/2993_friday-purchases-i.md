# 2993 — Friday Purchases I

## Deskripsi

**Soal:** [2993. Friday Purchases I](https://leetcode.com/problems/friday-purchases-i/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(d) where d = distinct weeks

**Algoritma:** —

**Fungsi Solusi:** `func findFridayPurchasesI(purchases []Purchase) []WeeklyResult`

## Solusi Go

```go
package main

// LeetCode #2993: Friday Purchases I
// https://leetcode.com/problems/friday-purchases-i/
// Difficulty: Medium (SQL problem — simulated in Go)
//
// Simulates: For Fridays in November 2023, calculate total spending per
// week of the month. Week number = CEIL(day_of_month / 7). Only include
// weeks with at least one purchase.

import (
	"fmt"
	"sort"
	"time"
)

// Purchase represents the Purchases database table.
type Purchase struct {
	UserID       int
	PurchaseDate string // format: "YYYY-MM-DD"
	AmountSpend  int
}

// WeeklyResult holds the output.
type WeeklyResult struct {
	WeekOfMonth int
	TotalAmount int
}

// findFridayPurchasesI simulates the SQL query.
// Time: O(n) | Space: O(d) where d = distinct weeks
// n = number of purchases.
func findFridayPurchasesI(purchases []Purchase) []WeeklyResult {
  // Membuat map untuk pencarian O(1): key → value
	weeklyAmount := make(map[int]int)

	for _, p := range purchases {
		// Parse the date.
		t, err := time.Parse("2006-01-02", p.PurchaseDate)
		if err != nil {
			continue
		}

		// Filter: November 2023 only.
		if t.Year() != 2023 || t.Month() != time.November {
			continue
		}

		// Filter: Friday only (weekday = 5 = time.Friday).
		if t.Weekday() != time.Friday {
			continue
		}

		// Compute week_of_month: CEIL(day / 7).
		day := t.Day()
		weekOfMonth := (day-1)/7 + 1

		weeklyAmount[weekOfMonth] += p.AmountSpend
	}

	var results []WeeklyResult
	for w, amt := range weeklyAmount {
		results = append(results, WeeklyResult{WeekOfMonth: w, TotalAmount: amt})
	}

	// Order by week_of_month ASC.
	sort.Slice(results, func(i, j int) bool {
		return results[i].WeekOfMonth < results[j].WeekOfMonth
	})

	return results
}

func main() {
	// Test data: various purchases in November 2023.
	purchases := []Purchase{
		// Friday Nov 3 (week 1)
		{UserID: 1, PurchaseDate: "2023-11-03", AmountSpend: 50},
		{UserID: 2, PurchaseDate: "2023-11-03", AmountSpend: 30},
		// Non-Friday should be excluded
		{UserID: 3, PurchaseDate: "2023-11-04", AmountSpend: 20},
		// Friday Nov 10 (week 2)
		{UserID: 4, PurchaseDate: "2023-11-10", AmountSpend: 100},
		// Friday Nov 17 (week 3) — no purchases, should not appear
		// Friday Nov 24 (week 4)
		{UserID: 5, PurchaseDate: "2023-11-24", AmountSpend: 200},
		{UserID: 6, PurchaseDate: "2023-11-24", AmountSpend: 75},
		// Outside November
		{UserID: 7, PurchaseDate: "2023-10-06", AmountSpend: 40},
		{UserID: 8, PurchaseDate: "2023-12-01", AmountSpend: 60},
	}

	results := findFridayPurchasesI(purchases)

	fmt.Println("Friday Purchases I (week_of_month | total_amount):")
	for _, r := range results {
		fmt.Printf("%d | %d\n", r.WeekOfMonth, r.TotalAmount)
	}
	// Expected output:
	// 1 | 80
	// 2 | 100
	// 4 | 275
}
```
