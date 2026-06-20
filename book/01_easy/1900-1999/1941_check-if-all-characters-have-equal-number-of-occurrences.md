# 1941 — Check If All Characters Have Equal Number Of Occurrences

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func CheckIfAllCharactersHaveEqualNumberOfOccurrences(s string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1) (max 26 chars)  
**Kompleksitas Ruang:** O(1) (max 26 chars)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1941: Check if All Characters Have Equal Number of Occurrences
// https://leetcode.com/problems/check-if-all-characters-have-equal-number-of-occurrences/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CheckIfAllCharactersHaveEqualNumberOfOccurrences("abacbc")) // true
	fmt.Println(CheckIfAllCharactersHaveEqualNumberOfOccurrences("aaabb"))  // false
}

// Time: O(n), Space: O(1) (max 26 chars)
func CheckIfAllCharactersHaveEqualNumberOfOccurrences(s string) bool {
  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[byte]int)
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		freq[s[i]]++
	}

	var target int
	for _, v := range freq {
		target = v
		break
	}
	for _, v := range freq {
		if v != target {
			return false
		}
	}
	return true
}
```
