# 0586 — Customer Placing The Largest Number Of Orders

## Deskripsi

**Soal:** [0586. Customer Placing The Largest Number Of Orders](https://leetcode.com/problems/customer-placing-the-largest-number-of-orders/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

**Fungsi Solusi:** `func CustomerPlacingTheLargestNumberOfOrders() string`

## Solusi Go

```go
package main

// LeetCode #586: Customer Placing the Largest Number of Orders
// https://leetcode.com/problems/customer-placing-the-largest-number-of-orders/
// Difficulty: Easy

import "fmt"

func CustomerPlacingTheLargestNumberOfOrders() string {
	return "SELECT customer_number FROM Orders GROUP BY customer_number ORDER BY COUNT(*) DESC LIMIT 1"
}

func main() {
	fmt.Println(CustomerPlacingTheLargestNumberOfOrders())
}
```
