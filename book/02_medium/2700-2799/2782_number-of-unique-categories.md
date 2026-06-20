# 2782 — Number Of Unique Categories

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func NumberOfUniqueCategories(categories []string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2782: Number of Unique Categories
// https://leetcode.com/problems/number-of-unique-categories/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func NumberOfUniqueCategories(categories []string) int {
  // Membuat map (HashMap) — pencarian O(1)
	seen := make(map[string]bool)
	for _, c := range categories {
		seen[c] = true
	}
	return len(seen)
}

func main() {
	fmt.Println(NumberOfUniqueCategories([]string{"a", "b", "a", "c"}))
	fmt.Println(NumberOfUniqueCategories([]string{"x", "x", "x"}))
	fmt.Println(NumberOfUniqueCategories([]string{}))
}
```
