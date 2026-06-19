package main

// LeetCode #3465: Find Products with Valid Serial Numbers
// https://leetcode.com/problems/find-products-with-valid-serial-numbers/
// Difficulty: Easy

import (
	"fmt"
	"regexp"
)

func main() {
	products := []InvProduct{
		{ProductID: 1, Name: "Widget", SerialNumber: "SN-12345-ABC"},
		{ProductID: 2, Name: "Gadget", SerialNumber: "invalid"},
		{ProductID: 3, Name: "Doohickey", SerialNumber: "SN-67890-XYZ"},
	}
	result := FindProductsWithValidSerialNumbers(products)
	for _, p := range result {
		fmt.Printf("%d: %s (%s)\n", p.ProductID, p.Name, p.SerialNumber)
	}
}

// InvProduct represents a product with a serial number.
type InvProduct struct {
	ProductID    int
	Name         string
	SerialNumber string
}

// FindProductsWithValidSerialNumbers returns products whose serial number matches a valid pattern.
// Time: O(n). Space: O(n).
func FindProductsWithValidSerialNumbers(products []InvProduct) []InvProduct {
	re := regexp.MustCompile(`^SN-\d{5}-[A-Z]{3}$`)
	result := []InvProduct{}
	for _, p := range products {
		if re.MatchString(p.SerialNumber) {
			result = append(result, p)
		}
	}
	return result
}
