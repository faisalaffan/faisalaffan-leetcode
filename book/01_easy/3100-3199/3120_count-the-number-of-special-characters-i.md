# 3120 — Count The Number Of Special Characters I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func CountTheNumberOfSpecialCharactersI(word string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3120: Count the Number of Special Characters I
// https://leetcode.com/problems/count-the-number-of-special-characters-i/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: numberOfSpecialChars
	fmt.Println(CountTheNumberOfSpecialCharactersI("aaAbcBC")) // 3
	fmt.Println(CountTheNumberOfSpecialCharactersI("abcd"))    // 0
	fmt.Println(CountTheNumberOfSpecialCharactersI("abAB"))   // 2
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: numberOfSpecialChars
func CountTheNumberOfSpecialCharactersI(word string) int {
  // Membuat map (HashMap) — pencarian O(1)
	lower := make(map[byte]bool)
  // Membuat map (HashMap) — pencarian O(1)
	upper := make(map[byte]bool)
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(word); i++ {
		c := word[i]
		if c >= 'a' && c <= 'z' {
			lower[c] = true
		} else if c >= 'A' && c <= 'Z' {
			upper[c] = true
		}
	}
	count := 0
	for c := byte('a'); c <= 'z'; c++ {
		if lower[c] && upper[c-'a'+'A'] {
			count++
		}
	}
	return count
}
```
