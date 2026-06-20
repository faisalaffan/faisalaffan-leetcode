# 3803 — Count Residue Prefixes

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func CountResiduePrefixes(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1) — at most 26 distinct chars

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3803: Count Residue Prefixes
// https://leetcode.com/problems/count-residue-prefixes/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CountResiduePrefixes("abc"))
	fmt.Println(CountResiduePrefixes("dd"))
	fmt.Println(CountResiduePrefixes("bob"))
}

// Time: O(n)
// Space: O(1) — at most 26 distinct chars
func CountResiduePrefixes(s string) int {
  // Membuat map (HashMap) — pencarian O(1)
	seen := make(map[byte]bool)
	count := 0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		seen[s[i]] = true
		if len(seen) == (i+1)%3 {
			count++
		}
	}
	return count
}
```
