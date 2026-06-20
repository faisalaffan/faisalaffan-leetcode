# 1475 — Final Prices With A Special Discount In A Shop

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func finalPrices(prices []int) []int

import "fmt"

func main()
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n^2), Space: O(1) excluding output  
**Kompleksitas Ruang:** O(1) excluding output

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1475: Final Prices With a Special Discount in a Shop
// https://leetcode.com/problems/final-prices-with-a-special-discount-in-a-shop/
// Difficulty: Easy
//
// LeetCode submission: func finalPrices(prices []int) []int

import "fmt"

func main() {
	fmt.Println(FinalPricesWithASpecialDiscountInAShop([]int{8, 4, 6, 2, 3})) // [4 2 4 2 3]
	fmt.Println(FinalPricesWithASpecialDiscountInAShop([]int{1, 2, 3, 4, 5})) // [1 2 3 4 5]
}

// Time: O(n^2), Space: O(1) excluding output
func FinalPricesWithASpecialDiscountInAShop(prices []int) []int {
  // Alokasi slice integer
	res := make([]int, len(prices))
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(prices); i++ {
		res[i] = prices[i]
		for j := i + 1; j < len(prices); j++ {
			if prices[j] <= prices[i] {
				res[i] -= prices[j]
				break
			}
		}
	}
	return res
}
```
