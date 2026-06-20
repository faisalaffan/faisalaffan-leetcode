# 0584 — Find Customer Referee

## Deskripsi

**Soal:** [0584. Find Customer Referee](https://leetcode.com/problems/find-customer-referee/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

**Fungsi Solusi:** `func FindCustomerReferee() string`

## Solusi Go

```go
package main

// LeetCode #584: Find Customer Referee
// https://leetcode.com/problems/find-customer-referee/
// Difficulty: Easy

import "fmt"

func FindCustomerReferee() string {
	return "SELECT name FROM Customer WHERE referee_id != 2 OR referee_id IS NULL"
}

func main() {
	fmt.Println(FindCustomerReferee())
}
```
