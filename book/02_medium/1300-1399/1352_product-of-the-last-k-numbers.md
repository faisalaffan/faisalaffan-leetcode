# 1352 — Product Of The Last K Numbers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func Constructor() ProductOfNumbers`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Prefix Sum

**Waktu:** O(1) for adding  |  **Ruang:** O(n) for prefix array

> 🎓 **Fresh Grad Tips:** Kuasai **Prefix Sum** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1352: Product of the Last K Numbers
// https://leetcode.com/problems/product-of-the-last-k-numbers/
// Difficulty: Medium

import "fmt"

type ProductOfNumbers struct {
	prefix []int // prefix product, reset after 0
}

func main() {
	pn := Constructor()
	pn.Add(3)
	pn.Add(0)
	pn.Add(2)
	pn.Add(5)
	pn.Add(4)
	fmt.Println(pn.GetProduct(2)) // 20
	fmt.Println(pn.GetProduct(3)) // 40
	fmt.Println(pn.GetProduct(4)) // 0

	pn2 := Constructor()
	pn2.Add(1)
	pn2.Add(2)
	pn2.Add(3)
	pn2.Add(4)
	fmt.Println(pn2.GetProduct(1)) // 4
	fmt.Println(pn2.GetProduct(4)) // 24
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{prefix: []int{1}}
}

// Time: O(1) for adding
// Space: O(n) for prefix array
func (this *ProductOfNumbers) Add(num int) {
	if num == 0 {
		this.prefix = []int{1}
		return
	}
	this.prefix = append(this.prefix, this.prefix[len(this.prefix)-1]*num)
}

// Time: O(1) for query
func (this *ProductOfNumbers) GetProduct(k int) int {
	if k >= len(this.prefix) {
		return 0
	}
	return this.prefix[len(this.prefix)-1] / this.prefix[len(this.prefix)-1-k]
}
```
