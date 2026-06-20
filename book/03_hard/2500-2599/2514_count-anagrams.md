# 2514 — Count Anagrams

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func countAnagrams(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2514: Count Anagrams
// https://leetcode.com/problems/count-anagrams/
// Difficulty: Hard
//
// For each word, number of distinct anagrams = word_len! / product(cnt[char]!).
// Multiply across all words. Use modular arithmetic with MOD = 1e9+7.

import (
	"fmt"
	"strings"
)

const MOD = 1000000007

func main() {
	// Example 1: "too" => 2
	fmt.Println(countAnagrams("too"))
	// Example 2: "aa aa" => 1
	fmt.Println(countAnagrams("aa aa"))
	// Edge: single char
	fmt.Println(countAnagrams("a"))
	// Edge: all same letters
	fmt.Println(countAnagrams("aaa"))
	// Edge: multiple words
	fmt.Println(countAnagrams("abc def ghi"))
}

func countAnagrams(s string) int {
	words := strings.Fields(s)
	result := int64(1)

	for _, word := range words {
		n := len(word)
		// Count character frequencies
  // Membuat map (HashMap) — pencarian O(1)
		cnt := make(map[rune]int)
		for _, ch := range word {
			cnt[ch]++
		}

		// result *= n! / product(cnt[c]!)
		// Compute n! * inverse(product(cnt[c]!))
		res := factorial(n)
		for _, c := range cnt {
			inv := modInv(factorial(c))
			res = (res * inv) % MOD
		}
		result = (result * res) % MOD
	}

	return int(result)
}

func factorial(n int) int64 {
	res := int64(1)
	for i := 2; i <= n; i++ {
		res = (res * int64(i)) % MOD
	}
	return res
}

func modInv(a int64) int64 {
	return modPow(a, MOD-2)
}

func modPow(a int64, b int) int64 {
	res := int64(1)
	for b > 0 {
		if b&1 == 1 {
			res = (res * a) % MOD
		}
		a = (a * a) % MOD
		b >>= 1
	}
	return res
}
```
