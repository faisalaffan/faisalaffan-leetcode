package main

// LeetCode #1045: Customers Who Bought All Products
// https://leetcode.com/problems/customers-who-bought-all-products/
// Difficulty: Medium
//
// Approach: Count distinct products per customer, compare to total products.
// Time: O(n) where n = len(customer_product)
// Space: O(m) where m = number of customers

import "fmt"

func main() {
	// Simulated: customer_id, product_key
	customer := []int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
	product := []int{1, 1, 1, 1, 1, 2, 2, 2, 2, 2}
	fmt.Println(customersWhoBoughtAllProducts(customer, product, 2)) // [1,2,3,4,5]

	customer2 := []int{1, 1, 2}
	product2 := []int{1, 2, 1}
	fmt.Println(customersWhoBoughtAllProducts(customer2, product2, 2)) // [1]
}

func customersWhoBoughtAllProducts(customer, product []int, totalProducts int) []int {
	bought := make(map[int]map[int]bool)
	customerSet := make(map[int]bool)

	for i := 0; i < len(customer); i++ {
		c, p := customer[i], product[i]
		customerSet[c] = true
		if bought[c] == nil {
			bought[c] = make(map[int]bool)
		}
		bought[c][p] = true
	}

	result := make([]int, 0)
	for c := range customerSet {
		if len(bought[c]) == totalProducts {
			result = append(result, c)
		}
	}

	return result
}
