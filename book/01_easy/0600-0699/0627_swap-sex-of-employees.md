# 0627 — Swap Sex Of Employees

## Deskripsi

**Soal:** [0627. Swap Sex Of Employees](https://leetcode.com/problems/swap-sex-of-employees/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

**Fungsi Solusi:** `func SwapSexOfEmployees() string`

## Solusi Go

```go
package main

// LeetCode #627: Swap Sex of Employees
// https://leetcode.com/problems/swap-sex-of-employees/
// Difficulty: Easy

import "fmt"

func SwapSexOfEmployees() string {
	return "UPDATE Salary SET sex = CASE WHEN sex = 'm' THEN 'f' ELSE 'm' END"
}

func main() {
	fmt.Println(SwapSexOfEmployees())
}
```
