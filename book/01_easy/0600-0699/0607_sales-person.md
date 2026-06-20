# 0607 — Sales Person

## Deskripsi

**Soal:** [0607. Sales Person](https://leetcode.com/problems/sales-person/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

**Fungsi Solusi:** `func SalesPerson() string`

## Solusi Go

```go
package main

// LeetCode #607: Sales Person
// https://leetcode.com/problems/sales-person/
// Difficulty: Easy

import "fmt"

func SalesPerson() string {
	return "SELECT s.name FROM SalesPerson s WHERE s.sales_id NOT IN (SELECT o.sales_id FROM Orders o JOIN Company c ON o.com_id = c.com_id WHERE c.name = 'RED')"
}

func main() {
	fmt.Println(SalesPerson())
}
```
