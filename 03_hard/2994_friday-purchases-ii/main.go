package main

// LeetCode #2994: Friday Purchases II (SQL simulation)
// https://leetcode.com/problems/friday-purchases-ii/
// Difficulty: Hard [Paid]

import (
	"fmt"
	"sort"
)

type Purchase struct {
	UserID       int
	PurchaseDate string
	Amount       float64
}

type fridaySummary struct {
	WeekStarting string
	TotalAmount  float64
	UserCount    int
}

func fridayPurchasesII(purchases []Purchase) []string {
	weekly := make(map[string]*fridaySummary)
	weekOrder := make([]string, 0)
	for _, p := range purchases {
		wk := p.PurchaseDate[:10]
		if _, ok := weekly[wk]; !ok {
			weekly[wk] = &fridaySummary{WeekStarting: wk}
			weekOrder = append(weekOrder, wk)
		}
		weekly[wk].TotalAmount += p.Amount
		weekly[wk].UserCount++
	}
	sort.Strings(weekOrder)
	var result []string
	for _, wk := range weekOrder {
		s := weekly[wk]
		result = append(result, fmt.Sprintf("%s|%.2f|%d", s.WeekStarting, s.TotalAmount, s.UserCount))
	}
	return result
}

func main() {
	purchases := []Purchase{
		{1, "2023-11-24", 100.50},
		{2, "2023-11-24", 50.25},
		{1, "2023-12-01", 200.00},
		{3, "2023-12-01", 75.00},
	}
	for _, r := range fridayPurchasesII(purchases) {
		fmt.Println(r)
	}
}
