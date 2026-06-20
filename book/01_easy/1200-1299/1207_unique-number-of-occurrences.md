# 1207 — Unique Number Of Occurrences

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func uniqueOccurrences(arr []int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1207: Unique Number of Occurrences
// https://leetcode.com/problems/unique-number-of-occurrences/
// Difficulty: Easy
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(uniqueOccurrences([]int{1, 2, 2, 1, 1, 3})) // true
	fmt.Println(uniqueOccurrences([]int{1, 2}))             // false
	fmt.Println(uniqueOccurrences([]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0})) // true
}

// LeetCode submission: uniqueOccurrences
func uniqueOccurrences(arr []int) bool {
  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[int]int)
	for _, v := range arr {
		freq[v]++
	}
  // Membuat map (HashMap) — pencarian O(1)
	seen := make(map[int]bool)
	for _, f := range freq {
		if seen[f] {
			return false
		}
		seen[f] = true
	}
	return true
}
```
