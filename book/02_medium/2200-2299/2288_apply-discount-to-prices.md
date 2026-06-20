# 2288 — Apply Discount To Prices

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func discountPrices(sentence string, discount int) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2288: Apply Discount to Prices
// https://leetcode.com/problems/apply-discount-to-prices/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"strconv"
	"strings"
)

func discountPrices(sentence string, discount int) string {
	words := strings.Split(sentence, " ")
	for i, w := range words {
		if len(w) > 1 && w[0] == '$' {
			if numStr, err := strconv.Atoi(w[1:]); err == nil && numStr >= 0 && w[1] != '0' && w[1:] == strconv.Itoa(numStr) {
				price := float64(numStr) * (100.0 - float64(discount)) / 100.0
				words[i] = fmt.Sprintf("$%.2f", price)
			} else if err == nil && numStr == 0 && w == "$0" {
				words[i] = "$0.00"
			} else if err == nil && numStr == 0 && strings.TrimLeft(w[1:], "0") == "" {
				// All zeros: keep as valid price
				words[i] = "$0.00"
			}
		}
	}
	return strings.Join(words, " ")
}

func main() {
	// Test case 1
	fmt.Println(discountPrices("there are $1 $2 and 5$ candies in the shop", 50))
	// Expected: "there are $0.50 $1.00 and 5$ candies in the shop"

	// Test case 2
	fmt.Println(discountPrices("1 2 $3 4 $5 $6 7 8$ $9 $10$", 100))
	// Expected: "1 2 $0.00 4 $0.00 $0.00 7 8$ $0.00 $10$"
}
```
