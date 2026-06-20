# 1357 — Apply Discount Every N Orders

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func NewCashier(n int, discount int, products []int, prices []int) Cashier`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(m) per bill where m = number of product types in the bill  |  **Ruang:** O(p) where p = number of unique products

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1357: Apply Discount Every n Orders
// https://leetcode.com/problems/apply-discount-every-n-orders/
// Difficulty: Medium

import "fmt"

type Cashier struct {
	n          int
	discount   int
	counter    int
	prices     map[int]int
}

func main() {
	cashier := NewCashier(3, 50, []int{1, 2, 3, 4, 5, 6, 7}, []int{100, 200, 300, 400, 300, 200, 100})
	fmt.Println(cashier.GetBill([]int{1, 2}, []int{1, 2}))    // 500.0
	fmt.Println(cashier.GetBill([]int{3, 7}, []int{10, 10}))  // 4000.0
	fmt.Println(cashier.GetBill([]int{1, 2, 3, 4, 5, 6, 7}, []int{1, 1, 1, 1, 1, 1, 1})) // 800.0
	fmt.Println(cashier.GetBill([]int{4}, []int{10}))         // 4000.0

	c2 := NewCashier(1, 10, []int{1, 2}, []int{100, 200})
	fmt.Println(c2.GetBill([]int{1}, []int{5})) // 450.0
}

func NewCashier(n int, discount int, products []int, prices []int) Cashier {
  // HashMap: O(1) lookup
	priceMap := make(map[int]int)
	for i, p := range products {
		priceMap[p] = prices[i]
	}
	return Cashier{n, discount, 0, priceMap}
}

// Time: O(m) per bill where m = number of product types in the bill
// Space: O(p) where p = number of unique products
func (this *Cashier) GetBill(product []int, amount []int) float64 {
	this.counter++
	total := 0
	for i, p := range product {
		total += this.prices[p] * amount[i]
	}

	result := float64(total)
	if this.counter%this.n == 0 {
		result = result * float64(100-this.discount) / 100.0
	}
	return result
}
```
