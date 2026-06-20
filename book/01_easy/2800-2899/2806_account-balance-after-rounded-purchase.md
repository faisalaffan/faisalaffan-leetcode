# 2806 — Account Balance After Rounded Purchase

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func AccountBalanceAfterRoundedPurchase(purchaseAmount int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2806: Account Balance After Rounded Purchase
// https://leetcode.com/problems/account-balance-after-rounded-purchase/
// Difficulty: Easy
// Time: O(1) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(AccountBalanceAfterRoundedPurchase(9))
	fmt.Println(AccountBalanceAfterRoundedPurchase(15))
}

func AccountBalanceAfterRoundedPurchase(purchaseAmount int) int {
	return 100 - ((purchaseAmount+5)/10)*10
}
```
