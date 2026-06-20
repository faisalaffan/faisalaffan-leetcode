# 1475 — Final Prices With A Special Discount In A Shop

## Deskripsi

**Soal:** [1475. Final Prices With A Special Discount In A Shop](https://leetcode.com/problems/final-prices-with-a-special-discount-in-a-shop/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n^2), Space: O(1) excluding output  
**Kompleksitas Ruang:** O(1) excluding output

**Algoritma:** —

**Fungsi Solusi:** `func finalPrices(prices []int) []int`

## Solusi Go

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
  // Membuat slice untuk menyimpan hasil
	res := make([]int, len(prices))
  // Loop standar: indeks 0 sampai n-1
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
