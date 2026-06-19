package main

// LeetCode #1596: The Most Frequently Ordered Products for Each Customer
// https://leetcode.com/problems/the-most-frequently-ordered-products-for-each-customer/
// Difficulty: Medium [Paid]

import "fmt"

func main() {
	// SQL problem: for each customer, find their most frequently ordered product(s).

	// Using maps directly instead of named structs
	orderData := []struct{ customerID, productID int }{
		{1, 10}, {1, 10}, {1, 20},
		{2, 20}, {2, 20}, {2, 30},
		{3, 10},
	}

	products := map[int]string{10: "Product A", 20: "Product B", 30: "Product C"}
	customers := map[int]string{1: "Alice", 2: "Bob", 3: "Charlie"}

	result := MostFrequentProducts(orderData, products, customers)
	fmt.Println("Most frequently ordered products per customer:")
	for _, r := range result {
		fmt.Printf("  %s -> %s (%d orders)\n", r.customerName, r.productName, r.count)
	}
}

type freqProduct struct {
	customerName string
	productName  string
	count        int
}

func MostFrequentProducts(orders []struct{ customerID, productID int }, products, customers map[int]string) []freqProduct {
	customerCounts := make(map[int]map[int]int)

	for _, o := range orders {
		if customerCounts[o.customerID] == nil {
			customerCounts[o.customerID] = make(map[int]int)
		}
		customerCounts[o.customerID][o.productID]++
	}

	result := make([]freqProduct, 0)
	for cid, prodCounts := range customerCounts {
		maxCount := 0
		for _, count := range prodCounts {
			if count > maxCount {
				maxCount = count
			}
		}
		for pid, count := range prodCounts {
			if count == maxCount {
				result = append(result, freqProduct{
					customerName: customers[cid],
					productName:  products[pid],
					count:        count,
				})
			}
		}
	}

	return result
}
