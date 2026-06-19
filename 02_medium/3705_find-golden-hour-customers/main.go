package main

// LeetCode #3705: Find Golden Hour Customers
// https://leetcode.com/problems/find-golden-hour-customers/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

type order struct {
	customerID int
	hour       int
	rating     int // 0 means no rating
	hasRating  bool
}

type customerResult struct {
	customerID         int
	totalOrders        int
	peakHourPercentage int
	averageRating      float64
}

func findGoldenHourCustomers(orders []order) []customerResult {
	type stats struct {
		total     int
		peakCount int
		ratingSum int
		rated     int
	}

	custStats := make(map[int]*stats)
	for _, o := range orders {
		if _, ok := custStats[o.customerID]; !ok {
			custStats[o.customerID] = &stats{}
		}
		s := custStats[o.customerID]
		s.total++
		if (o.hour >= 11 && o.hour <= 13) || (o.hour >= 18 && o.hour <= 20) {
			s.peakCount++
		}
		if o.hasRating {
			s.ratingSum += o.rating
			s.rated++
		}
	}

	var results []customerResult
	for id, s := range custStats {
		if s.total < 3 {
			continue
		}
		peakPct := s.peakCount * 100 / s.total
		if peakPct < 60 {
			continue
		}
		if s.rated == 0 || s.rated*2 < s.total {
			continue
		}
		avgRating := float64(s.ratingSum) / float64(s.rated)
		if avgRating < 4.0 {
			continue
		}
		avgRounded := float64(int(avgRating*100+0.5)) / 100
		results = append(results, customerResult{
			customerID:         id,
			totalOrders:        s.total,
			peakHourPercentage: peakPct,
			averageRating:      avgRounded,
		})
	}

	sort.Slice(results, func(i, j int) bool {
		if results[i].averageRating != results[j].averageRating {
			return results[i].averageRating > results[j].averageRating
		}
		return results[i].customerID > results[j].customerID
	})

	return results
}

func main() {
	orders := []order{
		{101, 12, 5, true},
		{101, 13, 4, true},
		{101, 19, 5, true},
		{102, 10, 3, true},
		{102, 14, 4, true},
		{103, 11, 5, true},
		{103, 12, 4, true},
		{103, 18, 5, true},
		{105, 11, 4, true},
		{105, 19, 5, true},
		{105, 20, 4, true},
	}
	results := findGoldenHourCustomers(orders)
	for _, r := range results {
		fmt.Printf("Customer %d: orders=%d peak=%d%% rating=%.2f\n", r.customerID, r.totalOrders, r.peakHourPercentage, r.averageRating)
	}
	if len(results) == 0 {
		fmt.Println("[]")
	}
}
