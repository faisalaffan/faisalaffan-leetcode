# 3415 — Find Products With Three Consecutive Digits

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func FindProductsWithThreeConsecutiveDigits(products []Product) []Product`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n). Space: O(n).  |  **Ruang:** O(n).


## 💻 Solusi Go

```go
package main

// LeetCode #3415: Find Products with Three Consecutive Digits
// https://leetcode.com/problems/find-products-with-three-consecutive-digits/
// Difficulty: Easy [Paid]

import (
	"fmt"
	"regexp"
)

func main() {
	products := []Product{
		{ProductID: 1, Name: "ABC123XYZ"},
		{ProductID: 2, Name: "Product456"},
		{ProductID: 3, Name: "Item12"},
	}
	result := FindProductsWithThreeConsecutiveDigits(products)
	for _, p := range result {
		fmt.Printf("%d: %s\n", p.ProductID, p.Name)
	}
}

// Product represents a product.
type Product struct {
	ProductID int
	Name      string
}

// FindProductsWithThreeConsecutiveDigits returns products whose name contains at least three consecutive digits.
// Time: O(n). Space: O(n).
func FindProductsWithThreeConsecutiveDigits(products []Product) []Product {
	re := regexp.MustCompile(`\d{3,}`)
	result := []Product{}
	for _, p := range products {
		if re.MatchString(p.Name) {
			result = append(result, p)
		}
	}
	return result
}
```
