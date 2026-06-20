# 1716 — Calculate Money In Leetcode Bank

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func TotalMoney(n int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(1), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1716: Calculate Money in Leetcode Bank
// https://leetcode.com/problems/calculate-money-in-leetcode-bank/
// Difficulty: Easy

import "fmt"

// Time: O(1), Space: O(1)
func TotalMoney(n int) int {
	weeks := n / 7
	days := n % 7
	// Sum of arithmetic progression: first week = 28, each week adds 7
	total := weeks*28 + 7*weeks*(weeks-1)/2
	// Remaining days
	total += days*(weeks+1) + days*(days-1)/2
	return total
}

func main() {
	fmt.Println(TotalMoney(4))
	fmt.Println(TotalMoney(10))
	fmt.Println(TotalMoney(20))
}
```
