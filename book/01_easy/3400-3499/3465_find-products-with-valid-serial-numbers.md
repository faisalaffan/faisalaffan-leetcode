# 3465 — Find Products With Valid Serial Numbers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindProductsWithValidSerialNumbers(products []InvProduct) []InvProduct
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n). Space: O(n).  
**Kompleksitas Ruang:** O(n).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

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
