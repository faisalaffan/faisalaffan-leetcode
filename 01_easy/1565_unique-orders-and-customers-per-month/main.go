package main

// LeetCode #1565: Unique Orders and Customers Per Month
// https://leetcode.com/problems/unique-orders-and-customers-per-month/
// Difficulty: Easy [Paid]
//
// This is a SQL problem. The solution is the SQL query below.
// Table: Orders (order_id, order_date, customer_id, invoice)

import "fmt"

func main() {
	fmt.Println(UniqueOrdersAndCustomersPerMonth())
}

// Time: N/A (SQL query), Space: N/A
func UniqueOrdersAndCustomersPerMonth() string {
	return `SELECT
  DATE_FORMAT(order_date, '%Y-%m') AS month,
  COUNT(DISTINCT order_id) AS order_count,
  COUNT(DISTINCT customer_id) AS customer_count
FROM Orders
WHERE invoice > 20
GROUP BY month
ORDER BY month;`
}
