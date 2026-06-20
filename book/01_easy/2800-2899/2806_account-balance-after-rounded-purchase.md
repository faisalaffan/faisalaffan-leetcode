# 2806 — Account Balance After Rounded Purchase

## Deskripsi

**Soal:** [2806. Account Balance After Rounded Purchase](https://leetcode.com/problems/account-balance-after-rounded-purchase/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** Dynamic Programming (DP)

## Solusi Go

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
