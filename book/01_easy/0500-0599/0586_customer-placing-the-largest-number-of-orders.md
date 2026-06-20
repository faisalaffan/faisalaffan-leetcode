# 0586 — Customer Placing The Largest Number Of Orders

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func CustomerPlacingTheLargestNumberOfOrders() string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

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
