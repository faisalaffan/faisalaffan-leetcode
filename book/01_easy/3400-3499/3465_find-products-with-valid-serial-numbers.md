# 3465 — Find Products With Valid Serial Numbers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string berisi tanda kurung: `()`, `[]`, `{}`. Tugasmu adalah memeriksa apakah string tersebut **valid** — setiap kurung buka harus ditutup oleh kurung yang sesuai dalam urutan benar.

Contoh valid: `()[]{}`, `({[]})`. Tidak valid: `(]`, `([)]`.

**Cara berpikir:** Gunakan Stack. Kurung buka → push. Kurung tutup → pop dan cek kecocokan. Di akhir, stack harus kosong.

**Fungsi Solusi:** `func FindProductsWithValidSerialNumbers(products []InvProduct) []InvProduct`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n). Space: O(n).  |  **Ruang:** O(n).


## 💻 Solusi Go

```go
package main

// LeetCode #3465: Find Products with Valid Serial Numbers
// https://leetcode.com/problems/find-products-with-valid-serial-numbers/
// Difficulty: Easy

import (
	"fmt"
	"regexp"
)

func main() {
	products := []InvProduct{
		{ProductID: 1, Name: "Widget", SerialNumber: "SN-12345-ABC"},
		{ProductID: 2, Name: "Gadget", SerialNumber: "invalid"},
		{ProductID: 3, Name: "Doohickey", SerialNumber: "SN-67890-XYZ"},
	}
	result := FindProductsWithValidSerialNumbers(products)
	for _, p := range result {
		fmt.Printf("%d: %s (%s)\n", p.ProductID, p.Name, p.SerialNumber)
	}
}

// InvProduct represents a product with a serial number.
type InvProduct struct {
	ProductID    int
	Name         string
	SerialNumber string
}

// FindProductsWithValidSerialNumbers returns products whose serial number matches a valid pattern.
// Time: O(n). Space: O(n).
func FindProductsWithValidSerialNumbers(products []InvProduct) []InvProduct {
	re := regexp.MustCompile(`^SN-\d{5}-[A-Z]{3}$`)
	result := []InvProduct{}
	for _, p := range products {
		if re.MatchString(p.SerialNumber) {
			result = append(result, p)
		}
	}
	return result
}
```
