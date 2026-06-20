# 0619 — Biggest Single Number

## Deskripsi

**Soal:** [0619. Biggest Single Number](https://leetcode.com/problems/biggest-single-number/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

**Fungsi Solusi:** `func BiggestSingleNumber() string`

## Solusi Go

```go
package main

// LeetCode #619: Biggest Single Number
// https://leetcode.com/problems/biggest-single-number/
// Difficulty: Easy

import "fmt"

func BiggestSingleNumber() string {
	return "SELECT MAX(num) AS num FROM (SELECT num FROM MyNumbers GROUP BY num HAVING COUNT(num) = 1) AS single_numbers"
}

func main() {
	fmt.Println(BiggestSingleNumber())
}
```
