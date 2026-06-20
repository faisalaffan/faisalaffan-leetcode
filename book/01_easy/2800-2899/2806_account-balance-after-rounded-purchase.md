# 2806 — Account Balance After Rounded Purchase

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func AccountBalanceAfterRoundedPurchase(purchaseAmount int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

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
