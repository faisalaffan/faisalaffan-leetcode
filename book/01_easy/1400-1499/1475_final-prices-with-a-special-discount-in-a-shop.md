# 1475 — Final Prices With A Special Discount In A Shop

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func finalPrices(prices []int) []int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n^2), Space: O(1) excluding output  |  **Ruang:** O(1) excluding output


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
  // Alokasi slice
	res := make([]int, len(prices))
  // Linear scan O(n)
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
