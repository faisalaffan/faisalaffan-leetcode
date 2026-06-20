# 2230 — The Users That Are Eligible For Discount

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func TheUsersThatAreEligibleForDiscount(purchases [][3]string) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2230: The Users That Are Eligible for Discount
// https://leetcode.com/problems/the-users-that-are-eligible-for-discount/
// Difficulty: Easy [Paid] (SQL)

import "fmt"

func main() {
	// User purchases: (user_id, amount, date)
	purchases := [][3]string{
		{"1", "120", "2024-01-15"},
		{"2", "50", "2024-02-01"},
		{"1", "80", "2024-03-01"},
	}
	fmt.Println(TheUsersThatAreEligibleForDiscount(purchases)) // [1]
}

// Time: O(n), Space: O(n)
func TheUsersThatAreEligibleForDiscount(purchases [][3]string) []int {
  // Membuat map (HashMap) — pencarian O(1)
	totals := make(map[int]int)
	for _, p := range purchases {
		id := 0
		for _, c := range p[0] {
			id = id*10 + int(c-'0')
		}
		amount := 0
		for _, c := range p[1] {
			amount = amount*10 + int(c-'0')
		}
		totals[id] += amount
	}

	var result []int
	for id, total := range totals {
		if total >= 100 {
			result = append(result, id)
		}
	}
	return result
}
```
