# 3042 — Count Prefix And Suffix Pairs I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func CountPrefixAndSuffixPairsI(words []string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** O(n^2 * m) where m is max word length  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3042: Count Prefix and Suffix Pairs I
// https://leetcode.com/problems/count-prefix-and-suffix-pairs-i/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: countPrefixSuffixPairs
	fmt.Println(CountPrefixAndSuffixPairsI([]string{"a", "aba", "ababa", "aa"})) // 4
	fmt.Println(CountPrefixAndSuffixPairsI([]string{"pa", "papa", "ma", "mama"})) // 2
}

// Time: O(n^2 * m) where m is max word length | Space: O(1)
// LeetCode submission name: countPrefixSuffixPairs
func CountPrefixAndSuffixPairsI(words []string) int {
	n := len(words)
	count := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if isPrefixAndSuffix(words[i], words[j]) {
				count++
			}
		}
	}
	return count
}

func isPrefixAndSuffix(a, b string) bool {
	if len(a) > len(b) {
		return false
	}
	// Check prefix
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	// Check suffix
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(a); i++ {
		if a[i] != b[len(b)-len(a)+i] {
			return false
		}
	}
	return true
}
```
