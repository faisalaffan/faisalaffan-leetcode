package main

// LeetCode #2686: Immediate Food Delivery III
// https://leetcode.com/problems/immediate-food-delivery-iii/
// Difficulty: Medium (SQL problem — simulated in Go)
//
// Simulates: For each order_date, calculate the percentage of immediate
// orders (where customer_pref_delivery_date == order_date).
// Round to 2 decimal places.

import (
	"fmt"
	"math"
	"sort"
)

// Delivery represents the Delivery database table.
type Delivery struct {
	DeliveryID              int
	CustomerID              int
	OrderDate               string // format: "YYYY-MM-DD"
	CustomerPrefDeliveryDate string // format: "YYYY-MM-DD"
}

// DeliveryResult holds the output.
type DeliveryResult struct {
	OrderDate           string
	ImmediatePercentage float64
}

// immediateFoodDeliveryIII simulates the SQL query.
// Time: O(n) | Space: O(d) where d = distinct order dates.
func immediateFoodDeliveryIII(deliveries []Delivery) []DeliveryResult {
	type dateStats struct {
		total     int
		immediate int
	}

	stats := make(map[string]*dateStats)

	for _, d := range deliveries {
		if stats[d.OrderDate] == nil {
			stats[d.OrderDate] = &dateStats{}
		}
		stats[d.OrderDate].total++
		if d.OrderDate == d.CustomerPrefDeliveryDate {
			stats[d.OrderDate].immediate++
		}
	}

	var results []DeliveryResult
	for date, s := range stats {
		// Round to 2 decimal places, matching SQL ROUND(x, 2).
		pct := math.Round(float64(s.immediate)/float64(s.total)*10000) / 100
		results = append(results, DeliveryResult{
			OrderDate:           date,
			ImmediatePercentage: pct,
		})
	}

	// Order by order_date ASC.
	sort.Slice(results, func(i, j int) bool {
		return results[i].OrderDate < results[j].OrderDate
	})

	return results
}

func main() {
	// Test data from the problem.
	deliveries := []Delivery{
		{DeliveryID: 1, CustomerID: 1, OrderDate: "2019-08-01", CustomerPrefDeliveryDate: "2019-08-01"},
		{DeliveryID: 2, CustomerID: 2, OrderDate: "2019-08-01", CustomerPrefDeliveryDate: "2019-08-02"},
		{DeliveryID: 3, CustomerID: 3, OrderDate: "2019-08-01", CustomerPrefDeliveryDate: "2019-08-01"},
		{DeliveryID: 4, CustomerID: 4, OrderDate: "2019-08-02", CustomerPrefDeliveryDate: "2019-08-02"},
		{DeliveryID: 5, CustomerID: 5, OrderDate: "2019-08-02", CustomerPrefDeliveryDate: "2019-08-03"},
		{DeliveryID: 6, CustomerID: 6, OrderDate: "2019-08-02", CustomerPrefDeliveryDate: "2019-08-02"},
		{DeliveryID: 7, CustomerID: 7, OrderDate: "2019-08-03", CustomerPrefDeliveryDate: "2019-08-03"},
		{DeliveryID: 8, CustomerID: 8, OrderDate: "2019-08-03", CustomerPrefDeliveryDate: "2019-08-03"},
		{DeliveryID: 9, CustomerID: 9, OrderDate: "2019-08-04", CustomerPrefDeliveryDate: "2019-08-05"},
		{DeliveryID: 10, CustomerID: 10, OrderDate: "2019-08-04", CustomerPrefDeliveryDate: "2019-08-06"},
	}

	results := immediateFoodDeliveryIII(deliveries)

	fmt.Println("Immediate Food Delivery III (order_date | immediate_percentage):")
	for _, r := range results {
		fmt.Printf("%s | %.2f\n", r.OrderDate, r.ImmediatePercentage)
	}
	// Expected output:
	// 2019-08-01 | 66.67
	// 2019-08-02 | 66.67
	// 2019-08-03 | 100.00
	// 2019-08-04 | 0.00
}
