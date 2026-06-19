package main

// LeetCode #3415: Find Products with Three Consecutive Digits
// https://leetcode.com/problems/find-products-with-three-consecutive-digits/
// Difficulty: Easy [Paid]

import (
	"fmt"
	"regexp"
)

func main() {
	products := []Product{
		{ProductID: 1, Name: "ABC123XYZ"},
		{ProductID: 2, Name: "Product456"},
		{ProductID: 3, Name: "Item12"},
	}
	result := FindProductsWithThreeConsecutiveDigits(products)
	for _, p := range result {
		fmt.Printf("%d: %s\n", p.ProductID, p.Name)
	}
}

// Product represents a product.
type Product struct {
	ProductID int
	Name      string
}

// FindProductsWithThreeConsecutiveDigits returns products whose name contains at least three consecutive digits.
// Time: O(n). Space: O(n).
func FindProductsWithThreeConsecutiveDigits(products []Product) []Product {
	re := regexp.MustCompile(`\d{3,}`)
	result := []Product{}
	for _, p := range products {
		if re.MatchString(p.Name) {
			result = append(result, p)
		}
	}
	return result
}
