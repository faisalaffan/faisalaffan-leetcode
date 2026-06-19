package main

// LeetCode #1398: Customers Who Bought Products A and B but Not C
// https://leetcode.com/problems/customers-who-bought-products-a-and-b-but-not-c/
// Difficulty: Medium

import "fmt"
import "sort"

func main() {
	// SQL problem - simulating in Go
	result := customersABnotC(
		[]struct {
			customerID   int
			customerName string
		}{
			{1, "Daniel"},
			{2, "Diana"},
			{3, "Elizabeth"},
			{4, "John"},
		},
		[]struct {
			orderID     int
			customerID  int
			productName string
		}{
			{1, 1, "A"},
			{2, 1, "B"},
			{3, 1, "C"},
			{4, 2, "A"},
			{5, 2, "B"},
			{6, 3, "A"},
		},
	)
	for _, r := range result {
		fmt.Printf("%d %s\n", r.id, r.name)
	}
	// Should output: 2 Diana (bought A and B but not C)
	// Note: 1 Daniel bought A and B but also C, excluded
	// Note: 3 Elizabeth bought A but not B, excluded
}

type customerResult struct {
	id   int
	name string
}

// Time: O(n) where n = number of orders
// Space: O(k) where k = number of customers
func customersABnotC(customers []struct {
	customerID   int
	customerName string
}, orders []struct {
	orderID     int
	customerID  int
	productName string
}) []customerResult {
	bought := make(map[int]map[string]bool)
	customerNames := make(map[int]string)

	for _, c := range customers {
		customerNames[c.customerID] = c.customerName
	}

	for _, o := range orders {
		if bought[o.customerID] == nil {
			bought[o.customerID] = make(map[string]bool)
		}
		bought[o.customerID][o.productName] = true
	}

	var result []customerResult
	for _, c := range customers {
		products := bought[c.customerID]
		if products["A"] && products["B"] && !products["C"] {
			result = append(result, customerResult{c.customerID, c.customerName})
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].id < result[j].id
	})

	return result
}
