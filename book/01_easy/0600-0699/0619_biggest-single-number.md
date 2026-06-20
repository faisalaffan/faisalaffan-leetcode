# 0619 — Biggest Single Number

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func BiggestSingleNumber() string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

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
