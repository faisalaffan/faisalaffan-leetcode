# 2082 — The Number Of Rich Customers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func TheNumberOfRichCustomers(transactions [][2]string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2082: The Number of Rich Customers
// https://leetcode.com/problems/the-number-of-rich-customers/
// Difficulty: Easy [Paid] (SQL)

import "fmt"

func main() {
	// (customer_id, amount)
	transactions := [][2]string{{"1", "500"}, {"2", "300"}, {"1", "600"}, {"3", "200"}}
	fmt.Println(TheNumberOfRichCustomers(transactions)) // 1 (only customer 1 has amount > 500)
}

// Time: O(n), Space: O(n)
func TheNumberOfRichCustomers(transactions [][2]string) int {
  // Membuat map (HashMap) — pencarian O(1)
	customers := make(map[int]bool)
	for _, t := range transactions {
		amount := 0
		for _, c := range t[1] {
			amount = amount*10 + int(c-'0')
		}
		if amount > 500 {
			id := 0
			for _, c := range t[0] {
				id = id*10 + int(c-'0')
			}
			customers[id] = true
		}
	}
	return len(customers)
}
```
