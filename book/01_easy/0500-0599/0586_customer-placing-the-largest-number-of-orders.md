# 0586 — Customer Placing The Largest Number Of Orders

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func CustomerPlacingTheLargestNumberOfOrders() string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

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
