# 3415 — Find Products With Three Consecutive Digits

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindProductsWithThreeConsecutiveDigits(products []Product) []Product
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n). Space: O(n).  
**Kompleksitas Ruang:** O(n).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

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
