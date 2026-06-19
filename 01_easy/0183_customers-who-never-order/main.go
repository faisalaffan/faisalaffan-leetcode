package main

// LeetCode #183: Customers Who Never Order
// https://leetcode.com/problems/customers-who-never-order/
// Difficulty: Easy

import "fmt"

func CustomersWhoNeverOrder() string {
	return "SELECT c.name AS Customers FROM Customers c LEFT JOIN Orders o ON c.id = o.customerId WHERE o.customerId IS NULL"
}

func main() {
	fmt.Println(CustomersWhoNeverOrder())
}
