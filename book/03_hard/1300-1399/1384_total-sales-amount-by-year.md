# 1384 — Total Sales Amount By Year

## Deskripsi

**Soal:** [1384. Total Sales Amount By Year](https://leetcode.com/problems/total-sales-amount-by-year/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

**Fungsi Solusi:** `func totalSalesByYear(sales []sale) map[int]map[int]int`

> **Ide Kunci:** Simulate SQL query in Go.

## Solusi Go

```go
package main

// LeetCode #1384: Total Sales Amount by Year
// https://leetcode.com/problems/total-sales-amount-by-year/
// Difficulty: Hard [Paid]
//
// Approach: Simulate SQL query in Go.
// Given products and sales (each with a period and average_daily_sales),
// compute total sales amount per year per product.
// A sale period may span multiple years; each year gets the proportional
// number of days multiplied by average_daily_sales.

import "fmt"

type sale struct {
	productID int
	startDate string // "YYYY-MM-DD"
	endDate   string
	avgDaily  int
}

type product struct {
	id   int
	name string
}

func totalSalesByYear(sales []sale) map[int]map[int]int {
	// result[year][productID] = total sales
  // Membuat map untuk pencarian O(1): key → value
	result := make(map[int]map[int]int)

	dim := []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	isLeap := func(y int) bool {
		return y%4 == 0 && (y%100 != 0 || y%400 == 0)
	}

	daysInYear := func(y int) int {
		if isLeap(y) {
			return 366
		}
		return 365
	}

	parseDate := func(s string) (y, m, d int) {
		fmt.Sscanf(s, "%d-%d-%d", &y, &m, &d)
		return
	}

	dayOfYear := func(y, m, d int) int {
		days := d
		for i := 1; i < m; i++ {
			days += dim[i]
		}
		if m > 2 && isLeap(y) {
			days++
		}
		return days
	}

	for _, s := range sales {
		y1, m1, d1 := parseDate(s.startDate)
		y2, m2, d2 := parseDate(s.endDate)

		if y1 == y2 {
			days := dayOfYear(y2, m2, d2) - dayOfYear(y1, m1, d1) + 1
			if result[y1] == nil {
				result[y1] = make(map[int]int)
			}
			result[y1][s.productID] += days * s.avgDaily
		} else {
			// First year: startDate to Dec 31
			daysY1 := daysInYear(y1) - dayOfYear(y1, m1, d1) + 1
			if result[y1] == nil {
				result[y1] = make(map[int]int)
			}
			result[y1][s.productID] += daysY1 * s.avgDaily

			// Middle full years
			for y := y1 + 1; y < y2; y++ {
				if result[y] == nil {
					result[y] = make(map[int]int)
				}
				result[y][s.productID] += daysInYear(y) * s.avgDaily
			}

			// Last year: Jan 1 to endDate
			daysY2 := dayOfYear(y2, m2, d2)
			if result[y2] == nil {
				result[y2] = make(map[int]int)
			}
			result[y2][s.productID] += daysY2 * s.avgDaily
		}
	}

	return result
}

func main() {
	sales := []sale{
		{1, "2020-01-01", "2020-12-31", 10},     // 366*10 = 3660 (leap)
		{2, "2020-06-01", "2021-06-01", 5},       // spans 2020-2021
		{3, "2019-11-01", "2020-02-29", 20},      // spans 2019-2020
	}

	result := totalSalesByYear(sales)
	for y := 2019; y <= 2021; y++ {
		if byYear, ok := result[y]; ok {
			for pid, total := range byYear {
				fmt.Printf("Year %d, Product %d: %d\n", y, pid, total)
			}
		}
	}
}
```
